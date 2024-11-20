package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"shnk.com/eventx/db"
	"shnk.com/eventx/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	//GET REQUESTS
	server.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "Wassup!"}) })
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	//POST REQUESTS
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	reqId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Not a valid input"})
		return
	}
	event, err := models.GetEventByID(reqId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Not a valid event"})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "Event queried successfully", "event": event})
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Body not properly parsed"})
		return
	}

	event.DateTime = time.Now()
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save data", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}
