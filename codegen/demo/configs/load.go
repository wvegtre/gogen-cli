package configs

import (
	"os"

	"github.com/spf13/viper"
)

func LoadConf() Conf {
	c := Conf{}
	v := loadConfig()
	if err := v.Unmarshal(&c); err != nil {
		panic(err)
	}
	return c
}

func loadConfig() *viper.Viper {
	v := viper.GetViper()
	v.SetConfigName(getConfigFileName()) // 配置文件名称(无扩展名)
	v.SetConfigType("toml")              // 如果配置文件的名称中没有扩展名，则需要配置此项
	v.AddConfigPath("../../configs")     // 还可以在工作目录中查找配置
	v.AddConfigPath("./configs")         // 还可以在工作目录中查找配置
	err := v.ReadInConfig()              // 查找并读取配置文件
	if err != nil {                      // 处理读取配置文件的错误
		panic(err)
	}
	return v
}

func getConfigFileName() string {
	env := os.Getenv("env")
	if env == "" {
		env = "local"
	}
	return "config-" + env
}
