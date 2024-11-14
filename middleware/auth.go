package middleware

import (
	"event-booking-rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authnenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == ""{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return 
	}

	err, userId := utils.ValidateToken(token)

	if err != nil{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorazited"})
		return
	}
	
	context.Set("usedId", userId)
	context.Next()
}