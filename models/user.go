package models

type User struct {
	Id         int    `xorm:"not null pk autoincr INT(64)" json:"id,omitempty"`
	Website    string `xorm:"not null VARCHAR(255)" json:"website,omitempty"`
	Email      string `xorm:"not null VARCHAR(255)" json:"email,omitempty"`
	Name       string `xorm:"not null VARCHAR(255)" json:"name,omitempty"`
	DeleteTime int    `xorm:"INT(11)" json:"delete_time,omitempty"`
}
