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
			"Signout",
			`/signout`,
			[]string{"get"},
			nil})

}
