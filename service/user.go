package service

import (
	. "HA-back-end/DBMgr"
	. "HA-back-end/models"
	"HA-back-end/public"
)
//
func Login(name, password, ip string)(error,string)  {
	var user User
	MySql.Preload("userName").First(&user, "userName = ?", name)
	MySql.Save(&user)
	return nil, ""
	//return errors.New("login fail"), "vvv"
}
//
func NewUser(name, password ,nickName string,gender int) (error, *User) {
	user := &User{
		Username:    name,
		Password:	 password,
		UserSex:  	 gender,
		NickName:   nickName,
	}
	if len(user.Password) > 0 {
		user.Password = public.Md5String(user.Password)
	} else {
		user.Password = public.Md5String(user.Username)
	}
	dbc := MySql.FirstOrCreate(user, User{Username: user.Username})
	if dbc.Error != nil {
		return dbc.Error, nil
	}
	return nil, user
}

func UpdateUser(name string,ColumnName ,ColumnVal string) (error, *User) {
	user := &User{	}
	if len(user.Password) > 0 {
		user.Password = public.Md5String(user.Password)
	} else {
		user.Password = public.Md5String(user.Username)
	}
	dbc:=MySql.Model(&user).Where("username = ?", name).Update(ColumnName, ColumnVal)
	if dbc.Error != nil {
		return dbc.Error, nil
	}
	return nil, user
}
//
func GetGamesInfo( userName string) (error error, gameIds string) {
	//dbc:=MySql.Raw("select havePlayGameId from users where username=?",userName).Scan(&gameIds)
	user := &User{	}
	dbc:=MySql.Where("username = ?", userName).First(&user)
	if dbc.Error != nil {
		return dbc.Error,""
	}
	gameIds=user.HavePlayGameID
	return nil,  gameIds
}
//
func GetUserByToken(token string) *User {
	var user User
	MySql.Preload("userName").First(&user, "token = ?", token)
	return &user
}

func NewUserLog(log *UserLog) {
	MySql.Create(log)
}
