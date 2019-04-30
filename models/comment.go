package models

import (
	"time"
)

type Comment struct {
	Id         int64     `xorm:"not null pk autoincr INT(64)" json:"commid,omitempty"`
	UserId     int64     `xorm:"not null INT(64)" json:"user_id,omitempty"`
	Url        string    `xorm:"not null VARCHAR(255)" json:"url,omitempty"`
	FatherId   int64     `xorm:"default 0 INT(64)" json:"father_id,omitempty"`
	SonId      int64     `xorm:"default 0 INT(64)" json:"son_id,omitempty"`
	Content    string    `xorm:"not null TEXT" json:"content,omitempty"`
	UpdateTime time.Time `xorm:"not null DATETIME" json:"update_time,omitempty"`
	DeleteTime time.Time `xorm:"DATETIME" json:"delete_time,omitempty"`
}

type CommentsExtend struct {
	User    User    `xorm:"extends" json:"user"`
	Comment Comment `xorm:"extends" json:"comment"`
}

const comment_user_cols = "user.id,user.name,user.website,user.email,comment.id,comment.user_id,comment.father_id,comment.son_id,comment.content,comment.update_time"

func GetCommentsByUrl(url string) ([]CommentsExtend, error) {
	comments := make([]CommentsExtend, 0)
	err := db.Table("comment").Join("RIGHT", "user", "user.id=comment.user_id").Cols(comment_user_cols).Where("url =?", url).Find(&comments)
	return comments, err
}

func AddComment(comment *Comment) error {
	return Insert(comment)
}
