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
	ginSetting.Use(gin.Recovery(), middleware.Logger())

	v1 := ginSetting.Group("/v1")
	router.AddUserRouter(v1)

	go func() {
		database.Db()
	}()

	ginSetting.Run(":8000")
}
