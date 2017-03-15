package controllers

import (
	"github.com/astaxie/beego"
)

type MessageController struct {
	beego.Controller
}

func (this *MessageController) Get() {
	this.Data["username"] = "xiaorun"
	this.TplName = "index.html"
}
