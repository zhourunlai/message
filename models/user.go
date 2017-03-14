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
	Id       int64   `orm:"pk"`
	Username string  `orm:"size(64)"`
	Contact  string  `orm:"size(64)"`
	User     *User   `orm:"rel(fk)"`       //设置一对多关系
	Chats    []*Chat `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Chat struct {
	Id        int64  `orm:"pk"`
	Sender    string `orm:"size(64)"`
	Receiver  string `orm:"size(64)"`
	Content   string
	Send_time int64
	Is_del    int8
	Is_read   int8
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
	id := o.QueryTable(new(User)).Filter("username", username).Filter("password", password).Exist()
	return id
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

func GetUser(username string) bool {
	o := orm.NewOrm()
	var u User
	err := o.QueryTable(new(User)).Filter("username", username).One(&u)
	if err != nil {
		return false
	}
	return true
}

func GetContact(username string) bool {
	o := orm.NewOrm()
	var c Contact
	// TODO
	_, err := o.QueryTable(new(Contact)).Filter("username", username).All(&c)
	if err != nil {
		return false
	}
	return true
}

func AddContact(username, contact string) bool {
	o := orm.NewOrm()
	u := Contact{Username: username, Contact: contact}
	_, err := o.Insert(&u)
	if err != nil {
		return false
	}
	return true
}

func DelContact(username, contact string) bool {
	o := orm.NewOrm()
	var c Contact
	err1 := o.QueryTable(new(Contact)).Filter("username", username).Filter("contact", contact).One(&c, "Id")
	if err1 == orm.ErrNoRows {
		return false
	}
	_, err2 := o.QueryTable(new(Contact)).Filter("id", c.Id).Delete()
	if err2 != nil {
		return false
	}
	return true
}

func GetChat(username, contact string) bool {
	o := orm.NewOrm()
	var t Chat
	// TODO
	_, err1 := o.QueryTable(new(Chat)).Filter("sender", username).Filter("receiver", contact).All(&t)
	_, err2 := o.QueryTable(new(Chat)).Filter("receiver", username).Filter("sender", contact).All(&t)
	if err1 != nil && err2 != nil {
		return false
	}
	return true
}

func DelChat(id int64) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Chat)).Filter("id", id).Delete()
	if err != nil {
		return false
	}
	return true
}

func UpdateChat(id int64) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Chat)).Filter("id", "id").Update(orm.Params{
		"is_read": 1,
	})
	if err != nil {
		return false
	}
	return true
}
