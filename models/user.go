package models

import "time"

//User model
type User struct {
	Id       int       `form:"-"`
	Mail     string    `orm:"unique;type(email)" form:"mail,mail" valid:"Email; MaxSize(100)"`
	Md5Mail  string    `orm:unique`
	Role     string    `form:"role"`
	Pass     string    `orm:"null;type(password)" form:"pass,password"`
	Created  time.Time `orm:"null;auto_now_add;type(datetime)"`
	Register *Register `orm:"null;rel(one);on_delete(set_null)"`
}

func (u *User) TableName() string {
	return "user"
}
