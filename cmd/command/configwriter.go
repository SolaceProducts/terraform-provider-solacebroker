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
	"embed"
	"os"
	"strings"
	"text/template"
)

var (
	//go:embed templates
	templatefiles embed.FS
)
var terraformTemplate *template.Template

func init() {
	var err error
	terraformTemplateString, _ := templatefiles.ReadFile("templates/terraform.template")
	terraformTemplate, err = template.New("Object Template").Funcs(template.FuncMap{
		"splitHCLResourceName": func(value string) []string {
			return strings.Split(value, " ")
		},
		"readHCLResourceName": func(slice []string, index int) string {
			return slice[index]
		},
	}).Parse(string(terraformTemplateString))
	if err != nil {
		panic(err)
	}
}

func GenerateTerraformFile(terraformObjectInfo *ObjectInfo) error {
	var codeStream bytes.Buffer
	err := terraformTemplate.Execute(&codeStream, terraformObjectInfo)
	if err != nil {
		LogCLIError("\nError: Templating error : " + err.Error() + "\n\n")
		os.Exit(1)
	}
	return os.WriteFile(terraformObjectInfo.FileName, codeStream.Bytes(), 0664)
}
