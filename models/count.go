package models

import (
	"lovely_server/helper"
	"time"
)

type Count struct {
	Id         int64     `xorm:"not null pk autoincr INT(64)" json:"id,omitempty"`
	Url        string    `xorm:"not null VARCHAR(255)" json:"url,omitempty"`
	Count      int64     `xorm:"not null INT(64)" json:"count,omitempty"`
	DeleteTime time.Time `xorm:"DATETIME" json:"delete_time,omitempty"`
}

func AddCount(count Count) error {
	//先查是否存在数据，若存在则增加
	has, oldCount, err := GetCountByUrl(count.Url)
	if has {
		helper.Debug("oldCount ---", oldCount)
		whereData := make([]interface{}, 0)
		oldCount.Count++
		whereData = append(whereData, oldCount.Url)
		err = Update(oldCount, "url=?", whereData, "count")
		return err
	}
	return Insert(count)
}

func GetCountByUrl(url string) (has bool, count Count, err error) {
	has, err = db.Table("count").Where("url =?", url).Get(&count)
	return has, count, err
}

func GetCountsByUrl(url string) ([]Count, error) {
	counts := make([]Count, 0)
	err := db.Table("count").Where("url =?", url).Find(&counts)
	return counts, err
}
