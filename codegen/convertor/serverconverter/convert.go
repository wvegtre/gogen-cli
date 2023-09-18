package serverconverter

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/wvegtre/gogen-cli/convertor/dbconverter"
)

type ServerConverter struct {
	config *serverConverterConfig
}

func NewServerConverter(ops ...ServerGenConfigOption) *ServerConverter {
	defaultConfig := newDefaultConfig()
	for _, op := range ops {
		op(defaultConfig)
	}
	return &ServerConverter{
		config: defaultConfig,
	}
}

type serverConverterConfig struct {
	// some configs for write file
	fileConfig
}

type fileConfig struct {
	SaveDir         string
	SaveProjectName string
	SavePackageName string
}

func newDefaultConfig() *serverConverterConfig {
	c := &serverConverterConfig{}
	return c
}

func (c *ServerConverter) genCompletePathPrefix() string {
	basePath := c.config.SaveDir
	basePath = c.appendSuffix(basePath)
	basePath += c.config.SaveProjectName
	basePath = c.appendSuffix(basePath)
	basePath += c.config.SavePackageName
	basePath = c.appendSuffix(basePath)
	return basePath
}

func (c *ServerConverter) appendSuffix(basePath string) string {
	if !strings.HasSuffix(basePath, "/") {
		basePath += "/"
	}
	return basePath
}

func (c *ServerConverter) Run(detail *dbconverter.OutputDetail) error {
	for name, values := range detail.GroupMap {
		basePath := c.genCompletePathPrefix() + name
		log.Println("create file folder, path:  ", basePath)
		err := os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			return errors.WithStack(err)
		}
		// write args.go
		outputFileContent, err := c.buildArgsFileContent(name, values, detail.TableMap)
		if err != nil {
			return errors.WithStack(err)
		}
		savePath := basePath + "/args.go"
		err = c.writeToFile(savePath, outputFileContent)
		if err != nil {
			return errors.WithStack(err)
		}
		// write service.go
		outputFileContent, err = c.buildServiceFileContent(name, values, detail.TableMap)
		if err != nil {
			return errors.WithStack(err)
		}
		savePath = basePath + "/service.go"
		err = c.writeToFile(savePath, outputFileContent)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (c *ServerConverter) buildArgsFileContent(
	name string, values []dbconverter.StructContentDetail,
	content map[string]*dbconverter.TableContentDetail) (string, error) {
	outputFile := "package %s\nimport(%s\n)\n%s"
	importContent := "\n\"github.com/wvegtre/gogen-cli/output/%s/internal/app/database/%s\""
	var outputContent string
	for _, v := range values {
		params := make(map[string]interface{})
		// StructName 就是通过 table name 处理后拼接上 "Model"，这里重新把 "Model" 移除就好了
		params["service_prefix"] = strings.Replace(v.StructName, "Model", "", -1)
		params["model"] = v.StructName
		params["group"] = name

		detail, ok := content[v.TableName]
		if !ok {
			return "", errors.New("not match content by table " + v.TableName)
		}
		params["args_slice"] = detail.FieldRow
		params["args_map"] = detail.FieldMap

		content, err := c.parseTemplate(params, "convert_args", "./gen/serverconverter/args.tpl")
		if err != nil {
			return "", errors.WithStack(err)
		}
		outputContent += content + "\n"
	}
	outputFile = fmt.Sprintf(outputFile, name, fmt.Sprintf(importContent, c.config.SaveProjectName, name), outputContent)
	return outputFile, nil
}

func (c *ServerConverter) buildServiceFileContent(
	name string, values []dbconverter.StructContentDetail,
	content map[string]*dbconverter.TableContentDetail) (string, error) {
	outputFile := "package %s\n\nimport(%s)\n%s"
	var importContent string
	importContent += "\n\"context\""
	importContent += "\n\"github.com/wvegtre/gogen-cli/output/%s/internal/app/database\""
	importContent += "\n\"github.com/wvegtre/gogen-cli/output/%s/internal/app/database/%s\""
	importContent += "\n\"github.com/pkg/errors\""
	var outputContent string
	for _, v := range values {
		params := make(map[string]interface{})
		// StructName 就是通过 table name 处理后拼接上 "Model"，这里重新把 "Model" 移除就好了
		params["service_prefix"] = strings.Replace(v.StructName, "Model", "", -1)
		params["model"] = v.StructName
		params["group"] = name

		content, err := c.parseTemplate(params, "convert_service", "./gen/serverconverter/service.tpl")
		if err != nil {
			return "", errors.WithStack(err)
		}
		outputContent += content + "\n"
	}
	outputFile = fmt.Sprintf(outputFile, name, fmt.Sprintf(
		importContent, c.config.SaveProjectName, c.config.SaveProjectName, name,
	), outputContent)
	return outputFile, nil
}

func (c *ServerConverter) parseTemplate(params map[string]interface{}, name, filePath string) (string, error) {
	t, err := template.New(name).ParseGlob(filePath)
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

func (c *ServerConverter) writeToFile(filePath string, content string) error {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println("Can not write file, err: ", err)
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
