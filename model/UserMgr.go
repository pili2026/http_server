package model

import (
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
	database.DbConnect.Create(&user)
	return user
}

func DeleteUser(userId string) bool {
	var user schema.User
	result := database.DbConnect.Where("id = ?", userId).Delete(&user)

	return result.RowsAffected != 0
}

func UpdateUser(userId string, user schema.User) schema.User {
	database.DbConnect.Model(&user).Where("id = ?", userId).Updates(user)
	return user
}
