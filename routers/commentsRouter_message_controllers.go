package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["message/controllers:UserController"] = append(beego.GlobalControllerRouter["message/controllers:UserController"],
		beego.ControllerComments{
			"Signin",
			`/signin`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["message/controllers:UserController"] = append(beego.GlobalControllerRouter["message/controllers:UserController"],
		beego.ControllerComments{
			"Signup",
			`/signup`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["message/controllers:UserController"] = append(beego.GlobalControllerRouter["message/controllers:UserController"],
		beego.ControllerComments{
			"Signout",
			`/signout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["message/controllers:UserController"] = append(beego.GlobalControllerRouter["message/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:username`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["message/controllers:UserController"] = append(beego.GlobalControllerRouter["message/controllers:UserController"],
		beego.ControllerComments{
			"GetContact",
			`/:username/contacts`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["message/controllers:UserController"] = append(beego.GlobalControllerRouter["message/controllers:UserController"],
		beego.ControllerComments{
			"AddContact",
			`/:username/contacts/:contact_username`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["message/controllers:UserController"] = append(beego.GlobalControllerRouter["message/controllers:UserController"],
		beego.ControllerComments{
			"DelContact",
			`/:username/contacts/:contact_username`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["message/controllers:UserController"] = append(beego.GlobalControllerRouter["message/controllers:UserController"],
		beego.ControllerComments{
			"GetChat",
			`/:username/contacts/:contact_username/chats`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["message/controllers:UserController"] = append(beego.GlobalControllerRouter["message/controllers:UserController"],
		beego.ControllerComments{
			"DelChat",
			`/:username/contacts/:contact_username/chats/:id`,
			[]string{"delete"},
			nil})

}
