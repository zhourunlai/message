package models

import (
	"net/url"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username    string `orm:"pk"`
	Password    string `orm:"size(64)"`
	Create_time int64
	Last_time   int64
	Contacts    []*Contact `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Contact struct {
	Username string  `orm:"pk"`
	Contact  string  `orm:"size(64)"`
	User     *User   `orm:"rel(fk)"`       //设置一对多关系
	Chats    []*Chat `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Chat struct {
	Id        int    `orm:"pk"`
	Sender    string `orm:"size(64)"`
	Receiver  string `orm:"size(64)"`
	Content   string
	Send_time int64
	Is_del    int
	Is_read   int
	Contact   *Contact `orm:"rel(fk)"` //设置一对多关系
}

func (u *User) TableName() string {
	return "users"
}

func (u *Contact) TableName() string {
	return "contacts"
}

func (u *Chat) TableName() string {
	return "chats"
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
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", conn, 5, 30)
	orm.RegisterModel(new(User), new(Contact), new(Chat))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func Signin(username, password string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	qs.Filter("username", username)
	qs.Filter("password", password)
	return qs.Exist()
}

func Signup(username, password string) bool {
	o := orm.NewOrm()
	u := User{Username: username, Password: password, Create_time: time.Now().Unix()}
	_, err := o.Insert(&u)
	if err != nil {
		return false
	}
	return true
}
