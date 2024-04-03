package cmd

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"terraform-provider-solacebroker/cmd/generator"
	internalbroker "terraform-provider-solacebroker/internal/broker"
	"terraform-provider-solacebroker/internal/broker/generated"
	"terraform-provider-solacebroker/internal/semp"
)

type IdentifyingAttribute struct {
	key, value string
}

type BrokerObjectAttributes []IdentifyingAttribute // Described as a set of identifying attributes

var cachedResources = make(map[string]map[string]interface{})

// Only used in this demo, for real broker instances the name is obtrained from the broker
func getInstanceName(brokerObjectAttributes BrokerObjectAttributes) string {
	instanceNamePrefix := ""
	for i := 0; i < len(brokerObjectAttributes)-1; i++ {
		instanceNamePrefix += brokerObjectAttributes[i].value + "-"
	}
	return instanceNamePrefix + brokerObjectAttributes[len(brokerObjectAttributes)-1].value
}

// TODO: Use the real broker object attributes to generate the TF config
//   - Query object attribute settings from broker
//   - Remove attribute for which the value is set to default
//   - Replace value for attributes that are identifying.
//     Example:
//     For a subscriptionTopic object defined as
//     /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic}
//     Replace msgVpnName, queueName in the TF config by expressions derived from the parent
//
// This will generate the config for a particular object
func generateConfig(brokerObjectType generator.BrokerObjectType, brokerObjectAttributes BrokerObjectAttributes) {
	// Query object attributes from broker using SEMP GET
	instanceName := getInstanceName(brokerObjectAttributes) // only used in this demo
	fmt.Printf("  ## Generated config for %s instance:\n  resource \"solacebroker_%s\" \"%s\"  {}\n\n", instanceName, brokerObjectType, instanceName)
}

// Returns the path template for all instances of a broker object type, additionally the identifier attributes and the path template for a single instance
func getAllInstancesPathTemplate(brokerObjectType generator.BrokerObjectType) (string, []string, string, error) {
	pathTemplate, err := getInstancePathTemplate(brokerObjectType)
	if err != nil {
		return "", nil, "", err
	}
	// Example path template: /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic}
	// Example all instances path template: /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions
	sections := strings.Split(pathTemplate, "/")
	if len(sections) < 2 || !strings.Contains(sections[len(sections)-1], "{") || !strings.Contains(sections[len(sections)-1], "}") {
		return "", nil, "", fmt.Errorf("cannot create all resources query from path template: %s", pathTemplate)
	}
	allInstancesPathTemplate := strings.Join(sections[:len(sections)-1], "/")
	rex := regexp.MustCompile(`{[^{}]*}`)
	matches := rex.FindAllStringSubmatch(sections[len(sections)-1], -1)
	// flatten matches into identifierAttributes
	var identifierAttributes []string
	for _, match := range matches {
		identifierAttributes = append(identifierAttributes, strings.TrimSuffix(strings.TrimPrefix(match[0], "{"), "}"))
	}
	return allInstancesPathTemplate, identifierAttributes, pathTemplate, nil
}

func getInstancePathTemplate(brokerObjectType generator.BrokerObjectType) (string, error) {
	i, ok := generator.DSLookup[brokerObjectType]
	if !ok {
		return "", fmt.Errorf("invalid broker object type")
	}
	dsEntity := internalbroker.Entities[i]
	return dsEntity.PathTemplate, nil
}

func getRequestPath(pathTemplate string, attributes BrokerObjectAttributes) (string, error) {
	// Example path template: /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic}
	// Example brokerObjectAttributes: [IdentifyingAttribute{key: "msgVpnName", value: "myvpn"}, IdentifyingAttribute{key: "queueName", value: "myqueue"}, IdentifyingAttribute{key: "subscriptionTopic", value: "mysubscription"}]
	// Example path: /msgVpns/myvpn/queues/myqueue/subscriptions/mysubscription
	for _, attr := range attributes {
		pathTemplate = strings.Replace(pathTemplate, "{"+attr.key+"}", url.PathEscape(attr.value), -1)
	}
	if strings.Contains(pathTemplate, "{") || strings.Contains(pathTemplate, "}") {
		return "", fmt.Errorf("missing attributes from path template %s", pathTemplate)
	}
	return pathTemplate, nil
}

func identifierToBrokerObjectAttributes(brokerObjectType generator.BrokerObjectType, identifier string) (BrokerObjectAttributes, error) {
	pathTemplate, err := getInstancePathTemplate(brokerObjectType)
	if err != nil {
		return nil, err
	}
	identifierValues := map[int]string{}
	brokerObjectAttributes := BrokerObjectAttributes{}
	if strings.Contains(identifier, "/") {
		ids := strings.Split(identifier, "/")
		for i, val := range ids {
			identifierValues[i] = val
		}
	} else {
		identifierValues[0] = identifier
	}
	if !strings.Contains(pathTemplate, "{") {
		return brokerObjectAttributes, nil
	}
	rex := regexp.MustCompile(`{[^{}]*}`)
	matches := rex.FindAllStringSubmatch(pathTemplate, -1)
	if len(matches) < len(identifierValues) {
		return nil, fmt.Errorf("error: too many provider specific identifiers. Required identifiers: " + fmt.Sprint(matches))
	}
	for i := range identifierValues {
		decodedPathVar, _ := url.PathUnescape(fmt.Sprint(identifierValues[i]))
		value := url.PathEscape(decodedPathVar)
		brokerObjectAttributes = append(brokerObjectAttributes, IdentifyingAttribute{key: strings.TrimSuffix(strings.TrimPrefix(matches[i][0], "{"), "}"), value: value})
	}
	return brokerObjectAttributes, nil
}

// Return the list of instances
//
// TODO:
//   - Query all instances of a BrokerObjectType from the broker
//   - Consider using filters: e.g: "List of all MsgVpn names" at https://docs.solace.com/API-Developer-Online-Ref-Documentation/swagger-ui/software-broker/config/index.html
//
// Returns one instance of the brokerObjectType if identifier has been provided, otherwise all instances that match the parentIdentifyingAttributes
func getInstances(context context.Context, client semp.Client, brokerObjectType generator.BrokerObjectType, identifier string, parentIdentifyingAttributes BrokerObjectAttributes) ([]BrokerObjectAttributes, error) {
	var instances []BrokerObjectAttributes

	if identifier != "" {
		// Determine the identifying attributes for the instance
		instanceIdentifyingAttributes, err := identifierToBrokerObjectAttributes(brokerObjectType, identifier)
		if err != nil {
			return nil, err
		}
		// Query broker if resource exists
		resourcePathTemplate, err := getInstancePathTemplate(brokerObjectType)
		if err != nil {
			return nil, err
		}
		requestPath, err := getRequestPath(resourcePathTemplate, instanceIdentifyingAttributes)
		if err != nil {
			return nil, err
		}
		_, err = client.RequestWithoutBodyForGenerator(context, generated.BasePath, http.MethodGet, requestPath, []map[string]any{})
		if err != nil {
			return nil, err
		}
		instances = append(instances, instanceIdentifyingAttributes)
	} else {
		// Query broker for all instances that match the parentIdentifyingAttributes
		allResourcesPathTemplate, childIdentifierAttributes, resourceInstancePathTemplate, err := getAllInstancesPathTemplate(brokerObjectType)
		if err != nil {
			return nil, err
		}
		requestPath, err := getRequestPath(allResourcesPathTemplate, parentIdentifyingAttributes)
		if err != nil {
			return nil, err
		}
		results, err := client.RequestWithoutBodyForGenerator(context, generated.BasePath, http.MethodGet, requestPath, []map[string]any{})
		if err != nil {
			// return nil, err
			return instances, nil
		}
		for _, result := range results {
			// Extract the identifying attributes from the result
			foundChildIndentifyingAttributes := parentIdentifyingAttributes
			skipAppendInstance := false
			for _, childIdentifierAttribute := range childIdentifierAttributes {
				if isSystemProvisionedAttribute(result[childIdentifierAttribute].(string)) {
					skipAppendInstance = true
					break
				}
				foundChildIndentifyingAttributes = append(foundChildIndentifyingAttributes, IdentifyingAttribute{key: childIdentifierAttribute, value: result[childIdentifierAttribute].(string)})
			}
			if !skipAppendInstance {
				instances = append(instances, foundChildIndentifyingAttributes)
				// also cache the results for later use
				path, _ := getRequestPath(resourceInstancePathTemplate, foundChildIndentifyingAttributes)
				cachedResources[path] = result
			}
		}

	}
	return instances, nil
}

func isSystemProvisionedAttribute(attribute string) bool {
	return strings.HasPrefix(attribute, "#") && attribute != "#DEAD_MESSAGE_QUEUE"
}

// Iterates all instances of a child object
func generateConfigForObjectInstances(context context.Context, client semp.Client, brokerObjectType generator.BrokerObjectType, identifier string, parentIdentifyingAttributes BrokerObjectAttributes) error {
	// brokerObjectType is the current object type
	// instances is the list of instances of the current object type
	instances, err := getInstances(context, client, brokerObjectType, identifier, parentIdentifyingAttributes)
	if err != nil {
		return fmt.Errorf("aborting, run into %w", err)
	}
	for i := 0; i < len(instances); i++ {
		generateConfig(brokerObjectType, instances[i])
		for _, subType := range generator.BrokerObjectRelationship[brokerObjectType] {
			fmt.Printf("  Now processing subtype %s\n\n", subType)
			// Will need to pass additional params like the parent name etc. so to construct the appropriate names
			err := generateConfigForObjectInstances(context, client, subType, "", instances[i])
			if err != nil {
				return fmt.Errorf("aborting, run into issues")
			}
		}
	}
	return nil
}
