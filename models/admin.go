package models

import (
	"time"
	"lovely_server/helper"
)

type Admin struct {
	Id         int64     `xorm:"not null pk autoincr INT(64)" json:"aid,omitempty"`
	UserName   string    `xorm:"not null unique VARCHAR(30)" json:"user_name,omitempty"`
	PassWord   string    `xorm:"not null unique VARCHAR(255)" json:"pass_word,omitempty"`
	Token      string    `xorm:"not null unique VARCHAR(32)" json:"token,omitempty"`
	DeleteTime time.Time `xorm:"comment('删除时间') DATETIME" json:"delete_time,omitempty"`
}

func AddAdmin(admin *Admin) error {
	has, _, err := GetAdminByName(admin.UserName)
	if helper.Error(err) {
		return err
	}
	if !has {
		admin.Token = helper.GetRandomString(32)
		err = Insert(admin)
		if !helper.Error(err) {
			admin.PassWord = GetPassword(admin.PassWord, admin.Id)
			var whereData []interface{}
			whereData = append(whereData, admin.Id)
			return Update(admin, "id=?", whereData, "pass_word")
		}
	}
	helper.Debug("has admin -- ", admin)
	return nil
}

func GetAdminByName(name string) (has bool, Admin Admin, err error) {
	has, err = db.Table("admin").Where("user_name =?", name).Get(&Admin)
	return
}

func CkeckAdmin(admin Admin) (has bool, Admin Admin, err error) {
	has, err = db.Table("admin").Where("user_name =? and pass_word=?", admin.UserName, admin.PassWord).Get(&Admin)
	if has {
		//更新token
		Admin.Token = helper.GetRandomString(32)
		var whereData []interface{}
		whereData = append(whereData, Admin.Id)
		err = Update(Admin, "id=?", whereData, "token")
	}
	Admin.PassWord = ""
	return
}

func GetAdminByToken(admin Admin) (has bool, Admin Admin, err error) {
	has, err = db.Table("admin").Where("token =?", admin.Token).Get(&Admin)
	Admin.PassWord = ""
	return
}
