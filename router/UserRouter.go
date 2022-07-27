package router

import (
	session "booking_system/middleware"
	"booking_system/model/schema"
	"booking_system/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/user", session.SetSession())

	user.GET("/all", service.CacheUsersDecorator(service.RedisUsers, "users", schema.User{}))
	user.GET("/:id", service.CacheUserDecorator(service.RedisUser, "id", "user_%s", schema.User{}))
	user.POST("/", service.PostUser)
	user.PUT("/:id", service.UpdateUser)
	user.DELETE("/:id", service.DeleteUser)

	// // TODO: not currently used
	// user.POST("/login", service.LoginUser)
	// user.GET("/check", service.CheckUserSession)

	// // TODO: not currently used
	// user.Use(session.AuthSession())
	// {
	// 	user.DELETE("/:id", service.DeleteUser)
	// 	user.GET("/logout", service.LogoutUser)

	// }
}
