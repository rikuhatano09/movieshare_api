package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rikuhatano09/movieshare_api/pkg/usecase"
)

func GetMovie(context *gin.Context) {
	movie := usecase.GetMovie()
	context.JSON(http.StatusOK, movie)
}
