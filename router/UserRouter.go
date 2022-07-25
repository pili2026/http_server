package router

import (
	"booking_system/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")

	user.GET("/", service.GetUsers)
	user.GET("/:id", service.GetUserById)
	user.POST("/", service.PostUser)
	user.PUT("/:id", service.UpdateUser)
	user.DELETE("/:id", service.DeleteUser)
}
