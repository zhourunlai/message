package models

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Username    string `orm:"pk"`
	Password    string `orm:"size(64)"`
	Create_time int
	Last_time   int
	Contacts    []*Contact `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Contact struct {
	Username string  `orm:"pk"`
	Contact  string  `orm:"size(64)"`
	User     *User   `orm:"rel(fk)"`       //设置一对多关系
	Chats    []*Chat `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Chat struct {
	Id        int
	Sender    string `orm:"size(64)"`
	Receiver  string `orm:"size(64)"`
	Content   string
	Send_time int
	Is_del    int
	Is_read   int
	Contact   *Contact `orm:"rel(fk)"` //设置一对多关系
}

func init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpass := beego.AppConfig.String("db.pass")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")

	conn := dbuser + ":" + dbpass + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		conn = conn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", conn, 5, 30)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	orm.RegisterModel(new(User), new(Contact), new(Chat))
}

func Signin(username, password string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("users")
	qs.Filter("username", username)
	qs.Filter("password", password)
	return qs.Exist()
}
