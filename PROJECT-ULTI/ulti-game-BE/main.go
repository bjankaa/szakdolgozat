package main

import (
	"exmaple.com/ulti-restapi/database"
	"exmaple.com/ulti-restapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDatabase()

	server := gin.Default()

	routes.AvaibleRoutes(server)

	server.Run(":3000")
}
