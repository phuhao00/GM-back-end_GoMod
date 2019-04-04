package service

import (
	. "HA-back-end/DBMgr"
	. "HA-back-end/models"
	"HA-back-end/public"
	"errors"
	"time"
)
//
func Login(name, password, ip string)(error,string)  {
	var user User
	MySql.Preload("Roles").First(&user, "name = ?", name)
	if user.Password == public.Md5String(password) {
		if len(user.Token) < 1 || user.TokenExpireTime.Unix() < time.Now().Unix() {
			user.Token = public.RandomString(32)
			tet := time.Now().Add(2 * 60 * 60 * time.Second)
			user.TokenExpireTime = &tet
		}
		user.LastLoginIp = ip
		user.LoginTimes++
		now := time.Now()
		user.UpdatedAt = &now
		MySql.Save(&user)
		return nil, user.Token
	}
	return errors.New("login fail"), ""
}
//
func NewUser(name, password string, roles []string, introduce string) (error, *User) {
	user := &User{
		Name:      name,
		Password:  password,
		CreatedAt: time.Now(),
		Avatar:    "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
	}
	if len(user.Password) > 0 {
		user.Password = public.Md5String(user.Password)
	} else {
		user.Password = public.Md5String(user.Name)
	}
	if len(introduce) > 0 {
		user.Introduction = introduce
	}

	dbc := MySql.FirstOrCreate(user, User{Name: user.Name})
	if dbc.Error != nil {
		return dbc.Error, nil
	}
	var rs []Role
	for _, v := range roles {
		var role Role
		dbc = MySql.Preload("Users").First(&role, "name = ?", v)
		if dbc.Error != nil {
			return dbc.Error, nil
		}
		rs = append(rs, role)
	}
	MySql.Model(user).Association("Roles").Replace(rs)
	return nil, user
}
