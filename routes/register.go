package routes

import (
	"event-booking-rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerEvent(context *gin.Context) {
	userId := context.GetInt64("usedId")
	
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Some thing is wrong with your request", "error": err.Error()})
		return
	}

	event, err := models.GetEventbyID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Regisration successfull!"})
}

func deleteRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.Id = eventId

	err = event.Cancel(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete registration."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully deleted!"})
}