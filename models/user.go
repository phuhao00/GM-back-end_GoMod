package models

import "time"

type User struct {
	ID              int32      `gorm:"primary_key"`
	Name        string     `gorm:"unique" json:"name"`
	Password        string     `json:"password"`
	Nick            string     `json:"nick"`
	Token           string     `json:"token"`
	TokenExpireTime *time.Time `json:"token_expire_time"`
	Avatar          string     `json:"avatar"`
	Introduction    string     `json:"introduction"`
	Roles           []string     `gorm:"many2many:user_role" json:"roles"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	LoginTimes      int32      `json:"login_times"`
	LastLoginIp     string     `json:"last_login_ip"`
}

type Role struct {
	ID    int32  `gorm:"primary_key"`
	Title string `gorm:"unique" json:"title"`
	Name  string `gorm:"unique" json:"name"`
	Users []User `gorm:"many2many:user_role" json:"users"`
}