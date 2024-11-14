package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "superkey"

func GenerateToken(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_email": email,
		"ID" : userId,
		"exp" : time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretkey))
}


func ValidateToken (token string) (error, int64){
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token)(interface {}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok{
			return nil, errors.New("Unexpected Error")
		}
		return []byte(secretkey), nil
	})

	if err != nil{
		return errors.New("Could not parse Token"), 0
	} 

	isToken := parsedToken.Valid

	if !isToken{
		return errors.New("Invalid Token"), 0
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return errors.New("Invalid Token Claim"), 0
	}

	// email := claims["email"].(string)
	userId := int64(claims["ID"].(float64))

	return nil, userId
}