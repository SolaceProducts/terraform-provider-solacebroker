package terraform

import (
	"context"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"net/http"
	"os"
	"regexp"
	"strings"
	internalbroker "terraform-provider-solacebroker/internal/broker"
	"terraform-provider-solacebroker/internal/broker/generated"
	"terraform-provider-solacebroker/internal/semp"
)

type BrokerObjectType string

type IdentifyingAttribute struct {
	key, value string
}
type BrokerObjectAttributes []IdentifyingAttribute

var BrokerObjectRelationship = map[BrokerObjectType][]BrokerObjectType{}

func CreateBrokerObjectRelationships() {
	BrokerObjectRelationship[("broker")] = []BrokerObjectType{}
	BrokerObjectRelationship[("msg_vpn_queue")] = []BrokerObjectType{
		"msg_vpn_queue_subscription",
	}
	for _, ds := range internalbroker.Entities {
		rex := regexp.MustCompile(`{[^{}]*}`)
		matches := rex.FindAllStringSubmatch(ds.PathTemplate, -1)

		for i := range matches {

			if i == 0 {
				var parent string
				parent = strings.TrimPrefix(matches[i][0], "{")
				parent = strings.TrimSuffix(parent, "}")

				//first item in node
				tfParentNameCollection, tfObjectExists := getDataSourceNameIfDatasource(parent, "")

				if tfObjectExists {
					for _, tfParentName := range tfParentNameCollection {
						_, ok := BrokerObjectRelationship[BrokerObjectType(tfParentName)]
						if !ok {
							BrokerObjectRelationship[BrokerObjectType(tfParentName)] = []BrokerObjectType{}
						}
					}
				}
			}
			if i != 0 {
				var parent string
				var tfParentNameCollection []string
				baseParent := strings.TrimPrefix(matches[0][0], "{")
				baseParent = strings.TrimSuffix(baseParent, "}")

				parent = strings.TrimPrefix(matches[i-1][0], "{")
				parent = strings.TrimSuffix(parent, "}")

				if baseParent != parent {
					tfParentNameCollection, _ = getDataSourceNameIfDatasource(baseParent, parent)
				} else {
					tfParentNameCollection, _ = getDataSourceNameIfDatasource(parent, "")
				}

				for _, tfParentName := range tfParentNameCollection {

					children, ok := BrokerObjectRelationship[BrokerObjectType(tfParentName)]
					if !ok {
						BrokerObjectRelationship[BrokerObjectType(tfParentName)] = []BrokerObjectType{}
						children = []BrokerObjectType{}
					}

					child := strings.TrimPrefix(matches[i][0], "{")
					child = strings.TrimSuffix(child, "}")

					tfChildNameCollection, tfValueExists := getDataSourceNameIfDatasource(parent, child)

					if tfValueExists {
						for _, tfChildName := range tfChildNameCollection {
							valExists := slices.Contains(children, BrokerObjectType(tfChildName))
							if !valExists {
								children = append(children, BrokerObjectType(tfChildName))
								BrokerObjectRelationship[BrokerObjectType(tfParentName)] = children
							}
							_, ok := BrokerObjectRelationship[BrokerObjectType(tfChildName)]
							if !ok {
								BrokerObjectRelationship[BrokerObjectType(tfChildName)] = []BrokerObjectType{}
							}
						}
					}
				}
			}
		}
	}
}

func getDataSourceNameIfDatasource(parent string, child string) ([]string, bool) {
	tfNames := map[string]string{}
	for _, ds := range internalbroker.Entities {
		rex := regexp.MustCompile(`{[^{}]*}`)
		matches := rex.FindAllStringSubmatch(ds.PathTemplate, -1)
		for i := range matches {
			if len(matches) <= 2 {

				if len(matches) == 1 {
					parentToSet := strings.TrimPrefix(matches[i][0], "{")
					parentToSet = strings.TrimSuffix(parentToSet, "}")

					if child == "" && parentToSet == parent {
						_, exists := tfNames[ds.TerraformName]
						if !exists {
							tfNames[ds.TerraformName] = ds.TerraformName
						}
					}
				}
				if len(matches) == 2 && i == 1 {
					parentToSet := strings.TrimPrefix(matches[0][0], "{")
					parentToSet = strings.TrimSuffix(parentToSet, "}")

					childToSet := strings.TrimPrefix(matches[1][0], "{")
					childToSet = strings.TrimSuffix(childToSet, "}")

					if child == childToSet && parentToSet == parent {
						_, exists := tfNames[ds.TerraformName]
						if !exists {
							tfNames[ds.TerraformName] = ds.TerraformName
						}
					}
				}
			}
		}
	}
	collectionTfNames := maps.Keys(tfNames)
	return collectionTfNames, len(collectionTfNames) > 0
}

func ParseTerraformObject(ctx context.Context, client semp.Client, resourceName string, brokerObjectTerraformName string, providerSpecificIdentifier string, parentBrokerResourceAttributes map[string]string) map[string]string {
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

	sempData, err := client.RequestWithoutBodyForGenerator(ctx, generated.BasePath, http.MethodGet, path, []map[string]any{})
	if err != nil {
		LogCLIError("SEMP called failed. " + err.Error() + " on path " + path)
		os.Exit(1)
	}

	resourceKey := "solacebroker_" + brokerObjectTerraformName + " " + resourceName

	resourceValues, err := GenerateTerraformString(entityToRead.Attributes, sempData, parentBrokerResourceAttributes)

	if len(resourceValues) == 1 {
		tfObject[strings.ToLower(resourceKey)] = resourceValues[0]
	} else {
		for i := range resourceValues {
			tfObject[strings.ToLower(resourceKey)+GenerateRandomString(6)] = resourceValues[i]
		}
	}
	return tfObject
}
