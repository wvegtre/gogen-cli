package convertor

import (
	"fmt"

	"gogen-cli/aigo/config"
	"gogen-cli/aigo/model"

	"github.com/pkg/errors"
)

type MySQLConvertor struct {
	projectName string
	tableName   string
	fields      model.TableFields
	*config.TemplatePaths
}

func NewMySQLConvertor(projectName, tableName string, fields model.TableFields, paths *config.TemplatePaths) *MySQLConvertor {
	return &MySQLConvertor{
		projectName:   projectName,
		tableName:     tableName,
		fields:        fields,
		TemplatePaths: paths,
	}
}

func (c MySQLConvertor) GenAPIParameter(tplFile, writeFile string) error {
	var listArgsRow []string
	var listArgsConvert []string
	for _, v := range c.fields {
		structName := SnakeToCamel(v.FieldName)
		structFiled := fmt.Sprintf("%s %s `json:%s` ",
			structName, fieldTypeMapping(v.FieldType), fmt.Sprintf(`"%s"`, v.FieldName))
		if v.FieldDescription != "" {
			structFiled += `// ` + v.FieldDescription
		}
		listArgsRow = append(listArgsRow, structFiled)
	}
	err := ParseTPLAndWrite(c.projectName, tplFile, writeFile, struct {
		TableName       string
		ListArgsRow     []string
		ListArgsConvert []string
	}{
		TableName:       SnakeToCamel(c.tableName),
		ListArgsConvert: listArgsConvert,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c MySQLConvertor) GenAPIRouter(tplFile, writeFile string) error {
	err := ParseTPLAndWrite(c.projectName, tplFile, writeFile, struct {
		RouterPrefix string
		TableName    string
		TableNameLow string
	}{
		RouterPrefix: c.tableName,
		TableName:    SnakeToCamel(c.tableName),
		TableNameLow: LowFirstLetter(SnakeToCamel(c.tableName)),
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c MySQLConvertor) GenServiceFunc(tplFile, writeFile string) error {
	var rowCovert []string
	var listArgsRow []string
	var createArgsRow []string
	var createArgsConvert []string
	var updateArgsRow []string
	var updateArgsConvert []string
	listArgMap := make(map[string]string)
	for _, v := range c.fields {
		structName := SnakeToCamel(v.FieldName)
		structFiled := fmt.Sprintf("%s %s ", structName, fieldTypeMapping(v.FieldType))
		structFiled += "`json:" + fmt.Sprintf(`"%s"`, v.FieldName) + "`"
		if v.FieldDescription != "" {
			structFiled += `// ` + v.FieldDescription
		}
		listArgsRow = append(listArgsRow, structFiled)
		listArgMap[v.FieldName] = structName
		rowCovert = append(rowCovert, structName)

		if !c.notAllowWriteWhenCreated(v.FieldName) {
			createArgsRow = append(createArgsRow, structFiled)
			createArgsConvert = append(createArgsConvert, structName)
		}

		if !c.notAllowWriteWhenUpdate(v.FieldName) {
			updateArgsRow = append(updateArgsRow, structFiled)
			updateArgsConvert = append(updateArgsConvert, structName)
		}
	}

	err := ParseTPLAndWrite(c.projectName, tplFile, writeFile, struct {
		TableName         string
		ListArgsRow       []string
		ListArgMap        map[string]string
		CreateArgsRow     []string
		CreateArgsConvert []string
		UpdateArgsRow     []string
		UpdateArgsConvert []string
		ModelFields       []string
		ModelConvert      []string
	}{
		TableName:         SnakeToCamel(c.tableName),
		ListArgsRow:       listArgsRow,
		ListArgMap:        listArgMap,
		CreateArgsRow:     createArgsRow,
		CreateArgsConvert: createArgsConvert,
		UpdateArgsRow:     updateArgsRow,
		UpdateArgsConvert: updateArgsConvert,
		ModelFields:       listArgsRow,
		ModelConvert:      rowCovert,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c MySQLConvertor) GenDBModel(tplFile, writeFile string) error {
	var structFields []string
	for _, v := range c.fields {
		structName := SnakeToCamel(v.FieldName)
		structFiled := fmt.Sprintf("%s %s `json:%s` ",
			structName, fieldTypeMapping(v.FieldType), fmt.Sprintf(`"%s"`, v.FieldName))
		if v.FieldDescription != "" {
			structFiled += `// ` + v.FieldDescription
		}
		structFields = append(structFields, structFiled)
	}
	err := ParseTPLAndWrite(c.projectName, tplFile, writeFile, struct {
		TableName    string
		DBTableName  string
		StructFields []string
	}{
		TableName:    SnakeToCamel(c.tableName),
		DBTableName:  c.tableName,
		StructFields: structFields,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c MySQLConvertor) notAllowWriteWhenCreated(field string) bool {
	return field == "id" || c.notAllowWriteWhenUpdate(field)
}

func (c MySQLConvertor) notAllowWriteWhenUpdate(field string) bool {
	return field == "created_at" || field == "updated_at"
}
