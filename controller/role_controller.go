package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sanjose/service"
)

func GetRolesForUser(c *gin.Context) {
	result := service.GetRolesForUser(c.Param("userID"))
	c.JSON(http.StatusOK, result)
}

func SetRolesForUser(c *gin.Context) {
	var input []string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := service.SetRolesForUser(c.Param("userID"), input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetRolesForUser(c.Param("userID")))
}
