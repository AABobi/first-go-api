package middlewares

import (
	"example/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		//TO z GIN metoda
		// No other code run after that
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No active token"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	context.Set("userId", userId)
	//Ensure that the next request handler in line will execute correctly
	context.Next()
}
