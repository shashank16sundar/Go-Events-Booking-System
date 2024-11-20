package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//GET REQUESTS
	server.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "Wassup!"}) })
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	//POST REQUESTS
	server.POST("/events", createEvent)
}
