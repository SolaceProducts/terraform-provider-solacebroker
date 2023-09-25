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

type GeneratorTerraformOutput struct {
	TerraformOutput  map[string]string
	SEMPDataResponse map[string]map[string]any
}
type BrokerObjectAttributes []IdentifyingAttribute

var BrokerObjectRelationship = map[BrokerObjectType][]BrokerObjectType{}

func CreateBrokerObjectRelationships() {
	for _, ds := range internalbroker.Entities {
		rex := regexp.MustCompile(`{[^{}]*}`)
		matches := rex.FindAllStringSubmatch(ds.PathTemplate, -1)

		BrokerObjectRelationship[BrokerObjectType(ds.TerraformName)] = []BrokerObjectType{}

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

					if len(children) == 0 {
						children = []BrokerObjectType{BrokerObjectType(ds.TerraformName)}
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
						}
					}
					BrokerObjectRelationship[BrokerObjectType(tfParentName)] = children
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

func ParseTerraformObject(ctx context.Context, client semp.Client, resourceName string, brokerObjectTerraformName string, providerSpecificIdentifier string, parentBrokerResourceAttributesRelationship map[string]string, parentResult map[string]any) GeneratorTerraformOutput {
	var objectName string
	tfObject := map[string]string{}
	tfObjectSempDataResponse := map[string]map[string]any{}
	entityToRead := internalbroker.EntityInputs{}
	for _, ds := range internalbroker.Entities {
		if strings.ToLower(ds.TerraformName) == strings.ToLower(brokerObjectTerraformName) {
			entityToRead = ds
		}
	}
	var path string
	var err error

	if len(parentResult) > 0 {
		path, err = ResolveSempPathWithParent(entityToRead.PathTemplate, parentResult)
		if err != nil {
			LogCLIError("Error calling Broker Endpoint")
			os.Exit(1)
		}
	} else {
		path, err = ResolveSempPath(entityToRead.PathTemplate, providerSpecificIdentifier)
		if err != nil {
			LogCLIError("Error calling Broker Endpoint")
			os.Exit(1)
		}
	}

	sempData, err := client.RequestWithoutBodyForGenerator(ctx, generated.BasePath, http.MethodGet, path, []map[string]any{})
	if err != nil {
		if err == semp.ErrResourceNotFound {
			// continue if error is resource not found
			if len(parentResult) > 0 {
				print("..")
			} else {
				LogCLIError("SEMP call failed. " + err.Error() + " on path " + path)
			}
			sempData = []map[string]any{}
		} else {
			LogCLIError("SEMP call failed. " + err.Error() + " on path " + path)
			os.Exit(1)
		}
	}

	resourceKey := "solacebroker_" + brokerObjectTerraformName + " " + resourceName

	resourceValues, err := GenerateTerraformString(entityToRead.Attributes, sempData, parentBrokerResourceAttributesRelationship)

	for i := range resourceValues {
		objectName = strings.ToLower(resourceKey) + GetNameForResource(strings.ToLower(resourceKey), resourceValues[i])
		tfObject[objectName] = resourceValues[i]
		tfObjectSempDataResponse[objectName] = sempData[i]
	}
	return GeneratorTerraformOutput{
		TerraformOutput:  tfObject,
		SEMPDataResponse: tfObjectSempDataResponse,
	}
}

func GetNameForResource(resourceTerraformName string, attributeResourceTerraform string) string {

	resourceName := GenerateRandomString(6) //use generated if not able to identify

	resourceTerraformName = strings.Split(resourceTerraformName, " ")[0]
	resourceTerraformName = strings.ReplaceAll(strings.ToLower(resourceTerraformName), "solacebroker_", "")
	resources := ConvertAttributeTextToMap(attributeResourceTerraform)

	//Get identifying attribute name to differentiate from multiples
	for _, ds := range internalbroker.Entities {
		if ds.TerraformName == resourceTerraformName {
			for _, attr := range ds.Attributes {
				if attr.Identifying &&
					(strings.Contains(strings.ToLower(attr.TerraformName), "name") ||
						strings.Contains(strings.ToLower(attr.TerraformName), "topic")) {
					// intentionally continue looping till we get the best name
					value, found := resources[attr.TerraformName]
					if strings.Contains(value, ".") {
						continue
					}
					if found {
						//sanitize name
						value = strings.ReplaceAll(value, " ", "_")
						value = strings.ReplaceAll(value, "#", "_")
						value = strings.ReplaceAll(value, "\\", "_")
						value = strings.ReplaceAll(value, "/", "_")
						value = strings.ReplaceAll(value, "\"", "")
						resourceName = "_" + value
					}
				}
			}
		}
	}
	return resourceName
}
