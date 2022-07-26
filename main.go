package main

import (
	"booking_system/database"
	"booking_system/middleware"
	"booking_system/router"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func setupLogFile() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func main() {
	setupLogFile()

	ginSetting := gin.Default()

	// TODO: Validator for registered passwords if required
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("password", middleware.UserPw)
	// }

	ginSetting.Use(gin.Recovery(), middleware.Logger())

	v1 := ginSetting.Group("/v1")
	router.AddUserRouter(v1)

	database.Db()
	database.MongoDb()
	defer database.MongoDisconnect()

	ginSetting.Run(":8000")
}
