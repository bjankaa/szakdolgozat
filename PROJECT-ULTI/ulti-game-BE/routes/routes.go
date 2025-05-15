package routes

import "github.com/gin-gonic/gin"

func AvaibleRoutes(server *gin.Engine) {

	server.POST("/auth", authenticationMethod)
	server.GET("/frontpage", getUsers)
}
