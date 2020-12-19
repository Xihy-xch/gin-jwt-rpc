package sql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"workspace/gin-jwt-rpc/models"
)

var db *gorm.DB

func Connection(username string, password string, dbName string) *gorm.DB {
	newDb, err := gorm.Open("mysql", username+":"+password+"@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("connect failed")
	}
	db = newDb
	return db
}

func CheckAccount(username string, password string) bool {
	var user models.User
	return !(db.Where("user_name = ? AND password = ?", username, password).Find(&user).RecordNotFound())
}
