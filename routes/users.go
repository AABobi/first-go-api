package routes

import (
	"example/models"
	"example/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(context *gin.Context) {
	users, err := models.GetAllUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch"})
		return
	}
	context.JSON(http.StatusOK, users)
}

func signup(context *gin.Context) {
	var user models.User
	//Przychodzi call z frontendu, jakiś json post i ShouldBindJSON zapisujemy wartośc do event
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("beforeSave")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.Save()
	if err != nil {
		fmt.Println("afterSave")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User create corect", "user": user})

}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println("login")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not login"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		fmt.Println("logi2")
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authoticed"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	fmt.Println("test", token)
	if err != nil {
		fmt.Println("log3")
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authoticed"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login succes", "token": token})
}
