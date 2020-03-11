package config

import (
	"../structs"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

// DBInit create connection to DB
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/dagun?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.User{})
	return db
}
