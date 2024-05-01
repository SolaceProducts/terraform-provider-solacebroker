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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
	"unicode"
)

type Color string

const (
	Reset Color = "\033[0m"
	Red   Color = "\033[31m"
)

var idStartYes = []*unicode.RangeTable{
	unicode.L,
	// unicode.Nl, // not included as it is not a valid starting character for an identifier
	unicode.Other_ID_Start,
	// include Hyphen and Underscore
	{
		R16: []unicode.Range16{
			{uint16('-'), uint16('-'), 1},
			{uint16('_'), uint16('_'), 1},
		},
	},
}

// This code defines the idContinueYes slice, which contains Unicode range tables for valid continuation characters in an identifier.
// The idContinueYes slice includes categories such as letter, number, mark, punctuation, and other valid continuation characters.
var idContinueYes = []*unicode.RangeTable{
	unicode.L,
	unicode.Nl,
	unicode.Other_ID_Start,
	unicode.Mn,
	unicode.Mc,
	unicode.Nd,
	unicode.Pc,
	unicode.Other_ID_Continue,
}

var idNo = []*unicode.RangeTable{
	unicode.Pattern_Syntax,
	unicode.Pattern_White_Space,
}

func StringWithDefaultFromEnv(name string, isMandatory bool, fallback string) string {
	envValue := os.Getenv("SOLACEBROKER_" + strings.ToUpper(name))
	if isMandatory && len(envValue) == 0 {
		ExitWithError("SOLACEBROKER_" + strings.ToUpper(name) + " is mandatory but not available")
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

// Workaround filter for system provisioned attributes
func isSystemProvisionedAttribute(attribute string) bool {
	systemProvisioned := strings.HasPrefix(attribute, "#") && attribute != "#DEAD_MSG_QUEUE"
	return systemProvisioned
}

func LogCLIError(err string) {
	_, _ = fmt.Fprintf(os.Stdout, "%s %s %s\n", Red, err, Reset)
}

func LogCLIInfo(info string) {
	_, _ = fmt.Fprintf(os.Stdout, "\n%s %s %s", Reset, info, Reset)
}

func ExitWithError(err string) {
	LogCLIError(err)
	os.Exit(1)
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

func resourcesToFormattedHCL(brokerResources []map[string]ResourceConfig) []map[string]string {
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

func SanitizeHclStringValue(value string) string {
	b, _ := json.Marshal(value)
	s := string(b)
	output := s[1 : len(s)-1]
	output = strings.ReplaceAll(output, "${", "$${")
	output = strings.ReplaceAll(output, "%{", "%%{")
	return output
}

func isStartRune(r rune) bool {
	return r == '-' || unicode.In(r, idStartYes...) && !unicode.In(r, idNo...)
}

func isContinueRune(r rune) bool {
	return r == '-' || unicode.In(r, idContinueYes...) && !unicode.In(r, idNo...)
}

// A valid Terraform identifier must satisfy the following conditions:
// - It must not be an empty string.
// - The first character must be a valid starting character for an identifier.
// - All subsequent characters must be valid continuation characters for an identifier.
func IsValidTerraformIdentifier(s string) bool {
	if s == "" {
		return false
	}
	runes := []rune(s)
	if !isStartRune(runes[0]) {
		return false
	}
	for _, r := range runes[1:] {
		if !isContinueRune(r) {
			return false
		}
	}
	return true
}

// makeValidForTerraformIdentifier replaces invalid characters in a string with hyphens ('-').
// It takes a string as input and iterates over each rune in the string.
// If the rune is not a valid continuation character or is in the idNo slice, it is replaced with a hyphen.
// The function returns the modified string with hyphens replacing the invalid characters.
func makeValidForTerraformIdentifier(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if i == 0 && !unicode.In(r, idStartYes...) || unicode.In(r, idNo...) {
			runes[i] = '-'
		} else {
			if !unicode.In(r, idContinueYes...) || unicode.In(r, idNo...) {
				runes[i] = '-'
			}
		}
	}
	return string(runes)
}
