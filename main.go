package main

import (
	"event-booking-rest-api/db"
	"event-booking-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	
	db.InitDb()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":5050")
}
