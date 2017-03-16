package controllers

import (
	"message/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// Cookie
var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()
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
		sess := u.StartSession()
		sess.Set("username", username)
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
	sess := u.StartSession()
	sess.Delete("username")
	u.Data["json"] = "signout success"
	u.ServeJSON()
}

// @Title getUser
// @Description user getUser
// @Param	username		query 	string	true		"The username for getUser"
// @Success 400 {string} getUser success, show username
// @Failure 401 getUser failed
// @router /:username [get]
func (u *UserController) Get() {
	username := u.GetString(":username")
	if models.GetUser(username) {
		u.Data["json"] = username
	} else {
		u.Data["json"] = "getUser failed"
	}
	u.ServeJSON()
}

// @Title getContact
// @Description user getContact
// @Param	username		query 	string	true		"The username for getContact"
// @Success 500 {string} getContact success, show contacts
// @Failure 501 getContact failed
// @router /:username/contacts [get]
func (u *UserController) GetContact() {
	username := u.GetString(":username")
	contacts, err := models.GetContact(username)
	if err == nil {
		u.Data["json"] = contacts
	} else {
		u.Data["json"] = "getContact failed"
	}
	u.ServeJSON()
}

// @Title addContact
// @Description user addContact
// @Param	username		query 	string	true		"The username for me"
// @Param	contact			query 	string	true		"The username for contact"
// @Success 600 {id} addContact success , show id
// @Failure 601 addContact failed
// @router /:username/contacts/:contact_username [get]
func (u *UserController) AddContact() {
	username := u.GetString(":username")
	contact := u.GetString(":contact_username")
	id, err := models.AddContact(username, contact)
	if err == nil {
		u.Data["json"] = id
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
// @Success 800 {string} getChat success, show chats
// @router /:username/contacts/:contact_username/chats [get]
func (u *UserController) GetChat() {
	username := u.GetString(":username")
	contact := u.GetString(":contact_username")
	u.Data["json"] = models.GetChat(username, contact)
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

// @Title getUnreadChat
// @Description user getUnreadChat
// @Param	username		query 	string	true		"The username for me"
// @Param	contact			query 	string	true		"The username for contact"
// @Success 800 {string} getUnreadChat success, show unread chats count
// @router /:username/contacts/:contact_username/chats/unread [get]
func (u *UserController) GetUnreadChat() {
	username := u.GetString(":username")
	contact := u.GetString(":contact_username")
	u.Data["json"] = models.GetUnreadChat(username, contact)
	u.ServeJSON()
}
