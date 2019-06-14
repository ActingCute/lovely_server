package models

import (
	"lovely_server/helper"
	"time"
)

type Twitter struct {
	Id         int64     `xorm:"not null pk autoincr INT(64)" json:"tw_id,omitempty"`
	Content    string    `xorm:"not null TEXT" json:"content,omitempty"`
	UpdateTime time.Time `xorm:"not null DATETIME" json:"update_time,omitempty"`
	DeleteTime time.Time `xorm:"DATETIME" json:"delete_time,omitempty"`
}

func AddTwitter(Twitter *Twitter) error {
	return Insert(Twitter)
}

func DeleteTwitter(Twitter *Twitter) error {
	var whereData []interface{}
	whereData = append(whereData, Twitter.Id)
	return Delete(Twitter, "id=?", whereData)
}

func GetTwitterLimit(limit int, start ...int) (list []Twitter, count int64, err error) {
	err = db.Table("twitter").Where(DELETE_TIME_IS_NULL).Limit(limit, start...).Find(&list)
	if !helper.Error(err) {
		count, err = db.Table("twitter").Where(DELETE_TIME_IS_NULL).Count(new(Twitter))
	}
	return
}
