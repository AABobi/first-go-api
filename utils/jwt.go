package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	//SigningMethodHS256 typ HMAC sprawdzany poniżej
	fmt.Println("GEneateTOken", userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), /*ten token jest wazny 2 h*/
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		//Sprawdzamy czy token jest takiego typu
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signin method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New(("Invalid token"))
	}
	fmt.Println(tokenIsValid)
	//Access to MapClaims ( above ) field
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	fmt.Println("claims", claims)
	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	//te nawiasy na końcu to type check
	//email := claims["email"].(string)
	fmt.Println(claims["userId"])
	var test1 float64 = claims["userId"].(float64)
	userId := int64(test1)
	return userId, nil
}
