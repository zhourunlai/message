package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username    string `orm:"size(64)"`
	Password    string `orm:"size(64)"`
	Create_time int
	Last_time   int
	Contacts    []*Contact `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Contact struct {
	Username string  `orm:"size(64)"`
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
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	orm.RegisterModel(new(User), new(Contact), new(Chat))
}

func Signin(username, password string) bool {
	return false
}
