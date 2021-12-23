package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rikuhatano09/movieshare_api/internal/interfaces/handler"
)

func main() {
	engine := gin.Default()

	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello MovieShare",
		})
	})
	engine.GET("/movies/random", handler.GetMovieAtRandom)
	engine.GET("/movies", handler.GetMovieList)

	engine.Run(":8000")
}
