package models

import (
	"time"
)

type Comment struct {
	Id         int64     `xorm:"not null pk autoincr INT(64)" json:"id,omitempty"`
	UserId     int64     `xorm:"not null INT(64)" json:"user_id,omitempty"`
	Url        string    `xorm:"not null VARCHAR(255)" json:"url,omitempty"`
	FatherId   int64     `xorm:"default 0 INT(64)" json:"father_id,omitempty"`
	SonId      int64     `xorm:"default 0 INT(64)" json:"son_id,omitempty"`
	Content    string    `xorm:"not null TEXT" json:"content,omitempty"`
	UpdateTime time.Time `xorm:"not null DATETIME" json:"update_time,omitempty"`
	DeleteTime time.Time `xorm:"DATETIME" json:"delete_time,omitempty"`
}

func AddComment(comment Comment) error {
	return Insert(comment)
}
