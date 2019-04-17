package models

import "time"
type User struct {
	ID             int          `gorm:"column:id;primary_key" json:"id"`
	Username       string 		`gorm:"column:username" json:"usernasme"`
	Password       string 		`gorm:"column:password" json:"password"`
	UserSex        int    	`gorm:"column:user_sex" json:"user_sex"`
	NickName       string		`gorm:"column:nick_name" json:"nick_name"`
	HavePlayGameID string 		`gorm:"column:havePlayGameId" json:"havePlayGameId"`
}

type UserLog struct {
	ID     int32     `gorm:"primary_key"`
	Name   string    `json:"name"`
	Time   time.Time `json:"time"`
	Url    string    `json:"url"`
	Ip     string    `json:"ip"`
	Status int32     `json:"status"`
}