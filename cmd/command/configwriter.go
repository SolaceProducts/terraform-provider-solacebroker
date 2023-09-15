package terraform

import (
	"bytes"
	"embed"
	"os"
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
	terraformTemplate, err = template.New("Object Template").Parse(string(terraformTemplateString))
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
