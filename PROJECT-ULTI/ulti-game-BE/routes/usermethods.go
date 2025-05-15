package routes

import (
	"errors"
	"fmt"
	"net/http"

	"exmaple.com/ulti-restapi/models"
	"exmaple.com/ulti-restapi/utility"
	"github.com/gin-gonic/gin"
)

func authenticationMethod(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user data the first occurence"})
	}

	if user.State == "signup" {
		err = signup(user)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnot save user"})
			return
		}
		context.JSON(http.StatusCreated, gin.H{"message": "User saved.", "user": user})

	} else if user.State == "login" {

		token, err := login(user)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"message:": "Failed to login in"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Logged in lets ggoooo", "token": token})

	} else {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unrecognisable data."})
	}
}

func signup(user models.User) error {

	err := user.Save()
	if err != nil {
		err = errors.New("credentials is invalid, at signup")
		return err
	}
	return err

}

func login(user models.User) (string, error) {

	err := user.ValidateUser()

	if err != nil {
		err = errors.New("credentials is invalid at login")
		fmt.Println("vagyide eljuttotam", err)
		return "", err
	}

	token, err := utility.GenerateToken(user.Email, user.ID)

	if err != nil {
		err = errors.New("credentials is invalid didn't get token")
		fmt.Println("másidae eljuttotam", token, err)
		return "", err
	}
	fmt.Println("ide eljuttotam", token, err)
	return token, nil
}

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch"})
		return
	}
	context.JSON(http.StatusOK, users)
}
