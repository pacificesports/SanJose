package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sanjose/model"
	"sanjose/service"
)

func GetVerificationForUser(c *gin.Context) {
	result := service.GetVerificationForUser(c.Param("userID"))
	c.JSON(http.StatusOK, result)
}

func SetVerificationForUser(c *gin.Context) {
	var input model.Verification
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := service.SetVerificationForUser(c.Param("userID"), input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if input.Status == "REJECTED" {

	}
	c.JSON(http.StatusOK, service.GetVerificationForUser(c.Param("userID")))
}
