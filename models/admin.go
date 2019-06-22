package models

import (
	"github.com/pkg/errors"
	"lovely_server/helper"
	"time"
)

type Admin struct {
	Id         int64     `xorm:"not null pk autoincr INT(64)" json:"aid,omitempty"`
	UserName   string    `xorm:"not null unique VARCHAR(30)" json:"user_name,omitempty"`
	PassWord   string    `xorm:"not null unique VARCHAR(255)" json:"pass_word,omitempty"`
	Token      string    `xorm:"not null unique VARCHAR(32)" json:"token,omitempty"`
	LastLogin  time.Time `xorm:"comment('最后登录时间') not null DATETIME" json:"last_login,omitempty"`
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
			admin.LastLogin = time.Now()
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
	has, Admin, err = GetAdminByName(admin.UserName)
	if !has {
		return has, admin, err
	}
	if Admin.PassWord == GetPassword(admin.PassWord, Admin.Id) {
		//更新token
		Admin.LastLogin = time.Now()
		Admin.Token = helper.GetRandomString(32)
		var whereData []interface{}
		whereData = append(whereData, Admin.Id)
		err = Update(Admin, "id=?", whereData, "token", "last_login")
		Admin.PassWord = ""
	} else {
		err = errors.New("pass_word error")
	}
	return
}

func GetAdminByToken(admin Admin) (has bool, Admin Admin, err error) {
	has, err = db.Table("admin").Where("token =?", admin.Token).Get(&Admin)
	Admin.PassWord = ""
	//登陆有效期检测
	if time.Now().Unix()-Admin.LastLogin.Unix() > 604800 {
		err = errors.New("登陆过期")
	}
	return
}
