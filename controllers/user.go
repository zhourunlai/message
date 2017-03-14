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
// @Success 100 {string} signin success
// @Failure 101 signin failed
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

// @Title signup
// @Description user signup
// @Param	username		query 	string	true		"The username for signup"
// @Param	password		query 	string	true		"The password for signup"
// @Success 200 {string} signup success
// @Failure 201 signup failed
// @router /signup [post]
func (u *UserController) Signup() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Signup(username, password) {
		u.Data["json"] = "signup success"
	} else {
		u.Data["json"] = "signup failed"
	}
	u.ServeJSON()
}

// @Title signout
// @Description user signout
// @Success 300 {string} signout success
// @Failure 301 signout failed
// @router /signout [get]
func (u *UserController) Signout() {
	u.Data["json"] = "signout success"
	u.ServeJSON()
}

// @Title getUser
// @Description user getUser
// @Param	username		query 	string	true		"The username for getUser"
// @Success 400 {string} getUser success
// @Failure 401 getUser failed
// @router /:username [get]
func (u *UserController) Get() {
	username := u.GetString(":username")
	if models.GetUser(username) {
		u.Data["json"] = "getUser success"
	} else {
		u.Data["json"] = "getUser failed"
	}
	u.ServeJSON()
}

// @Title getContact
// @Description user getContact
// @Param	username		query 	string	true		"The username for getContact"
// @Success 500 {string} getContact success
// @Failure 501 getContact failed
// @router /:username/contact [get]
func (u *UserController) GetContact() {
	username := u.GetString("username")
	if models.GetContact(username) {
		u.Data["json"] = "getContact success"
	} else {
		u.Data["json"] = "getContact failed"
	}
	u.ServeJSON()
}

// @Title addContact
// @Description user addContact
// @Param	username		query 	string	true		"The username for addContact"
// @Success 600 {string} addContact success
// @Failure 601 addContact failed
// @router /:username/contact/:contact_username [get]
func (u *UserController) AddContact() {
	username := u.GetString("username")
	contact := u.GetString("contact_username")
	if models.AddContact(username, contact) {
		u.Data["json"] = "addContact success"
	} else {
		u.Data["json"] = "addContact failed"
	}
	u.ServeJSON()
}
