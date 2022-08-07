package router

import (
	session "booking_system/middleware"
	"booking_system/model/schema"
	"booking_system/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/user", session.SetSession())

	user.GET("/", service.CacheUsersDecorator(service.RedisUsers, "users", schema.User{}))
	user.GET("/:id", service.CacheUserDecorator(service.RedisUser, "id", "user_%s", schema.User{}))
	user.POST("/", service.PostUser)
	user.PUT("/:id", service.UpdateUser)
	user.DELETE("/:id", service.DeleteUser)

	mogUser := user.Group("/mongo")
	mogUser.GET("/", service.MongoFindUsers)
	mogUser.GET("/:id", service.MongoFindUser)
	mogUser.POST("/", service.MongoCreateUser)
	mogUser.PUT("/:id", service.MongoUpdateUser)
	mogUser.DELETE("/:id", service.MongoDeleteUser)

	// TODO: not currently used
	// user.POST("/login", service.LoginUser)
	// user.GET("/check", service.CheckUserSession)

	// TODO: not currently used
	// user.Use(session.AuthSession())
	// {
	// 	user.DELETE("/:id", service.DeleteUser)
	// 	user.GET("/logout", service.LogoutUser)

	// }
}
