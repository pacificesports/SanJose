package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"sanjose/service"
	"sanjose/utils"
	"strings"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/sj/ping", Ping)
	router.GET("/users", GetAllUsers)
	router.GET("/users/:userID", GetUserByID)
	router.POST("/users/:userID", CreateUser)
	router.GET("/users/:userID/roles", GetRolesForUser)
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.SugarLogger.Infoln("GATEWAY REQUEST ID: " + c.GetHeader("Request-ID"))
		c.Next()
	}
}

func AuthChecker() gin.HandlerFunc {
	return func(c *gin.Context) {

		var requestUserID string
		var requestUserRoles []string

		ctx := context.Background()
		client, err := service.FirebaseAdmin.Auth(ctx)
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}
		if c.GetHeader("Authorization") != "" {
			token, err := client.VerifyIDToken(ctx, strings.Split(c.GetHeader("Authorization"), "Bearer ")[1])
			if err != nil {
				utils.SugarLogger.Errorln("error verifying ID token")
				requestUserID = "null"
			} else {
				utils.SugarLogger.Infoln("Decoded User ID: " + token.UID)
				requestUserID = token.UID
				requestUserRoles = service.GetRolesForUser(requestUserID)
			}
		} else {
			utils.SugarLogger.Infoln("No user token provided")
			requestUserID = "null"
		}
		c.Set("userID", requestUserID)
		// The main authentication gateway per request path
		// The requesting user's ID and roles are pulled and used below
		// Any path can also be quickly halted if not ready for prod
		if c.FullPath() == "/users/:userID" {
			// Creating or modifying a user requires the requesting user
			// to have a matching user ID or the ADMIN role
			if c.Request.Method == "POST" {
				if requestUserID != c.Param("userID") && !contains(requestUserRoles, "ADMIN") {
					//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You do not have permission to edit this resource"})
				}
			}
		}
		c.Next()
	}
}

func contains(s []string, element string) bool {
	for _, i := range s {
		if i == element {
			return true
		}
	}
	return false
}
