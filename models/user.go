package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
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
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/message?charset=utf8")
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.RegisterModel(new(User), new(Contact), new(Chat))
	orm.DefaultTimeLoc = time.UTC
}

func Signin(username, password string) bool {
	o := orm.NewOrm()
	sql := "SELECT * FROM user WHERE username=" + username + " AND password=" + password
	num, err := o.Raw(sql).QueryRows()
	if num != 0 && err != nil {
		return true
	}
	return false
}
