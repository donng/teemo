package model

import (
	"fmt"
	"log"

	"github.com/donng/teemo/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID uint `gorm:"primary_key" json:"id"`
}

func init() {
	conf := setting.Setting.Mysql
	log.Println(conf)

	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.DBName))
	if err != nil {
		panic(fmt.Sprintf("connect mysql error, err: %s", err))
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.TablePrefix + defaultTableName
	}
}
