package main

import (
	"./config"
	"./handler"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &handler.InDB{DB: db}
	router := gin.Default()

	router.POST("/service-user/login", inDB.LoginUser)
	router.POST("/service-user/register", inDB.UserRegistration)
	router.Run(":3030")
}