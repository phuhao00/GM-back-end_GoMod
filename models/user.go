package models

import "time"

type User struct {
	ID              int32      `gorm:"primary_key"`
	Username       string    	`gorm:"unique" json:"username"`
	Password        string     `json:"password"`
	User_sex		int32		`json:"user_sex"`
	Nick_name            string     `json:"nick_name"`
}

type UserLog struct {
	ID     int32     `gorm:"primary_key"`
	Name   string    `json:"name"`
	Time   time.Time `json:"time"`
	Url    string    `json:"url"`
	Ip     string    `json:"ip"`
	Status int32     `json:"status"`
}