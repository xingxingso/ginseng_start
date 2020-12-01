package service

import (
	"fmt"

	"github.com/xingxingso/ginseng_start/config"
	"github.com/xingxingso/ginseng_start/model"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]model.User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *model.User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *model.User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *model.User, id string) (err error) {
	fmt.Println(user)
	config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *model.User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
