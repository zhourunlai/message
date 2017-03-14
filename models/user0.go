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
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} signin success
// @Failure 403 user not exist
// @router /signin [get]
func (u *UserController) Signin() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Signin(username, password) {
		u.Data["json"] = "signin success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title signup
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} signup success
// @Failure 403 body is empty
// @router /signup [post]
func (u *UserController) Signup() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Signup(username, password) {
		u.Data["json"] = "signup success"
	} else {
		u.Data["json"] = "body is empty"
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
