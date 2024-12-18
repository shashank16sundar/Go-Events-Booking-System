package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shnk.com/eventx/models"
	"shnk.com/eventx/utils"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user details"})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not add user to the database"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Successfully added user to the database"})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user details"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Could not login"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully logged in!", "token": token})
}
