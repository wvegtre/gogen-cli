package db

import (
	"strings"

	"echo-shopping/scripts/codegen/tools/dbconverter"
	"echo-shopping/scripts/codegen/tools/serviceconverter"

	"github.com/pkg/errors"
)

type MySQLDBConnect struct {

	// key: group_name, values include table_name and struct_name
	groupMap map[string][]dbconverter.StructContentDetail
}

// GenStructByDBFields 利用第三方 convert 包完成以下几件事
/*
	1. connect mysql client
	2. query mysql table fields
	3. gen struct by table fields
*/
func (c *MySQLDBConnect) GenStructByDBFields(parameter GenParameter) error {
	dc := dbconverter.DBConfig{
		UserName: parameter.UserName,
		Password: parameter.Password,
		IP:       parameter.IP,
		Port:     parameter.Port,
		Database: parameter.Database,
	}
	options := []dbconverter.MySQLConfigOption{
		dbconverter.WithSaveDir("./output/"),
		dbconverter.WithSaveFileDefaultName("model"),
		//dbconverter.WithAllInOneFile(false),
	}
	if parameter.Charset != "" {
		dbconverter.WithCharset(parameter.Charset)
	}
	if len(parameter.TargetTable) > 0 {
		tables := strings.Split(parameter.TargetTable, ",")
		dbconverter.WithTables(tables)
	}
	groupMap, err := dbconverter.NewMySQLConverter(dc, options...).Run()
	if err != nil {
		return errors.WithStack(err)
	}
	// 记录关联关系，后边生成 service 层代码会用到
	c.groupMap = groupMap
	return nil
}

func (c *MySQLDBConnect) GenServiceForDBStruct() error {
	options := []serviceconverter.ServiceGenConfigOption{
		serviceconverter.WithSaveDir("./output/"),
		serviceconverter.WithSaveFileDefaultName("service"),
		//dbconverter.WithAllInOneFile(false),
	}
	err := serviceconverter.NewServiceConverter(options...).Run(c.groupMap)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
