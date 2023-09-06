package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sanjose/config"
	"sanjose/model"
	"sanjose/service"
)

func GetVerificationForUser(c *gin.Context) {
	result := service.GetVerificationForUser(c.Param("userID"))
	if result.UserID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "No verification found for user with given id: " + c.Param("userID")})
		return
	}
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
	user := service.GetUserByID(c.Param("userID"))
	go service.Discord.ChannelMessageSend(config.DiscordChannel, "<@"+user.ID+"> "+user.FirstName+" "+user.LastName+"'s verification status is now `"+input.Status+"`")
	if input.Status == "REQUESTED" {
		go service.DiscordLogUserVerificationRequested(user)
	} else if input.Status == "ACCEPTED" {
		go service.DiscordLogUserVerificationApproved(user)
	}
	c.JSON(http.StatusOK, service.GetVerificationForUser(c.Param("userID")))
}
