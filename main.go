package main

import (
	"github.com/gin-gonic/gin"
	"shnk.com/eventx/db"
	"shnk.com/eventx/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
