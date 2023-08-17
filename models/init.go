package models

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"morris/im/helper"
)

type mysqlConfig struct {
	driverName string
	dsn        string
	showSql    bool
	maxConn    int
}

var DbEngin *xorm.Engine

// InitMysql  数据库链接
func InitMysql() {
	mysqlConfig := initConfig()

	err := errors.New("")
	DbEngin, err = xorm.NewEngine(mysqlConfig.driverName, mysqlConfig.dsn)
	if nil != err && "" != err.Error() {
		log.Fatal(err.Error())
	}

	//是否显示SQL语句
	DbEngin.ShowSQL(mysqlConfig.showSql)
	//数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(mysqlConfig.maxConn)

	//自动User
	DbEngin.Sync2(new(UserModel))

	//测试链接
	err = DbEngin.Ping()
	if err != nil {
		log.Fatal("数据库连接失败：", err.Error())
	}

}

func initConfig() mysqlConfig {

	subViper := helper.ViperConfig.Sub("db.mysql")
	config := mysqlConfig{
		driverName: "mysql",
		dsn:        subViper.GetString("dsn"),
		showSql:    subViper.GetBool("showSql"),
		maxConn:    subViper.GetInt("maxConn"),
	}

	if config.dsn == "" {
		log.Fatal("请配置数据库dsn")
	}
	if config.maxConn == 0 {
		config.maxConn = 10
	}

	return config

}
