package serviceconverter

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"echo-shopping/scripts/codegen/tools/dbconverter"

	"github.com/pkg/errors"
)

type ServiceConverter struct {
	config *serviceConverterConfig
}

func NewServiceConverter(ops ...ServiceGenConfigOption) *ServiceConverter {
	defaultConfig := newDefaultConfig()
	for _, op := range ops {
		op(defaultConfig)
	}
	return &ServiceConverter{
		config: defaultConfig,
	}
}

type serviceConverterConfig struct {
	// some configs for write file
	fileConfig
}

type fileConfig struct {
	//AllInOneFile        bool
	SaveDir string
	//SaveFilePrefix      string
	SaveFileDefaultName string // default model.go, but invalid when AllInOneFile=false
}

func newDefaultConfig() *serviceConverterConfig {
	c := &serviceConverterConfig{}
	c.fileConfig = fileConfig{
		SaveDir:             "./",
		SaveFileDefaultName: "service",
	}
	return c
}

func (c *ServiceConverter) Run(groupMap map[string][]dbconverter.StructContentDetail) error {
	for name, values := range groupMap {
		outputFileContent, err := c.buildFileContent(name, values)
		if err != nil {
			return errors.WithStack(err)
		}
		savePath := c.config.SaveDir + name + "/" + c.config.SaveFileDefaultName + ".go"
		err = c.writeToFile(savePath, outputFileContent)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (c *ServiceConverter) buildFileContent(name string, values []dbconverter.StructContentDetail) (string, error) {
	outputFile := `package %s
		
		%s
		`
	var outputContent string
	for _, v := range values {
		params := make(map[string]string)
		// StructName 就是通过 table name 处理后拼接上 "Model"，这里重新把 "Model" 移除就好了
		params["service_prefix"] = strings.Replace(v.StructName, "Model", "", -1)
		params["model"] = v.StructName
		content, err := c.parseTemplate(params)
		if err != nil {
			return "", errors.WithStack(err)
		}
		outputContent += content + "\n"
	}
	outputFile = fmt.Sprintf(outputFile, name, outputContent)
	return outputFile, nil
}

func (c *ServiceConverter) parseTemplate(params map[string]string) (string, error) {
	t, err := template.New("convert").ParseGlob("./tools/serviceconverter/service.tpl")
	if err != nil {
		return "", errors.WithStack(err)
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, params)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return buf.String(), nil
}

func (c *ServiceConverter) writeToFile(filePath string, content string) error {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println("Can not write file")
		return errors.WithStack(err)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return errors.WithStack(err)
	}
	log.Println("write to field succeed. path: ", filePath)
	return nil
}
