package DBMgr

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var MySql *gorm.DB
var err error

type ConnMysql struct {
	Name string
	User string
	Host string
	Password string
}
func LoadMysqlConfig() *ConnMysql {

	return &ConnMysql{
		"game",
		"root",
		"0.0.0.0:3306",
		"huhao123",
	}
}
func (self *ConnMysql) OpenDB() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	self.Init(self.User, self.Password, self.Host, self.Name)
}

func (self *ConnMysql)Init(user, password, host, name string)  {
	MySql, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, name))
	if err != nil {
		log.Println(err)
	}
	MySql.LogMode(true)
	//DB.SingularTable(true)
	MySql.DB().SetMaxOpenConns(1000)
	MySql.DB().SetMaxIdleConns(500)
}