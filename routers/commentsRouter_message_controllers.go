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
			`/:username/contact`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["message/controllers:UserController"] = append(beego.GlobalControllerRouter["message/controllers:UserController"],
		beego.ControllerComments{
			"AddContact",
			`/:username/contact/:contact_username`,
			[]string{"get"},
			nil})

}
