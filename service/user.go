package service

import (
	. "HA-back-end/DBMgr"
	. "HA-back-end/models"
	"HA-back-end/public"
	"errors"
)
//
func Login(name, password, ip string)(error,string)  {
	var user User
	MySql.Preload("userName").First(&user, "userName = ?", name)
	MySql.Save(&user)
	return nil, ""
	return errors.New("login fail"), "vvv"
}
//
func NewUser(name, password string, roles []string, introduce string) (error, *User) {
	user := &User{
		Username:      name,
		Password: password,
	}
	if len(user.Password) > 0 {
		user.Password = public.Md5String(user.Password)
	} else {
		user.Password = public.Md5String(user.Username)
	}
	if len(introduce) > 0 {
	}

	dbc := MySql.FirstOrCreate(user, User{Username: user.Username})
	if dbc.Error != nil {
		return dbc.Error, nil
	}
	var rs []Role
	for _, v := range roles {
		var role Role
		dbc = MySql.Preload("User").First(&role, "name = ?", v)
		if dbc.Error != nil {
			return dbc.Error, nil
		}
		rs = append(rs, role)
	}
	MySql.Model(user).Association("Roles").Replace(rs)
	return nil, user
}

func GetUserByToken(token string) *User {
	var user User
	MySql.Preload("userName").First(&user, "token = ?", token)
	return &user
}

func NewUserLog(log *UserLog) {
	MySql.Create(log)
}