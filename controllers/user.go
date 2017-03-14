package controllers

import (
	"message/models"
	"strconv"

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
// @router /:username/contacts [get]
func (u *UserController) GetContact() {
	username := u.GetString(":username")
	if models.GetContact(username) {
		u.Data["json"] = "getContact success"
	} else {
		u.Data["json"] = "getContact failed"
	}
	u.ServeJSON()
}

// @Title addContact
// @Description user addContact
// @Param	username		query 	string	true		"The username for me"
// @Param	contact			query 	string	true		"The username for contact"
// @Success 600 {string} addContact success
// @Failure 601 addContact failed
// @router /:username/contacts/:contact_username [get]
func (u *UserController) AddContact() {
	username := u.GetString(":username")
	contact := u.GetString(":contact_username")
	if models.AddContact(username, contact) {
		u.Data["json"] = "addContact success"
	} else {
		u.Data["json"] = "addContact failed"
	}
	u.ServeJSON()
}

// @Title delContact
// @Description user delContact
// @Param	username		query 	string	true		"The username for me"
// @Param	contact			query 	string	true		"The username for contact"
// @Success 700 {string} delContact success
// @Failure 701 delContact failed
// @router /:username/contacts/:contact_username [delete]
func (u *UserController) DelContact() {
	username := u.GetString(":username")
	contact := u.GetString(":contact_username")
	if models.DelContact(username, contact) {
		u.Data["json"] = "delContact success"
	} else {
		u.Data["json"] = "delContact failed"
	}
	u.ServeJSON()
}

// @Title getChat
// @Description user getChat
// @Param	username		query 	string	true		"The username for me"
// @Param	contact			query 	string	true		"The username for contact"
// @Success 800 {string} getChat success
// @Failure 801 getChat failed
// @router /:username/contacts/:contact_username/chats [get]
func (u *UserController) GetChat() {
	username := u.GetString(":username")
	contact := u.GetString(":contact_username")
	if models.GetChat(username, contact) {
		u.Data["json"] = "getChat success"
	} else {
		u.Data["json"] = "getChat failed"
	}
	u.ServeJSON()
}

// @Title delChat
// @Description user delContact
// @Param	id		query 	string	true		"The id for chat"
// @Success 900 {string} delChat success
// @Failure 901 delChat failed
// @router /:username/contacts/:contact_username/chats/:id [delete]
func (u *UserController) DelChat() {
	id_str := u.GetString(":id")
	id_int, err := strconv.ParseInt(id_str, 10, 64)
	if err != nil {
	}
	if models.DelChat(id_int) {
		u.Data["json"] = "delChat success"
	} else {
		u.Data["json"] = "delChat failed"
	}
	u.ServeJSON()
}

// @Title updateChat
// @Description user updateChat
// @Param	id		query 	string	true		"The id for chat"
// @Success 1000 {string} updateChat success
// @Failure 1001 updateChat failed
// @router /:username/contacts/:contact_username/chats/:id [get]
func (u *UserController) UpdateChat() {
	id_str := u.GetString(":id")
	id_int, err := strconv.ParseInt(id_str, 10, 64)
	if err != nil {
	}
	if models.UpdateChat(id_int) {
		u.Data["json"] = "updateChat success"
	} else {
		u.Data["json"] = "updateChat failed"
	}
	u.ServeJSON()
}
