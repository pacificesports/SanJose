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
