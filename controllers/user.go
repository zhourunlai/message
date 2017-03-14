package controllers

import (
	"message/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title signin
// @Description user signin
// @Param	username		query 	string	true		"The username for signin"
// @Param	password		query 	string	true		"The password for signin"
// @Success 200 {string} signin success
// @Failure 401 signin failed
// @router /signin [get]
func (u *UserController) Signin() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Signin(username, password) {
		u.Data["json"] = "signin success"
	} else {
		u.Data["json"] = "signin failed"
	}
	u.ServeJSON()
}

// @Title signout
// @Description Logs out current logged in user session
// @Success 200 {string} signout success
// @router /signout [get]
func (u *UserController) Signout() {
	u.Data["json"] = "signout success"
	u.ServeJSON()
}
