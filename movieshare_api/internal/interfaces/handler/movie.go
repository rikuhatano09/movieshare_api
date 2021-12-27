package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rikuhatano09/movieshare_api/pkg/domain/contract"
	"github.com/rikuhatano09/movieshare_api/pkg/usecase"
)

func GetMovieAtRandom(context *gin.Context) {
	genre := convertStringToStringPointer(context.Query("genre"))

	movie, error := usecase.GetMovieAtRandom(genre)
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
	userId := convertStringToStringPointer(context.Query("userId"))
	movieList, error := usecase.GetMovieList(title, genre, userId)
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
	movie, error := usecase.GetMovieByID(id)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Internal server error: %s", error.Error()),
		})
		return
	}
	context.JSON(http.StatusOK, movie)
}

func CreateMovie(context *gin.Context) {
	requestBody := contract.MoviePostRequestBody{}

	error := context.ShouldBindJSON(&requestBody)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Bad request error: %s", error.Error()),
		})
		return
	}

	movie, error := usecase.CreateMovie(requestBody)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Internal server error: %s", error.Error()),
		})
		return
	}
	context.JSON(http.StatusOK, movie)
}

func PutMovie(context *gin.Context) {
	id, error := strconv.ParseUint(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Bad request error: %s", error.Error()),
		})
		return
	}
	requestBody := contract.MoviePutRequestBody{}

	error = context.ShouldBindJSON(&requestBody)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Bad request error: %s", error.Error()),
		})
		fmt.Println("aa")
		return
	}

	movie, error := usecase.PutMovie(requestBody, id)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Internal server error: %s", error.Error()),
		})
		return
	}
	context.JSON(http.StatusOK, movie)

}
