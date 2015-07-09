package models

import "time"

//Register model
type Register struct {
	Id      int       `form:"-"`
	Mail    string    `orm:"unique;type(email)" form:"mail,mail" valid:"Email; MaxSize(100)"`
	IP      string    `form:"text"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	User    *User     `orm:"reverse(one)"`
}

func (r *Register) TableName() string {
	return "register"
}
