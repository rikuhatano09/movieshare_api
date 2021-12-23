package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rikuhatano09/movieshare_api/pkg/usecase"
)

func GetMovieAtRandom(context *gin.Context) {
	movie, error := usecase.GetMovieAtRandom()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Internal server error: %s", error.Error()),
		})
		return
	}
	context.JSON(http.StatusOK, movie)
}
