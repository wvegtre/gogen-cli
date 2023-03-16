package convertor

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/wvegtre/gogen-cli/convertor/config"
	"github.com/wvegtre/gogen-cli/convertor/dbconverter"
	"github.com/wvegtre/gogen-cli/convertor/serverconverter"
	"github.com/wvegtre/gogen-cli/convertor/serviceconverter"
)

type MySQLDBConnect struct {
	Config *config.GenConfig
	// key: group_name, values include table_name and struct_name
	Output *dbconverter.OutputDetail
}

// GenStructByDBFields 利用第三方 convert 包完成以下几件事
/*
	1. connect mysql client
	2. query mysql table fields
	3. gen struct by table fields
*/
func (c *MySQLDBConnect) GenStructByDBFields(parameter GenDBCodeParameter) error {
	dc := dbconverter.DBConfig{
		UserName: parameter.UserName,
		Password: parameter.Password,
		IP:       parameter.IP,
		Port:     parameter.Port,
		Database: parameter.Database,
	}
	options := []dbconverter.MySQLConfigOption{
		dbconverter.WithSaveDir(c.Config.Output.Dir),
		dbconverter.WithSaveProjectName(c.Config.Output.ProjectName),
		dbconverter.WithSavePackageName("internal/app/database"),
	}
	if parameter.Charset != "" {
		dbconverter.WithCharset(parameter.Charset)
	}
	if len(c.Config.Drivers.Mysql.Tables) > 0 {
		tables := strings.Split(c.Config.Drivers.Mysql.Tables, ",")
		dbconverter.WithTables(tables)
	}
	detail, err := dbconverter.NewMySQLConverter(dc, options...).Run()
	if err != nil {
		return errors.WithStack(err)
	}
	// 记录关联关系，后边生成 service 层代码会用到
	c.Output = detail
	return nil
}

func (c *MySQLDBConnect) GenServiceForDBStruct() error {
	options := []serviceconverter.ServiceGenConfigOption{
		serviceconverter.WithSaveDir(c.Config.Output.Dir),
		serviceconverter.WithSaveProjectName(c.Config.Output.ProjectName),
		serviceconverter.WithSavePackageName("internal/app/database/"),
	}
	err := serviceconverter.NewServiceConverter(options...).Run(c.Output.GroupMap)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *MySQLDBConnect) GenServerFunc() error {
	options := []serverconverter.ServerGenConfigOption{
		serverconverter.WithSaveDir(c.Config.Output.Dir),
		serverconverter.WithSaveProjectName(c.Config.Output.ProjectName),
		serverconverter.WithSavePackageName("internal/app/server/"),
	}
	err := serverconverter.NewServerConverter(options...).Run(c.Output)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
