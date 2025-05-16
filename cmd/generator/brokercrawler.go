// terraform-provider-solacebroker
//
// Copyright 2024 Solace Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package generator

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	internalbroker "terraform-provider-solacebroker/internal/broker"
	"terraform-provider-solacebroker/internal/broker/generated"
	"terraform-provider-solacebroker/internal/semp"
)

type IdentifyingAttribute struct {
	key, value string
}

type IdentifyingAttributes []IdentifyingAttribute // Described as a set of identifying attributes

type BrokerObjectInstanceInfo struct {
	resourceTypeAndName   string
	identifyingAttributes IdentifyingAttributes
}

var rootBrokerObjectPathTemplate string
var rootBrokerObjectResourceName string
var cachedResources map[string]interface{}
var brokerResources []map[string]ResourceConfig
var variables map[string]VariableConfig

func buildResourceTypeAndName(brokerObjectType BrokerObjectType, resourceInstancePathTemplate string, foundChildIndentifyingAttributes IdentifyingAttributes) (string, error) {
	var resourceTypeAndName string
	// Replace rootBrokerObjectPathTemplate part with rootBrokerObjectResourceName
	convertedPath := strings.Replace(resourceInstancePathTemplate, rootBrokerObjectPathTemplate, rootBrokerObjectResourceName, 1)
	// Split path by /
	sections := strings.Split(convertedPath, "/")
	resourceTypeAndName = sections[0]
	for i := 1; i < len(sections); i++ {
		if strings.HasPrefix(sections[i], "{") && strings.HasSuffix(sections[i], "}") {
			// in case of multiple attributes replace all },{ with }_{
			adjustedSection := strings.Replace(sections[i], "},{", "}_{", -1)
			resourceTypeAndName += adjustedSection
		} else {
			resourceTypeAndName += "_"
		}
	}
	var err error
	resourceTypeAndName, err = substituteVariables(resourceTypeAndName, foundChildIndentifyingAttributes, false)
	if err != nil {
		return "", err
	}
	resourceTypeAndName = string(brokerObjectType) + " " + makeValidForTerraformIdentifier(resourceTypeAndName)
	resourceTypeAndName = "solacebroker_" + resourceTypeAndName
	// Loop while cachedResources has an entry with this key. Change the key appending __number until it is unique
	i := 2
	modifiedResourceTypeAndName := resourceTypeAndName
	for cachedResources[modifiedResourceTypeAndName] != nil {
		modifiedResourceTypeAndName = resourceTypeAndName + "__" + fmt.Sprint(i)
		i++
	}
	resourceTypeAndName = modifiedResourceTypeAndName
	return resourceTypeAndName, nil
}

// Returns the path template for all instances of a broker object type, additionally the identifier attributes and the path template for a single instance
func getAllInstancesPathTemplate(brokerObjectType BrokerObjectType) (string, []string, string, error) {
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

func getInstancePathTemplate(brokerObjectType BrokerObjectType) (string, error) {
	i, ok := DSLookup[brokerObjectType]
	if !ok {
		return "", fmt.Errorf("invalid broker object type")
	}
	dsEntity := internalbroker.Entities[i]
	return dsEntity.PathTemplate, nil
}

// This function substitutes variables in a template string with values from the attributes, optionally escaping the values
func substituteVariables(template string, attributes IdentifyingAttributes, doEscape bool) (string, error) {
	// Example template: /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic}
	// Example brokerObjectAttributes: [IdentifyingAttribute{key: "msgVpnName", value: "myvpn"}, IdentifyingAttribute{key: "queueName", value: "myqueue"}, IdentifyingAttribute{key: "subscriptionTopic", value: "mysubscription"}]
	// Example results: /msgVpns/myvpn/queues/myqueue/subscriptions/mysubscription
	result := template
	for _, attr := range attributes {
		var value string
		if doEscape {
			value = url.PathEscape(attr.value)
		} else {
			value = attr.value
		}
		result = strings.Replace(result, "{"+attr.key+"}", value, -1)
	}
	return result, nil
}

func identifierToBrokerObjectAttributes(brokerObjectType BrokerObjectType, identifier string) (IdentifyingAttributes, error) {
	pathTemplate, err := getInstancePathTemplate(brokerObjectType)
	if err != nil {
		return nil, err
	}
	identifierValues := map[int]string{}
	brokerObjectAttributes := IdentifyingAttributes{}
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
	if len(matches) != len(identifierValues) {
		return nil, fmt.Errorf("incorrect identifier: \"%q\". Following required identifier elements are expected: %v", identifier, matches)
	}
	for i := range identifierValues {
		decodedPathVar, _ := url.PathUnescape(fmt.Sprint(identifierValues[i]))
		brokerObjectAttributes = append(brokerObjectAttributes, IdentifyingAttribute{key: strings.TrimSuffix(strings.TrimPrefix(matches[i][0], "{"), "}"), value: decodedPathVar})
	}
	return brokerObjectAttributes, nil
}

// Returns one instance of the brokerObjectType if identifier has been provided, otherwise all instances that match the parentIdentifyingAttributes
// Communicates with the broker via the SEMP client to fetch the instances
// As a side effect, it will also construct an identifier for an object instance, prep the attributes and cache the results for later use
func getInstances(context context.Context, client semp.Client, brokerObjectType BrokerObjectType, identifier string, parent BrokerObjectInstanceInfo) ([]BrokerObjectInstanceInfo, error) {
	var instances []BrokerObjectInstanceInfo

	if identifier != "" {
		// Return a single instance of the brokerObjectType that matches the identifier
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
		requestPath, err := substituteVariables(resourcePathTemplate, instanceIdentifyingAttributes, true)
		if err != nil {
			return nil, err
		}
		results, err := client.RequestWithoutBodyForGenerator(context, generated.BasePath, http.MethodGet, requestPath, []map[string]any{})
		if err != nil {
			return nil, err
		}
		resourceTypeAndName, err := buildResourceTypeAndName(brokerObjectType, resourcePathTemplate, instanceIdentifyingAttributes)
		if err != nil {
			return nil, err
		}
		// create a resource config from results[0]
		attributes := internalbroker.Entities[DSLookup[BrokerObjectType(brokerObjectType)]].Attributes
		resourceValues, tfVariables, err := processSempResults(resourceTypeAndName, attributes, results, BrokerObjectInstanceInfo{})
		if err != nil {
			return nil, err
		}
		element := make(map[string]ResourceConfig)
		element[resourceTypeAndName] = resourceValues[0]
		brokerResources = append(brokerResources, element)
		for key, value := range tfVariables {
			variables[key] = value
		}
		instances = append(instances, BrokerObjectInstanceInfo{
			resourceTypeAndName:   resourceTypeAndName,
			identifyingAttributes: instanceIdentifyingAttributes,
		})

	} else {
		// Query broker for all instances that match the parentIdentifyingAttributes
		allResourcesPathTemplate, childIdentifierAttributes, resourceInstancePathTemplate, err := getAllInstancesPathTemplate(brokerObjectType)
		if err != nil {
			return nil, err
		}
		requestPath, err := substituteVariables(allResourcesPathTemplate, parent.identifyingAttributes, true)
		if err != nil {
			return nil, err
		}
		results, err := client.RequestWithoutBodyForGenerator(context, generated.BasePath, http.MethodGet, requestPath, []map[string]any{})
		if err != nil {
			// Fail except if the path is invalid - this means the generator SEMP schema is trying
			// to fetch a resource that doesn't exist in an older broker
			if !errors.Is(err, semp.ErrInvalidPath) {
				return nil, err
			}
			LogCLIInfo(fmt.Sprintf("     Resource %s unknown on broker, check broker and generator SEMP versions\n", brokerObjectType))
		}
		for _, result := range results {
			// Extract the identifying attributes from the result
			foundChildIndentifyingAttributes := parent.identifyingAttributes
			skipAppendInstance := false
			for _, childIdentifierAttribute := range childIdentifierAttributes {
				// Skip system provisioned objects
				if isSystemProvisionedAttribute(result[childIdentifierAttribute].(string)) {
					// Workaround while waiting for SOL-117252
					if string(brokerObjectType) == "msg_vpn_acl_profile" || string(brokerObjectType) == "msg_vpn_client_profile" || string(brokerObjectType) == "msg_vpn_client_username" || string(brokerObjectType) == "msg_vpn" {
						skipAppendInstance = true
						break
					}
				}
				foundChildIndentifyingAttributes = append(foundChildIndentifyingAttributes, IdentifyingAttribute{key: childIdentifierAttribute, value: result[childIdentifierAttribute].(string)})
			}
			if !skipAppendInstance {
				// also cache the results for later use
				resourceTypeAndName, err := buildResourceTypeAndName(brokerObjectType, resourceInstancePathTemplate, foundChildIndentifyingAttributes)
				if err != nil {
					return nil, err
				}
				cachedResources[resourceTypeAndName] = "" // using cachedResources as a set
				// create a resource config from result
				var elems []map[string]interface{}
				elems = append(elems, result)
				attributes := internalbroker.Entities[DSLookup[BrokerObjectType(brokerObjectType)]].Attributes
				resourceValues, tfVariables, err := processSempResults(resourceTypeAndName, attributes, elems, parent)
				if err != nil {
					return nil, err
				}
				element := make(map[string]ResourceConfig)
				element[resourceTypeAndName] = resourceValues[0]
				brokerResources = append(brokerResources, element)
				for key, value := range tfVariables {
					variables[key] = value
				}
				instances = append(instances, BrokerObjectInstanceInfo{
					resourceTypeAndName:   resourceTypeAndName,
					identifyingAttributes: foundChildIndentifyingAttributes,
				})
			}
		}

	}
	return instances, nil
}

// Main entry point to generate the config for a broker object
func fetchBrokerConfig(context context.Context, client semp.Client, brokerObjectType BrokerObjectType, brokerResourceName string, identifier string) ([]map[string]ResourceConfig, map[string]VariableConfig, error) {
	var err error
	cachedResources = make(map[string]interface{})
	variables = map[string]VariableConfig{}
	rootBrokerObjectResourceName = brokerResourceName
	rootBrokerObjectPathTemplate, err = getInstancePathTemplate(brokerObjectType)
	if err != nil {
		return nil, nil, err
	}
	err = GenerateConfigForObjectInstances(context, client, brokerObjectType, identifier, BrokerObjectInstanceInfo{})
	if err != nil {
		return nil, nil, err
	}
	return brokerResources, variables, nil
}

// This is a recursive function that generates the config for a broker object and its children
// The entry point is the parent object with the identifier. For child objects the identifier is empty
// It will call itself for each child object instance
func GenerateConfigForObjectInstances(context context.Context, client semp.Client, brokerObjectType BrokerObjectType, identifier string, parentInstanceInfo BrokerObjectInstanceInfo) error {
	// brokerObjectType is the current object type
	// instances is the list of instances of the current object type
	LogCLIInfo(fmt.Sprintf("  ## Fetching config for resource %s\n", brokerObjectType))
	instances, err := getInstances(context, client, brokerObjectType, identifier, parentInstanceInfo)
	if err != nil {
		return err
	}
	for i := 0; i < len(instances); i++ {
		for _, subType := range BrokerObjectRelationship[brokerObjectType] {
			// Will need to pass additional params like the parent name etc. so to construct the appropriate names
			err := GenerateConfigForObjectInstances(context, client, subType, "", instances[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
