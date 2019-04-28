package models

import (
	"lovely_server/helper"
	"time"
)

type User struct {
	Id         int64     `xorm:"not null pk autoincr INT(64)" json:"id,omitempty"`
	Website    string    `xorm:"VARCHAR(255)" json:"website,omitempty"`
	Email      string    `xorm:"VARCHAR(255)" json:"email,omitempty"`
	Name       string    `xorm:"not null comment('名字') unique VARCHAR(255)" json:"name,omitempty"`
	DeleteTime time.Time `xorm:"comment('删除时间') DATETIME" json:"delete_time,omitempty"`
}

func AddUser(user *User) error {
	//先查是否存在数据，若存在则增加
	has, oldUser, err := GetUserByName(user.Name)
	if helper.Error(err) {
		return err
	}
	if has {
		helper.Debug("oldUser ---", oldUser)
		whereData := make([]interface{}, 0)
		if len(user.Website) > 0 {
			oldUser.Website = user.Website
		}
		if len(user.Email) > 0 {
			oldUser.Email = user.Email
		}
		whereData = append(whereData, oldUser.Name)
		err = Update(oldUser, "name=?", whereData, "website", "email")
		return err
	}
	return Insert(user)
}

func GetUserByName(name string) (has bool, user User, err error) {
	has, err = db.Table("user").Where("name =?", name).Get(&user)
	return
}
