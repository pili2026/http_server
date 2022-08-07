package model

import (
	"github.com/google/uuid"

	"booking_system/database"
	"booking_system/model/schema"
)

func GetUsers() []schema.User {
	var users []schema.User
	database.DbConnect.Find(&users)
	return users
}

func GetUserById(userId string) schema.User {
	var user schema.User
	database.DbConnect.Where("id = ?", userId).First(&user)
	return user
}

func CreateUser(user schema.User) schema.User {
	id := uuid.NewSHA1(uuid.NameSpaceDNS, []byte(user.Email))
	user.Id = id
	database.DbConnect.Create(&user)
	return user
}

func DeleteUser(userId string) bool {
	var user schema.User
	result := database.DbConnect.Where("id = ?", userId).Delete(&user)

	return result.RowsAffected != 0
}

func UpdateUser(userId string, user schema.User) bool {
	result := database.DbConnect.Model(&user).Where("id = ?", userId).Updates(user)
	return result.RowsAffected != 0

}

func CheckUserPassword(account string, password string) schema.User {
	user := schema.User{}
	database.DbConnect.Where("account = ? and password = ?", account, password).First(&user)
	return user
}
