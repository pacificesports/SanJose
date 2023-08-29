package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sanjose/model"
	"sanjose/service"
)

func GetAllUsers(c *gin.Context) {
	result := service.GetAllUsers()
	c.JSON(http.StatusOK, result)
}

func GetUserByID(c *gin.Context) {
	result := service.GetUserByID(c.Param("userID"))
	if result.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "No user found with given id: " + c.Param("userID")})
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func GetUsersWithVerificationStatus(c *gin.Context) {
	result := service.GetUsersWithVerificationStatus(c.Param("status"))
	c.JSON(http.StatusOK, result)
}

func CreateUser(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// Set the user id to ensure that the user can only modify their own account
	input.ID = c.Param("userID")
	if err := service.CreateUser(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetUserByID(input.ID))
}
