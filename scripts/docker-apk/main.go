package main

import (
	_ "light-apk/routers"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	dbuser := beego.AppConfig.String("dbuser")
	dbpwd := beego.AppConfig.String("dbpwd")
	dbname := beego.AppConfig.String("dbname")
	dbhost := beego.AppConfig.String("dbhost")
	maxIdle, _ := beego.AppConfig.Int("maxidle")
	maxConn, _ := beego.AppConfig.Int("maxconn")

	orm.RegisterDriver("mysql", orm.DRMySQL)

	conn := dbuser + ":" + dbpwd + "@tcp(" + dbhost + ")/" + dbname + "?charset=utf8&loc=Local"
	err := orm.RegisterDataBase("default", "mysql", conn, maxIdle, maxConn)
	if err != nil {
		beego.Error(err.Error)
	}
	orm.Debug = true

	//自动建表
	name := "default"                         //数据库别名
	force := false                            //不强制建数据库
	verbose := true                           //打印建表过程
	err = orm.RunSyncdb(name, force, verbose) //建表

	if err != nil {
		beego.Error(err)
	}
	beego.Debug("初始化")

}

func main() {
	beego.SetStaticPath("/download", "static/files")

	beego.SetLogger("console", `{"level":7}`) // debug 级别以上的信息会被输出
	beego.SetLogger("file", `{"filename":"logs/`+beego.AppConfig.String("appname")+`.log"}`)

	beego.Run()
}
