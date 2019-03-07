package main

import (
	_ "github.com/FadhilAhsan/backend-to-do-list-Golang/routers"
	_ "github.com/go-sql-driver/mysql"
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/mytodolistdb?charset=utf8", 30)

	beego.Run()
}
