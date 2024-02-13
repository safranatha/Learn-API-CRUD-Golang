package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabseInit() {
	var err error
	source_database := "root:@tcp(127.0.0.1:3306)/api_golang?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(source_database), &gorm.Config{})
	if err != nil {
		panic("cannot connect database")
	}
	fmt.Println("database connect")
}
