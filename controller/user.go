package controller

import (
	"ginseng_start/model"
	"ginseng_start/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var user []model.User
	err := service.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := service.CreateUser(&user)
	if err != nil {
		// fmt.Println(err.Error())
		log.Println(err.Error())
		println(err.Error())
		// gin.Logger()
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user model.User
	err := service.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		// c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")
	err := service.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
		return
	}
	c.BindJSON(&user)
	err = service.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")
	err := service.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
