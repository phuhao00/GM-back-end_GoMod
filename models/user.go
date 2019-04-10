package models

import "time"

type User struct {
	ID              int32      `gorm:"primary_key"`
	Username       string    	`gorm:"unique" json:"username"`
	Password        string     `json:"password"`
	User_sex		int32		`json:"user_sex"`
	Nick_name            string     `json:"nick_name"`
}

type Role struct {
	ID    int32  `gorm:"primary_key"`
	Title string `gorm:"unique" json:"title"`
	Name  string `gorm:"unique" json:"name"`
	Users []User `gorm:"many2many:user_role" json:"user"`
}

type UserLog struct {
	ID     int32     `gorm:"primary_key"`
	Name   string    `json:"name"`
	Time   time.Time `json:"time"`
	Url    string    `json:"url"`
	Ip     string    `json:"ip"`
	Status int32     `json:"status"`
}