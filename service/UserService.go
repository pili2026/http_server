package service

import (
	"booking_system/model"
	"booking_system/model/schema"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET all user
func GetUsers(ctx *gin.Context) {
	users := model.GetUsers()
	ctx.JSON(http.StatusOK, users)
}

// GET user by id
func GetUserById(ctx *gin.Context) {
	IdParam := ctx.Param("id")
	user := model.GetUserById(IdParam)

	if user.Id.String() != IdParam {
		ctx.JSON(http.StatusNotFound, "User not found")
		return
	}
	log.Println("User->", user)
	ctx.JSON(http.StatusOK, user)
}

// POST
func PostUser(ctx *gin.Context) {
	user := schema.User{}
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	}
	createdUser := model.CreateUser(user)
	ctx.JSON(http.StatusOK, createdUser)
}

// PUT
func UpdateUser(ctx *gin.Context) {
	user := schema.User{}
	IdParam := ctx.Param("id")
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Error")
		return
	}

	userResult := model.UpdateUser(IdParam, user)

	if !userResult {
		ctx.JSON(http.StatusNotFound, "Can't update user due to not found")
		log.Println("Can't update user due to not found")
		return
	}
	ctx.JSON(http.StatusNoContent, "Updated")

}

// Delete
func DeleteUser(ctx *gin.Context) {
	user := model.DeleteUser(ctx.Param("id"))

	if !user {
		ctx.JSON(http.StatusNotFound, "Can't delete user due to not found")
		log.Println("Can't delete user due to not found")
		return
	}
	ctx.JSON(http.StatusOK, "Deleted")
}
