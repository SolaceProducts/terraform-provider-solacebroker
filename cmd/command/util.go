package terraform

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"terraform-provider-solacebroker/internal/broker"
	"time"
)

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
		return 0, errors.New(fmt.Errorf("%v is not valid; %q cannot be parsed as a duration: %w", ("SOLACEBROKER_" + strings.ToUpper(name)), envValue, err).Error())
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

	for i, _ := range identifiersValues {
		if i < len(out) {
			generatedPath = strings.ReplaceAll(generatedPath, out[i][0], identifiersValues[i])
		}
	}
	if len(out) > len(identifiersValues) {
		//remove unused vars
		for i, _ := range out {
			generatedPath = strings.ReplaceAll(generatedPath, out[i][0], "")
		}
	}

	path := strings.TrimSuffix(generatedPath, ",")
	path = strings.TrimSuffix(path, "/")
	return path, nil
}

func GenerateTerraformString(attributes []*broker.AttributeInfo, values []map[string]interface{}) ([]string, error) {
	var tfBrokerObjects []string
	for k, _ := range values {
		tfAttributes := "\t"
		for _, attr := range attributes {
			systemProvisioned := false
			//if attr.Sensitive {
			//	// write-only attributes can't be retrieved so we don't expose them in the datasource
			//	continue
			//}
			//if !attr.Identifying && attr.ReadOnly {
			//	// read-only attributes should only be in the datasource
			//	continue
			//}

			valuesRes := values[k][attr.SempName]
			//if reflect.TypeOf(attr.Default) != nil {
			//	continue
			//}
			switch attr.BaseType {
			case broker.String:
				if reflect.TypeOf(valuesRes) == nil || valuesRes == "" {
					continue
				}
				if strings.Contains(valuesRes.(string), "#") {
					systemProvisioned = true
				}
				val := attr.TerraformName + "\t\t\t\t\t\t=\t\"" + valuesRes.(string) + "\""
				tfAttributes += val
			case broker.Int64:
				if valuesRes == nil {
					continue
				}
				intValue := valuesRes
				val := attr.TerraformName + "\t\t\t\t\t\t=\t" + fmt.Sprintf("%v", intValue)
				tfAttributes += val
			case broker.Bool:
				if valuesRes == nil {
					continue
				}
				boolValue := valuesRes.(bool)
				val := attr.TerraformName + "\t\t\t\t\t\t=\t" + strconv.FormatBool(boolValue)
				tfAttributes += val
			case broker.Struct:
				valueJson, err := json.Marshal(valuesRes)
				if err != nil {
					continue
				}
				val := attr.TerraformName + "\t\t\t\t\t\t=\t" + string(valueJson)
				tfAttributes += val
			}
			if attr.Deprecated && systemProvisioned {
				tfAttributes += "	# Note: This attribute is deprecated and may also be system provisioned."
			} else if attr.Deprecated && !systemProvisioned {
				tfAttributes += "	# Note: This attribute is deprecated."
			} else if !attr.Deprecated && systemProvisioned {
				tfAttributes += "	# Note: This attribute may be system provisioned."
			}
			tfAttributes += "\n\t"
		}
		tfBrokerObjects = append(tfBrokerObjects, tfAttributes)
	}
	return tfBrokerObjects, nil
}

func ConvertToAlphabetic(n int) string {
	return strings.ToLower(string(rune('A' + n)))
}

func LogCLIError(err string) {
	fmt.Fprintf(os.Stdout, "\033[0;31m%s \033[0m\n", err)
}

func LogCLIInfo(info string) {
	fmt.Fprintf(os.Stdout, "\u001B[0m%s \033[0m\n", info)
}
