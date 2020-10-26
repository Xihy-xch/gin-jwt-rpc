package sql

import (
	"fmt"
	"models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connection(username string, password string, dbName string) *gorm.DB {
	newDb, err := gorm.Open("mysql", username+":"+password+"@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("connect failed")
	}
	db = newDb
	defer db.Close()
	return db
}

func CheckAccount(username string, password string) bool {
	var user models.User
	return !(db.Where("user_name = ? AND password = ?", username, password).Find(&user).RecordNotFound())
}
