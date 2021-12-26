package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

func GetMovieList(context *gin.Context) {
	title := convertStringToStringPointer(context.Query("title"))
	genre := convertStringToStringPointer(context.Query("genre"))
	user_id := convertStringToStringPointer(context.Query("userId"))
	movieList, error := usecase.GetMovieList(title, genre, user_id)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Internal server error: %s", error.Error()),
		})
		return
	}
	context.JSON(http.StatusOK, movieList)
}

func GetMovieByID(context *gin.Context) {
	id, error := strconv.ParseUint(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Bad request error: %s", error.Error()),
		})
		return
	}
	movie, error := usecase.GetMovieID(id)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Internal server error: %s", error.Error()),
		})
		return
	}
	context.JSON(http.StatusOK, movie)
}
