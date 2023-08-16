package models

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

const (
	driveName = "mysql"
	dsName    = "root:123456@(127.0.0.1:3306)/tech-chat?charset=utf8"
	showSQL   = true
	maxCon    = 10
	NoError   = "no" //没有错误
)

var DbEngin *xorm.Engine

// InitMysql  数据库链接
func InitMysql() {
	err := errors.New(NoError)
	DbEngin, err = xorm.NewEngine(driveName, dsName)
	if nil != err && NoError != err.Error() {
		log.Fatal(err.Error())
	}

	//是否显示SQL语句
	DbEngin.ShowSQL(showSQL)
	//数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(maxCon)

	//自动User
	DbEngin.Sync2(new(UserModel))

	//测试链接
	err = DbEngin.Ping()
	if err != nil {
		log.Fatal("数据库连接失败：", err.Error())
	}

}
