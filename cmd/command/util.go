// terraform-provider-solacebroker
//
// Copyright 2023 Solace Corporation. All rights reserved.
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
package terraform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"terraform-provider-solacebroker/internal/broker"
	"text/tabwriter"
	"time"
)

type Color string

const (
	Reset Color = "\033[0m"
	Red   Color = "\033[31m"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type ResourceAttributeInfo struct {
	AttributeValue string
	Comment        string
}

type ResourceConfig struct {
	ResourceAttributes map[string]ResourceAttributeInfo // indexed by resource attribute name
}

type ObjectInfo struct {
	Registry        string
	BrokerURL       string
	Username        string
	Password        string
	BearerToken     string
	FileName        string
	BrokerResources []map[string]string
}

func StringWithDefaultFromEnv(name string, isMandatory bool, fallback string) string {
	envValue := os.Getenv("SOLACEBROKER_" + strings.ToUpper(name))
	if isMandatory && len(envValue) == 0 {
		LogCLIError("SOLACEBROKER_" + strings.ToUpper(name) + " is mandatory but not available")
		os.Exit(1)
	} else if len(envValue) == 0 {
		return fallback //default to fallback
	}
	return envValue
}

func Int64WithDefaultFromEnv(name string, isMandatory bool, fallback int64) (int64, error) {
	envName := "SOLACEBROKER_" + strings.ToUpper(name)
	s, ok := os.LookupEnv(envName)
	if !ok && isMandatory {
		return 0, errors.New("SOLACEBROKER_" + strings.ToUpper(name) + " is mandatory but not available")
	} else if !ok {
		return fallback, nil //default to fallback
	}
	return strconv.ParseInt(s, 10, 64)
}

func BooleanWithDefaultFromEnv(name string, isMandatory bool, fallback bool) (bool, error) {
	envName := "SOLACEBROKER_" + strings.ToUpper(name)
	s, ok := os.LookupEnv(envName)
	if !ok && isMandatory {
		return false, errors.New("SOLACEBROKER_" + strings.ToUpper(name) + " is mandatory but not available")
	} else if !ok {
		return fallback, nil //default to fallback
	}
	return strconv.ParseBool(s)
}

func DurationWithDefaultFromEnv(name string, isMandatory bool, fallback time.Duration) (time.Duration, error) {
	envValue := os.Getenv("SOLACEBROKER_" + strings.ToUpper(name))
	if isMandatory && len(envValue) == 0 {
		return 0, errors.New("SOLACEBROKER_" + strings.ToUpper(name) + " is mandatory but not available")
	} else if len(envValue) == 0 {
		return fallback, nil //default to fallback
	}
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"
	d, err := time.ParseDuration(envValue)
	if err != nil {
		return 0, errors.New(fmt.Errorf("%v is not valid; %q cannot be parsed as a duration: %w", "SOLACEBROKER_"+strings.ToUpper(name), envValue, err).Error())
	}
	return d, nil
}

func ResolveSempPath(pathTemplate string, v string) (string, error) {
	identifiersValues := map[int]string{}
	if strings.Contains(v, "/") {
		identifier := strings.Split(v, "/")
		for i, val := range identifier {
			identifiersValues[i] = val
		}
	} else {
		identifiersValues[0] = v
	}
	if !strings.Contains(pathTemplate, "{") {
		return pathTemplate, nil
	}
	rex := regexp.MustCompile(`{[^{}]*}`)
	out := rex.FindAllStringSubmatch(pathTemplate, -1)
	generatedPath := pathTemplate
	if len(out) < len(identifiersValues) {
		LogCLIError("\nError: Too many provider specific identifiers. Required identifiers: " + fmt.Sprint(out))
		os.Exit(1)
	}

	for i := range identifiersValues {
		if i < len(out) {
			//encode all, if param is url friendly this does nothing.
			value := url.QueryEscape(fmt.Sprint(identifiersValues[i]))
			generatedPath = strings.ReplaceAll(generatedPath, out[i][0], value)
		}
	}
	if len(out) > len(identifiersValues) {
		//remove unused vars
		for i := range out {
			generatedPath = strings.ReplaceAll(generatedPath, out[i][0], "")
		}
	}

	//special conditions
	//for /url/{param1},{param2},{param3}.. valid semp query to fetch all items when all parameters are empty
	//for /url1/{param1},{param2},{param3}/url2.. invalid when all parameters are empty as it becomes /url1/url2
	if strings.Contains(generatedPath, "/,,/") || strings.Contains(generatedPath, "/,/") {
		return "", errors.New("not all parameters found. SEMP call will be invalid")
	}

	path := strings.ReplaceAll(generatedPath, "/,,", "")
	path = strings.ReplaceAll(path, "/,", "")
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	if strings.Contains(path, "//") {
		return "", errors.New("not all parameters found. SEMP call will be invalid")
	}
	return path, nil
}
func ResolveSempPathWithParent(pathTemplate string, parentValues map[string]any) (string, error) {

	if !strings.Contains(pathTemplate, "{") {
		return pathTemplate, nil
	}
	rex := regexp.MustCompile(`{[^{}]*}`)
	out := rex.FindAllStringSubmatch(pathTemplate, -1)
	generatedPath := pathTemplate

	for i := range out {
		key := strings.TrimPrefix(out[i][0], "{")
		key = strings.TrimSuffix(key, "}")
		value, found := parentValues[key]

		if found {
			//url encode, if param is url friendly this does nothing.
			value = url.QueryEscape(fmt.Sprint(value))
			generatedPath = strings.ReplaceAll(generatedPath, out[i][0], fmt.Sprint(value))
		}
	}

	//remove unused vars
	for i := range out {
		generatedPath = strings.ReplaceAll(generatedPath, out[i][0], "")
	}
	//special conditions
	//for /url/{param1},{param2},{param3}.. valid semp query to fetch all items when all parameters are empty
	//for /url1/{param1},{param2},{param3}/url2.. invalid when all parameters are empty as it becomes /url1/url2
	if strings.Contains(generatedPath, "/,,/") || strings.Contains(generatedPath, "/,/") {
		return "", errors.New("not all parameters found. SEMP call will be invalid")
	}

	path := strings.ReplaceAll(generatedPath, "/,,", "")
	path = strings.ReplaceAll(path, "/,", "")
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	if strings.Contains(path, "//") {
		return "", errors.New("not all parameters found. SEMP call will be invalid")
	}
	return path, nil
}

func newAttributeInfo(value string) ResourceAttributeInfo {
	return ResourceAttributeInfo{
		AttributeValue: value,
		Comment:        "",
	}
}

func addCommentToAttributeInfo(info ResourceAttributeInfo, comment string) ResourceAttributeInfo {
	return ResourceAttributeInfo{
		AttributeValue: info.AttributeValue,
		Comment:        comment,
	}
}

func GenerateTerraformString(attributes []*broker.AttributeInfo, values []map[string]interface{}, parentBrokerResourceAttributes map[string]string, brokerObjectTerraformName string) ([]ResourceConfig, error) {
	var tfBrokerObjects []ResourceConfig
	var attributesWithDefaultValue = []string{} // list of attributes, collected but not used
	for k := range values {
		resourceConfig := ResourceConfig{
			ResourceAttributes: map[string]ResourceAttributeInfo{},
		}
		systemProvisioned := false
		for _, attr := range attributes {
			attributeParentNameAndValue, attributeExistInParent := parentBrokerResourceAttributes[attr.TerraformName]
			if attr.Sensitive {
				// write-only attributes can't be retrieved, so we don't expose them
				continue
			}
			if !attr.Identifying && attr.ReadOnly {
				// read-only attributes should only be in the datasource
				continue
			}
			valuesRes := values[k][attr.SempName]
			if attr.Identifying && attributeExistInParent {
				resourceConfig.ResourceAttributes[attr.TerraformName] = newAttributeInfo(attributeParentNameAndValue)
				continue
			} else if attr.TerraformName == "client_profile_name" && attributeExistInParent {
				//peculiar use case where client_profile is not identifying for msg_vpn_client_username but it is dependent
				resourceConfig.ResourceAttributes[attr.TerraformName] = newAttributeInfo(attributeParentNameAndValue)
				continue
			}

			switch attr.BaseType {
			case broker.String:
				if reflect.TypeOf(valuesRes) == nil || valuesRes == "" {
					continue
				}
				if attr.Identifying && valuesRes.(string)[0] == '#' && valuesRes.(string) != "#DEAD_MESSAGE_QUEUE" {
					systemProvisioned = true
				}
				if reflect.TypeOf(attr.Default) != nil && fmt.Sprint(attr.Default) == fmt.Sprint(valuesRes) {
					//attributes with default values will be skipped
					// WORKAROUND: Except if attribute is "authentication_basic_type" in "msg_vpn"
					if brokerObjectTerraformName != "msg_vpn" || attr.TerraformName != "authentication_basic_type" {
						attributesWithDefaultValue = append(attributesWithDefaultValue, attr.TerraformName)
						continue
					} else {
						LogCLIInfo("Applying workaround: not ignoring default for `msg_vpn` attribute `authentication_basic_type`")
					}
				}
				val := "\"" + SanitizeHclStringValue(valuesRes.(string)) + "\""
				resourceConfig.ResourceAttributes[attr.TerraformName] = newAttributeInfo(val)
			case broker.Int64:
				if valuesRes == nil {
					continue
				}
				intValue := valuesRes
				if reflect.TypeOf(attr.Default) != nil && fmt.Sprint(attr.Default) == fmt.Sprint(intValue) {
					//attributes with default values will be skipped
					attributesWithDefaultValue = append(attributesWithDefaultValue, attr.TerraformName)
					continue
				}
				val := fmt.Sprintf("%v", intValue)
				resourceConfig.ResourceAttributes[attr.TerraformName] = newAttributeInfo(val)
			case broker.Bool:
				if valuesRes == nil {
					continue
				}
				boolValue := valuesRes.(bool)
				if reflect.TypeOf(attr.Default) != nil && fmt.Sprint(attr.Default) == fmt.Sprint(boolValue) {
					//attributes with default values will be skipped
					attributesWithDefaultValue = append(attributesWithDefaultValue, attr.TerraformName)
					continue
				}
				val := strconv.FormatBool(boolValue)
				resourceConfig.ResourceAttributes[attr.TerraformName] = newAttributeInfo(val)
			case broker.Struct:
				valueJson, err := json.Marshal(valuesRes)
				if err != nil {
					continue
				}
				if reflect.TypeOf(attr.Default) != nil && fmt.Sprint(attr.Default) == fmt.Sprint(valuesRes) {
					//attributes with default values will be skipped
					attributesWithDefaultValue = append(attributesWithDefaultValue, attr.TerraformName)
					continue
				}
				output := strings.ReplaceAll(string(valueJson), "clearPercent", "clear_percent")
				output = strings.ReplaceAll(output, "setPercent", "set_percent")
				output = strings.ReplaceAll(output, "clearValue", "clear_value")
				output = strings.ReplaceAll(output, "setValue", "set_value")
				val := output
				resourceConfig.ResourceAttributes[attr.TerraformName] = newAttributeInfo(val)
			}
			if attr.Deprecated && systemProvisioned {
				addCommentToAttributeInfo(resourceConfig.ResourceAttributes[attr.TerraformName],
					" # Note: This attribute is deprecated and may also be system provisioned.")
			} else if attr.Deprecated && !systemProvisioned {
				addCommentToAttributeInfo(resourceConfig.ResourceAttributes[attr.TerraformName],
					" # Note: This attribute is deprecated.")
			} else if !attr.Deprecated && systemProvisioned {
				addCommentToAttributeInfo(resourceConfig.ResourceAttributes[attr.TerraformName],
					" # Note: This attribute may be system provisioned.")
			}
		}
		if !systemProvisioned {
			tfBrokerObjects = append(tfBrokerObjects, resourceConfig)
		} else {
			//add to maintain index, it will not be included in generation
			tfBrokerObjects = append(tfBrokerObjects, ResourceConfig{
				ResourceAttributes: nil,
			})
		}
	}
	return tfBrokerObjects, nil
}

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomString(n int) string {
	return "_" + randStr(n)
}

func LogCLIError(err string) {
	_, _ = fmt.Fprintf(os.Stdout, "%s %s %s\n", Red, err, Reset)
}

func LogCLIInfo(info string) {
	_, _ = fmt.Fprintf(os.Stdout, "\n%s %s %s", Reset, info, Reset)
}

func GetParentResourceAttributes(parentObjectName string, brokerParentResource map[string]ResourceConfig) map[string]string {
	parentResourceAttributes := map[string]string{}
	parentResourceName := strings.ReplaceAll(parentObjectName, " ", ".")
	for parentResourceObject := range brokerParentResource {
		resourceAttributes := brokerParentResource[parentResourceObject].ResourceAttributes
		for resourceAttributeName := range resourceAttributes {
			parentResourceAttributes[resourceAttributeName] = parentResourceName + "." + resourceAttributeName
		}
	}
	return parentResourceAttributes
}

func ConvertAttributeTextToMap(attribute string) map[string]string {
	attributeMap := map[string]string{}
	attributeSlice := strings.Split(attribute, "\n")
	for i := range attributeSlice {
		keyValue := strings.ReplaceAll(attributeSlice[i], "\t", "")
		if strings.Contains(keyValue, "=") {
			attributeMap[strings.Split(keyValue, "=")[0]] = strings.ReplaceAll(strings.Split(keyValue, "=")[1], "\"", "")
		}
	}
	return attributeMap
}

func IndexOf(elm BrokerObjectType, data []BrokerObjectType) int {
	for k, v := range data {
		if elm == v {
			return k
		}
	}
	return -1
}

func ToFormattedHCL(brokerResources []map[string]ResourceConfig) []map[string]string {
	var formattedResult []map[string]string
	for _, resources := range brokerResources {
		resourceCollection := make(map[string]string)
		for resourceTypeAndName := range resources {
			formattedResource := hclFormatResource(resources[resourceTypeAndName])
			resourceCollection[resourceTypeAndName] = formattedResource
		}
		formattedResult = append(formattedResult, resourceCollection)
	}
	return formattedResult
}

func hclFormatResource(resourceConfig ResourceConfig) string {
	var attributeNames []string
	for attributeName := range resourceConfig.ResourceAttributes {
		attributeNames = append(attributeNames, attributeName)
	}
	sort.Strings(attributeNames)
	var b bytes.Buffer
	w := tabwriter.NewWriter(&b, 0, 0, 2, ' ', 0)
	for pos := range attributeNames {
		attributeName := attributeNames[pos]
		attributeConfigLine := "\t" + attributeName + "\t" + "= "
		attributeConfigLine += resourceConfig.ResourceAttributes[attributeName].AttributeValue
		attributeConfigLine += resourceConfig.ResourceAttributes[attributeName].Comment
		fmt.Fprintln(w, attributeConfigLine)
	}
	w.Flush()
	config := b.String()
	return config
}

func SanitizeHclIdentifierName(name string) string {
	name = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(strings.ReplaceAll(name, " ", ""), "_")
	if len(name) == 0 || (name[0] >= '0' && name[0] <= '9') || (len(name) == 1 && name[0] == '_') {
		//just prepend static string to avoid checking all characters
		name = "gn_" + name
	}
	return name
}

func SanitizeHclStringValue(value string) string {
	b, err := json.Marshal(value)
	if err != nil {
	}
	s := string(b)
	output := s[1 : len(s)-1]
	output = strings.ReplaceAll(output, "${", "$${")
	output = strings.ReplaceAll(output, "%{", "%%{")
	return output
}
