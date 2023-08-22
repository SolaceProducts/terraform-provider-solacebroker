package terraform

import (
	"context"
	"golang.org/x/exp/slices"
	"net/http"
	"os"
	"regexp"
	"strings"
	internalbroker "terraform-provider-solacebroker/internal/broker"
	"terraform-provider-solacebroker/internal/semp"
)

type BrokerObjectType string

type IdentifyingAttribute struct {
	key, value string
}
type BrokerObjectAttributes []IdentifyingAttribute

var BrokerObjectRelationship = map[BrokerObjectType][]BrokerObjectType{}

func CreateBrokerObjectRelationships() {
	for _, ds := range internalbroker.Entities {
		rex := regexp.MustCompile(`{[^{}]*}`)
		matches := rex.FindAllStringSubmatch(ds.PathTemplate, -1)

		for i, _ := range matches {

			if i == 0 {
				var parent string
				parent = strings.TrimPrefix((matches[i][0]), "{")
				parent = strings.TrimSuffix(parent, "}")

				//first item in node
				tfParentName, tfObjectExists := getDataSourceNameIfDatasource(parent, "")

				if tfObjectExists {
					_, ok := BrokerObjectRelationship[BrokerObjectType(tfParentName)]
					if !ok {
						BrokerObjectRelationship[BrokerObjectType(tfParentName)] = []BrokerObjectType{}
					}
				}
			}
			if i != 0 {
				var parent string
				var tfParentName string
				baseParent := strings.TrimPrefix((matches[0][0]), "{")
				baseParent = strings.TrimSuffix(baseParent, "}")

				parent = strings.TrimPrefix((matches[i-1][0]), "{")
				parent = strings.TrimSuffix(parent, "}")

				if baseParent != parent {
					tfParentName, _ = getDataSourceNameIfDatasource(baseParent, parent)
				} else {
					tfParentName, _ = getDataSourceNameIfDatasource(parent, "")
				}

				children, ok := BrokerObjectRelationship[BrokerObjectType(tfParentName)]
				if !ok {
					BrokerObjectRelationship[BrokerObjectType(tfParentName)] = []BrokerObjectType{}
					children = []BrokerObjectType{}
				}

				child := strings.TrimPrefix((matches[i][0]), "{")
				child = strings.TrimSuffix(child, "}")

				tfChildName, tfValueExists := getDataSourceNameIfDatasource(parent, child)

				if tfValueExists {
					valExists := slices.Contains(children, BrokerObjectType(tfChildName))
					if !valExists {
						children = append(children, BrokerObjectType(tfChildName))
						BrokerObjectRelationship[BrokerObjectType(tfParentName)] = children
					}
				}
			}
		}
	}
}

func getDataSourceNameIfDatasource(parent string, child string) (string, bool) {
	for _, ds := range internalbroker.Entities {
		rex := regexp.MustCompile(`{[^{}]*}`)
		matches := rex.FindAllStringSubmatch(ds.PathTemplate, -1)
		for i, _ := range matches {
			if len(matches) <= 2 {

				if len(matches) == 1 {
					parentToSet := strings.TrimPrefix((matches[i][0]), "{")
					parentToSet = strings.TrimSuffix(parentToSet, "}")

					if child == "" && parentToSet == parent {
						return ds.TerraformName, true
					}
				}
				if len(matches) == 2 && i == 1 {
					parentToSet := strings.TrimPrefix((matches[0][0]), "{")
					parentToSet = strings.TrimSuffix(parentToSet, "}")

					childToSet := strings.TrimPrefix((matches[1][0]), "{")
					childToSet = strings.TrimSuffix(childToSet, "}")

					if child == childToSet && parentToSet == parent {
						return ds.TerraformName, true
					}
				}
			}
		}
	}

	return "", false
}

func ParseTerraformObject(ctx context.Context, client semp.Client, resourceName string, brokerObjectTerraformName string, providerSpecificIdentifier string) map[string]string {
	tfObject := map[string]string{}
	LogCLIInfo("Generating terraform config for " + brokerObjectTerraformName)
	entityToRead := internalbroker.EntityInputs{}
	for _, ds := range internalbroker.Entities {
		if strings.ToLower(ds.TerraformName) == strings.ToLower(brokerObjectTerraformName) {
			entityToRead = ds
		}
	}

	path, err := ResolveSempPath(entityToRead.PathTemplate, providerSpecificIdentifier)
	if err != nil {
		LogCLIError("Error calling Broker Endpoint")
		os.Exit(1)
	}

	sempData, err := client.RequestWithoutBodyForGenerator(ctx, http.MethodGet, path)
	if err != nil {
		LogCLIError("SEMP called failed. " + err.Error() + " on path " + path)
		os.Exit(1)
	}

	resourceKey := "solacebroker_" + brokerObjectTerraformName + " " + resourceName

	resourceValues, err := GenerateTerraformString(entityToRead.Attributes, sempData)

	if len(resourceValues) == 1 {
		tfObject[strings.ToLower(resourceKey)] = resourceValues[0]
	} else {
		for i, _ := range resourceValues {
			tfObject[strings.ToLower(resourceKey)+ConvertToAlphabetic(i)] = resourceValues[i]
		}
	}
	return tfObject
}
