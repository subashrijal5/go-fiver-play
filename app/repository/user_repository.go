package repository

import (
	"github.com/subashrijal5/go-fiber-boilerplate/app/model"
	"github.com/subashrijal5/go-fiber-boilerplate/utils/database"
)

var db = database.GetDB()

func GetUserByEmail(email string) (model.User, error) {
	var user model.User
	data := db.First(&user, "email = ?", email)
	return user, data.Error
}

func CreateUser(user model.User) error {
	data := db.Create(&user)
	return data.Error
}

func GetUserByID(id uint) (model.User, error) {
	var user model.User
	data := db.First(&user, id)
	return user, data.Error
}

func UpdateUser(user model.User) error {
	data := db.Save(&user)
	return data.Error
}

func DeleteUser(user model.User) error {
	data := db.Delete(&user)
	return data.Error
}

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	data := db.Find(&users)
	return users, data.Error
}
