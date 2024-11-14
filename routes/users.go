package routes

import (
	"event-booking-rest-api/models"
	"event-booking-rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Something is wrong with your request", "error": err.Error()})
		return
	}
	
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save your request", "error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully created"})
}

func login(c *gin.Context){
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something is wrong with your request", "error": err.Error()})
	}

	err = user.ValidateUser()

	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.User_email, user.ID)

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authnenticate user"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message" : "Login Successfull", "token": token})

}
