package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	_ "test/models/chapter05"
	_ "test/routers"
)

func init() {
	username, _ := beego.AppConfig.String("username")
	password, _ := beego.AppConfig.String("password")
	host, _ := beego.AppConfig.String("host")
	port, _ := beego.AppConfig.String("port")
	database, _ := beego.AppConfig.String("database")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", username, password, host, port, database)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", datasource)

	fmt.Println("连接数据库")

	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		panic(err)
	}

	orm.Debug = true
}

func main() {
	//过滤器
	//beego.InsertFilter("/cms/*", beego.BeforeRouter, chapter04.MyFilter)
	orm.RunCommand()

	//logs
	//logs.SetLogger("console")
	//把日志写进文件
	//logs.SetLogger(logs.AdapterFile, `{"filename":"logs/beego_project.log"}`)
	//多日志写入
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/beego_project.log", "separate":["error", "info"]}`)

	beego.Run()
}
