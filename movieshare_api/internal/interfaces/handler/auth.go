package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rikuhatano09/movieshare_api/pkg/domain/contract"
)

// LoginHandler handles login requests
func LoginHandler(context *gin.Context) {
	requestBody := contract.LoginPostRequestBody{}

	err := context.ShouldBindJSON(requestBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Bad request error: %s", err.Error()),
		})
		return
	}
}