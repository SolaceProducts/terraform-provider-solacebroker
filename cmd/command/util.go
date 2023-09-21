package terraform

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"terraform-provider-solacebroker/internal/broker"
	"time"
)

type Color string

const (
	Reset Color = "\033[0m"
	Red   Color = "\033[31m"
)

const (
	AttributesStart     string = "\t"
	AttributeKeyEnd            = "\t\t\t\t\t\t"
	AttributeValueStart        = "\t"
	AttributeValueEnd          = "\t\n"
	AttributesEnd              = "\n\t"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

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
			generatedPath = strings.ReplaceAll(generatedPath, out[i][0], identifiersValues[i])
		}
	}
	if len(out) > len(identifiersValues) {
		//remove unused vars
		for i := range out {
			generatedPath = strings.ReplaceAll(generatedPath, out[i][0], "")
		}
	}

	path := strings.TrimSuffix(generatedPath, ",")
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
		path = path + "?count=10"
	}
	return path, nil
}

func GenerateTerraformString(attributes []*broker.AttributeInfo, values []map[string]interface{}, parentBrokerResourceAttributes map[string]string) ([]string, error) {
	var tfBrokerObjects []string
	var attributesWithDefaultValue = []string{}
	for k := range values {
		tfAttributes := AttributesStart
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
				tfAttributes += attr.TerraformName + AttributeKeyEnd + "=" + AttributeValueStart + attributeParentNameAndValue + AttributeValueEnd + "\t"
				continue
			}
			switch attr.BaseType {
			case broker.String:
				if reflect.TypeOf(valuesRes) == nil || valuesRes == "" {
					continue
				}
				if attr.Identifying && strings.Contains(valuesRes.(string), "#") {
					systemProvisioned = true
				}
				if reflect.TypeOf(attr.Default) != nil && attr.Default == valuesRes.(string) {
					//attributes with default values will be skipped
					attributesWithDefaultValue = append(attributesWithDefaultValue, attr.TerraformName)
					continue
				}
				val := attr.TerraformName + AttributeKeyEnd + "=" + AttributeValueStart + "\"" + valuesRes.(string) + "\""
				if strings.Contains(valuesRes.(string), "{") {
					val = attr.TerraformName + AttributeKeyEnd + "=" + AttributeValueStart + valuesRes.(string)
				}
				tfAttributes += val
			case broker.Int64:
				if valuesRes == nil {
					continue
				}
				intValue := valuesRes
				if reflect.TypeOf(attr.Default) != nil && attr.Default == intValue {
					//attributes with default values will be skipped
					attributesWithDefaultValue = append(attributesWithDefaultValue, attr.TerraformName)
					continue
				}
				val := attr.TerraformName + AttributeKeyEnd + "=" + AttributeValueStart + fmt.Sprintf("%v", intValue)
				tfAttributes += val
			case broker.Bool:
				if valuesRes == nil {
					continue
				}
				boolValue := valuesRes.(bool)
				if reflect.TypeOf(attr.Default) != nil && attr.Default == boolValue {
					//attributes with default values will be skipped
					attributesWithDefaultValue = append(attributesWithDefaultValue, attr.TerraformName)
					continue
				}
				val := attr.TerraformName + AttributeKeyEnd + "=" + AttributeValueStart + strconv.FormatBool(boolValue)
				tfAttributes += val
			case broker.Struct:
				valueJson, err := json.Marshal(valuesRes)
				if err != nil {
					continue
				}
				if reflect.TypeOf(attr.Default) != nil && attr.Default == valuesRes {
					//attributes with default values will be skipped
					attributesWithDefaultValue = append(attributesWithDefaultValue, attr.TerraformName)
					continue
				}
				val := attr.TerraformName + AttributeKeyEnd + "=" + AttributeValueStart + string(valueJson)
				tfAttributes += val
			}
			if attr.Deprecated && systemProvisioned {
				tfAttributes += "	# Note: This attribute is deprecated and may also be system provisioned."
			} else if attr.Deprecated && !systemProvisioned {
				tfAttributes += "	# Note: This attribute is deprecated."
			} else if !attr.Deprecated && systemProvisioned {
				tfAttributes += "	# Note: This attribute may be system provisioned."
			}
			tfAttributes += AttributesEnd
		}
		if !systemProvisioned {
			tfBrokerObjects = append(tfBrokerObjects, tfAttributes)
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

func GetParentResourceAttributes(brokerParentResource map[string]string) map[string]string {
	parentResourceAttributes := map[string]string{}
	for parentResourceObject := range brokerParentResource {
		resourceAttributes := strings.Split(brokerParentResource[parentResourceObject], "\n")
		for n := range resourceAttributes {
			if len(strings.TrimSpace(resourceAttributes[n])) > 0 {
				parentResourceName := strings.ReplaceAll(parentResourceObject, " ", ".")
				parentResourceAttribute := strings.Split(strings.Replace(resourceAttributes[n], "\t", "", -1), "=")[0]
				parentResourceAttributes[parentResourceAttribute] = parentResourceName + "." + parentResourceAttribute
			}
		}
	}
	return parentResourceAttributes
}
