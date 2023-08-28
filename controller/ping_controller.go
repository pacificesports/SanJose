package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sanjose/config"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": config.Service.Name + " v" + config.Version + " is online!"})
}
