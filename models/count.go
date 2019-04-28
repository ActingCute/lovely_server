package models

import (
	"time"
)

type Count struct {
	Id         int64     `xorm:"not null pk autoincr INT(64)" json:"id,omitempty"`
	Url        int64     `xorm:"not null INT(64)" json:"url,omitempty"`
	Count      int64     `xorm:"not null INT(64)" json:"count,omitempty"`
	DeleteTime time.Time `xorm:"DATETIME" json:"delete_time,omitempty"`
}
