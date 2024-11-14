package routes

import (
	"event-booking-rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authnenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerEvent)
	authenticated.DELETE("/events/:id/register", deleteRegistration)
	server.POST("/signup", signup)
	server.POST("/login", login)
}