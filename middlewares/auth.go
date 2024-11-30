package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shnk.com/eventx/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please login to create an event"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please login with right credentials to create an event", "error": err.Error()})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
