package main

import (
	_ "message/docs"
	_ "message/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	// 设置静态资源目录
	beego.SetStaticPath("/dist", "views/dist")
	beego.Run()
}
