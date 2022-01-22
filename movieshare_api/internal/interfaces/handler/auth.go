package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rikuhatano09/movieshare_api/pkg/domain/contract"
	"github.com/rikuhatano09/movieshare_api/pkg/usecase"
)

// LoginHandler handles login requests
func LoginHandler(context *gin.Context) {
	err := usecase.CreateSessionCookie(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Bad request error: %s", err.Error()),
		})
		return 
	}
	context.JSON(http.StatusOK, contract.VerificationResponse{
		Status: "Succefully logged in the user.",
	})
	return
}