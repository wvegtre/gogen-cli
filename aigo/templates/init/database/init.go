package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"gen-templates/configs"

	_ "github.com/go-sql-driver/mysql"

	gsql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGROMClientForMySQL(conf configs.Conf) *gorm.DB {
	mySqlDb := newMySQLClient(conf)
	gormDB, err := gorm.Open(gsql.New(gsql.Config{
		Conn: mySqlDb,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return gormDB
}

func newMySQLClient(conf configs.Conf) *sql.DB {
	pwd := os.Getenv("mysql_password")
	client, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		conf.Database.MySQL.User, pwd,
		conf.Database.MySQL.Host, conf.Database.MySQL.Port,
		conf.Database.MySQL.DB,
	))
	if err != nil {
		panic(fmt.Errorf("mysql connect failed. %w", err))
	}
	//设置数据库最大连接数
	client.SetConnMaxLifetime(time.Duration(conf.Database.MySQL.MaxLifetimeSecond) * time.Second)
	//设置上数据库最大闲置连接数
	client.SetMaxIdleConns(conf.Database.MySQL.MaxIdle)
	//验证连接
	if err := client.Ping(); err != nil {
		panic(fmt.Errorf("mysql ping failed. %w", err))
	}
	fmt.Println("mysql connect success.")
	return client
}
