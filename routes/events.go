package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"shnk.com/eventx/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}
	context.JSON(http.StatusOK, events)
}

// INDIVIDUAL EVENT REQUESTS
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
	ctx.JSON(http.StatusOK, gin.H{"message": "Event queried successfully", "event": event})
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

func deleteEvent(ctx *gin.Context) {
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

	err = event.Delete()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in deleting the event"})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted event", "event": event})
}

func updateEvent(ctx *gin.Context) {
	reqId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Not a valid input"})
		return
	}
	_, err = models.GetEventByID(reqId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Not a valid event"})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Body not properly parsed"})
		return
	}
	updatedEvent.ID = reqId
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error in updating the event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully updated", "event": updatedEvent})
}
