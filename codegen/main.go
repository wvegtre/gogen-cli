package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/wvegtre/gogen-cli/convertor"
	"github.com/wvegtre/gogen-cli/convertor/config"
	"github.com/wvegtre/tools-cli/tools/myfile/readfile"
)

func init() {

}

func main() {
	// 1. 读取配置、
	c, err := loadConfigs()
	if err != nil {
		log.Println("loadConfigs failed, ", err)
		os.Exit(0)
	}
	operator := convertor.NewMySQLDBConnect(c)
	p := completeDBConnectConfig(c.Drivers)
	//p := mockForDebug()
	err = validator.New().Struct(&p)
	if err != nil {
		log.Println("GenStructByDBFields Parameter validate failed, ", err)
		os.Exit(0)
	}
	log.Println("input parse succeed, start running...")
	err = operator.GenStructByDBFields(p)
	if err != nil {
		log.Println("GenStructByDBFields Failed, err: ", err)
		os.Exit(0)
	}
	err = operator.GenServiceForDBStruct()
	if err != nil {
		log.Println("GenServiceForDBStruct Failed, err: ", err)
		os.Exit(0)
	}
	err = operator.GenServerFunc()
	if err != nil {
		log.Println("GenServerFunc Failed, err: ", err)
		os.Exit(0)
	}
	log.Println("end running. all filed output to target path.")
}

func loadConfigs() (*config.GenConfig, error) {
	result := readfile.Read("./gen/config/config.json")
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

func completeDBConnectConfig(c config.GenConfigDrivers) convertor.GenDBCodeParameter {
	p := convertor.GenDBCodeParameter{
		UserName: c.Mysql.UserName,
		IP:       c.Mysql.IP,
		Port:     c.Mysql.Port,
		Database: c.Mysql.DB,
		Charset:  c.Mysql.Charset,
	}
	fmt.Println("--- tips: selected mysql as default driver. ---")
	fmt.Print("Please enter your db password: ")
	p.Password = getInputLine()
	if p.Password == "" {
		fmt.Println("!!! exit because password required")
		os.Exit(0)
	}
	return p
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
