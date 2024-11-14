package routes

import (
	"event-booking-rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)

}

func getEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Some thing is wrong with your request", "error": err.Error()})
		return
	}

	event, err := models.GetEventbyID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Something is wrong with your request", "error": err.Error()})
		return
	}

	userId := context.GetInt64("usedId")
	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Some thing is wrong with your request", "error": err.Error()})
		return
	}

	userId := context.GetInt64("usedId")
	event, err := models.GetEventbyID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "error": err.Error()})
		return
	}

	if userId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data", "error": err.Error()})
	}

	updatedEvent.Id = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update data", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Some thing is wrong with your request", "error": err.Error()})
		return
	}

	event, err := models.GetEventbyID(eventId)
	userId := context.GetInt64("usedId")
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "error": err.Error()})
		return
	}
	
	if userId != event.UserID{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not unauthorized to delete the event"})
	} 
	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete data", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
