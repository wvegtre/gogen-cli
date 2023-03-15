package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"echo-shopping/scripts/codegen/db"

	"github.com/go-playground/validator/v10"
)

func init() {

}

func main() {
	// 1. 连接数据库
	p := inputForDBConnectParameter()
	//p := mockForDebug()
	err := validator.New().Struct(&p)
	if err != nil {
		log.Println("GenStructByDBFields Parameter validate failed, ", err)
		os.Exit(0)
	}
	log.Println("input parse succeed, start running...")
	err = p.Operator.GenStructByDBFields(p)
	if err != nil {
		log.Println("GenStructByDBFields Failed, err: ", err)
		os.Exit(0)
	}
	err = p.Operator.GenServiceForDBStruct()
	if err != nil {
		log.Println("GenServiceForDBStruct Failed, err: ", err)
		os.Exit(0)
	}
	log.Println("end running. all filed output to target path.")
}

func inputForDBConnectParameter() db.GenParameter {
	p := db.GenParameter{}
	fmt.Print("Please enter your db driver: ")
	driver := getInputLine()
	// TODO 改成可选项
	switch driver {
	case "mysql":
		p.Operator = db.NewMySQLDBConnect()
	default:
		fmt.Println("--- tips: selected mysql as default driver. ---")
		p.Operator = db.NewMySQLDBConnect()
	}
	fmt.Print("Please enter your db user: ")
	p.UserName = getInputLineWithDefault("mysql_user")
	fmt.Print("Please enter your db password: ")
	p.Password = getInputLine()
	if p.Password == "" {
		fmt.Println("!!! exit because password required")
		os.Exit(0)
	}
	fmt.Print("Please enter your db ip: ")
	p.IP = getInputLine()
	fmt.Print("Please enter your db port: ")
	p.Port = getInputLine()
	if p.IP == "" || p.Port == "" {
		domain := os.Getenv("mysql_domain")
		arr := strings.Split(domain, ":")
		p.IP = arr[0]
		p.Port = arr[1]
		fmt.Println("--- Find db ip or port empty, get from os env. ip:port -> ", p.IP, ":", p.Port, " ---")
	}
	fmt.Print("Please enter your database: ")
	p.Database = getInputLineWithDefault("mysql_db")
	// TODO 改成可选项
	fmt.Println("Please enter target table")
	p.TargetTable = getInputLine()
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
