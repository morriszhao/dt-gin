package helper

import (
	"github.com/spf13/viper"
	"log"
)

var ViperConfig *viper.Viper

func InitConfig() {
	ViperConfig = viper.New()
	ViperConfig.AddConfigPath("./config") //添加配置文件搜索路径    可以添加多个   和php一样  以入口文件为 根路径
	ViperConfig.SetConfigType("yaml")     //配置文件类型
	ViperConfig.SetConfigName("app.yaml") //配置文件名

	if err := ViperConfig.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("use config file -> %s\n", ViperConfig.ConfigFileUsed())

}
