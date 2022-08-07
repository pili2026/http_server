package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"booking_system/database"
	"booking_system/model"
	"booking_system/model/schema"
)

func MongoCreateUser(ctx *gin.Context) {
	message := ""
	user := schema.MongoUser{}
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}

	createdUser, isExisted := model.MgoCreateUser(user)

	if isExisted {
		message = "This user already existed"
	} else {
		message = "Create user successfully"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
		"user":    createdUser,
	})
}

func MongoFindUsers(ctx *gin.Context) {
	users := model.MgoFindUsers()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Find all user successfully",
		"user":    users,
	})
}

func MongoFindUser(ctx *gin.Context) {
	user := schema.MongoUser{}
	user = model.MgoFindUserById(ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Find user successfully",
		"user":    user,
	})
}

func MongoUpdateUser(ctx *gin.Context) {
	user := schema.User{}
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}

	user = model.MgoUpdateUser(ctx.Param("id"), user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Updated user successfully",
		"user":    user,
	})
}

func MongoDeleteUser(ctx *gin.Context) {
	user := model.MgoDeleteUserById(ctx.Param("id"))

	if !user {
		ctx.JSON(http.StatusNotFound, "Error")
		return
	}
	ctx.JSON(http.StatusOK, "Deleted")
}

func RedisCacheMgoUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user := schema.User{}

	// FIXME: Find function need fix it.
	// database.DbConnect.Find(&user, id)
	database.DbConnect.Where("id = ?", id).First(&user)
	ctx.Set("dbResult", user)
}

func RedisCacheMgoUsers(ctx *gin.Context) {
	users := model.MgoFindUsers()
	ctx.Set("dbMgoUsers", users)
}
