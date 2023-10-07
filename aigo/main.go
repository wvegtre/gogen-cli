package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"gogen-cli/aigo/config"
	"gogen-cli/aigo/convertor"
	"gogen-cli/aigo/detector"

	"github.com/go-playground/validator/v10"
	"github.com/logrusorgru/aurora"
	"github.com/pkg/errors"
	"github.com/wvegtre/tools-cli/tools/myfile/readfile"
)

func main() {
	ctx := context.Background()
	// load json config
	c, err := loadConfigs()
	if err != nil {
		log.Println(aurora.Red("GetTTSProduct error"), err)
		os.Exit(0)
	}
	// complete db password
	dbAuth := completeDBConnectConfig(c.Drivers.Mysql)
	// connect to db and get table fields
	err = validator.New().Struct(dbAuth)
	if err != nil {
		log.Println(aurora.Red("GenStructByDBFields Parameter validate failed"), err)
		os.Exit(0)
	}
	mySQLDetector := detector.NewMysqlDetector(dbAuth)
	targetTables := c.Tables
	if len(c.Tables) == 0 {
		targetTables, err = mySQLDetector.GetAllTables()
		if err != nil {
			log.Println("GetAllTables failed, ", err)
			os.Exit(0)
		}
	}
	err = parseAndWriteInitAPIRouter(c, targetTables)
	if err != nil {
		log.Println(aurora.Red("parseAndWriteInitAPIRouter failed"), err)
		os.Exit(0)
	}
	for _, tableName := range targetTables {
		tableFields, err := mySQLDetector.GetTableFields(ctx, tableName)
		if err != nil {
			log.Println(aurora.Red("GetTableFields failed"), err)
			os.Exit(0)
		}
		mySQLConvertor := convertor.NewMySQLConvertor(c.Output.ProjectName, tableName, tableFields, c.TemplatePaths)
		for _, v := range defaultParseAndWrite(c, mySQLConvertor, tableName) {
			err = v.fn(v.tplFile, v.writeFile)
			if err != nil {
				log.Println(aurora.Red("defaultParseAndWrite failed"), err)
				os.Exit(0)
			}
		}
		log.Println(aurora.Green(fmt.Sprintf("%s table generate success", tableName)))
	}

	// 无用文件删除
	for _, v := range []string{
		".DS_Store", ".tpl",
	} {
		err = convertor.DeleteFileWithSuffixName(c.Output.Dir+c.Output.ProjectName, v)
		if err != nil {
			log.Println(aurora.Red("DeleteFileWithSuffixName failed"), err)
			os.Exit(0)
		}
	}

	// 项目关键词批量替换
	err = convertor.ReplaceFileDir(c.Output.Dir+c.Output.ProjectName, c.Output.ProjectName)
	if err != nil {
		log.Println(aurora.Red("ReplaceFileDir failed"), err)
		os.Exit(0)
	}
	log.Println(aurora.Green("finish...."))
}

func loadConfigs() (*config.GenConfig, error) {
	result := readfile.Read("./config/config.json")
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "read failed")
	}
	c := &config.GenConfig{}
	err := json.Unmarshal(result.JSONDetail, c)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return c, nil
}

func completeDBConnectConfig(c *config.DBAuth) *config.DBAuth {
	//fmt.Println("--- tips: selected mysql as default driver. ---")
	//fmt.Print("Please enter your db password: ")
	c.Password = os.Getenv("db_password")
	if c.Password == "" {
		// TODO 本地调试
		//fmt.Println("!!! exit because password required")
		//os.Exit(0)
		c.Password = "123456"
	}
	return c
}

// 读取命令行输入的内容，如果输入空白字符串，则返回默认值
func getInputLineWithDefault(key string) string {
	line := getInputLine()
	line = strings.TrimSpace(line)
	if line == "" {
		line = os.Getenv(key)
		fmt.Println("--- Find input empty, get from os env. key: ", key, ",value: ", line, " ---")
	}
	return line
}

// 读取命令行输入的内容，
func getInputLine() string {
	inputReader := bufio.NewReader(os.Stdin)
	line, _, _ := inputReader.ReadLine()
	return string(line)
}

type parseAndWrite struct {
	tplFile   string
	writeFile string
	fn        func(tplFile, writeFile string) error
}

func parseAndWriteInitAPIRouter(c *config.GenConfig, tableNames []string) error {
	// 1. api_route.tpl
	tplFile := c.TemplatePaths.Basic + c.TemplatePaths.InitAPIRouter + "/api_route.tpl"
	writeFile := c.Output.Dir + c.Output.ProjectName + c.TemplatePaths.InitAPIRouter + "/router.go"
	err := convertor.ParseTPLAndWrite(c.Output.ProjectName, tplFile, writeFile, struct {
		APIPrefix string
	}{
		APIPrefix: strings.ReplaceAll(c.Output.ProjectName, "-", "_"),
	})
	if err != nil {
		return errors.WithStack(err)
	}
	// 2. add_v1_api_router.tpl
	tplFile = c.TemplatePaths.Basic + c.TemplatePaths.AddV1APIRouter + "/add_api_route.tpl"
	writeFile = c.Output.Dir + c.Output.ProjectName + c.TemplatePaths.AddV1APIRouter + "/router.go"
	var routers []string
	for _, v := range tableNames {
		routers = append(routers, convertor.SnakeToCamel(v))
	}
	err = convertor.ParseTPLAndWrite(c.Output.ProjectName, tplFile, writeFile, struct {
		TableNames []string
	}{
		TableNames: routers,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	// 3. app_main.tpl
	tplFile = c.TemplatePaths.Basic + "/cmd/app/app_main.tpl"
	writeFile = c.Output.Dir + c.Output.ProjectName + "/cmd/app/main.go"
	err = convertor.ParseTPLAndWrite(c.Output.ProjectName, tplFile, writeFile, struct{}{})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func defaultParseAndWrite(c *config.GenConfig, mySQLConvertor *convertor.MySQLConvertor, tableName string) []parseAndWrite {
	return []parseAndWrite{
		{
			tplFile:   c.TemplatePaths.Basic + c.TemplatePaths.APIParameter + "/api_args.tpl",
			writeFile: c.Output.Dir + c.Output.ProjectName + c.TemplatePaths.APIParameter + fmt.Sprintf("/%s.go", tableName),
			fn:        mySQLConvertor.GenAPIParameter,
		},
		{
			tplFile:   c.TemplatePaths.Basic + c.TemplatePaths.APIRouters + "/api_router.tpl",
			writeFile: c.Output.Dir + c.Output.ProjectName + c.TemplatePaths.APIRouters + fmt.Sprintf("/%s.go", tableName),
			fn:        mySQLConvertor.GenAPIRouter,
		},
		{
			tplFile:   c.TemplatePaths.Basic + c.TemplatePaths.InternalService + "/service.tpl",
			writeFile: c.Output.Dir + c.Output.ProjectName + c.TemplatePaths.InternalService + fmt.Sprintf("/%s.go", tableName),
			fn:        mySQLConvertor.GenServiceFunc,
		},
		{
			tplFile:   c.TemplatePaths.Basic + c.TemplatePaths.InternalDatabase + "/db_model.tpl",
			writeFile: c.Output.Dir + c.Output.ProjectName + c.TemplatePaths.InternalDatabase + fmt.Sprintf("/%s.go", tableName),
			fn:        mySQLConvertor.GenDBModel,
		},
	}
}
