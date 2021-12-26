package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/rikuhatano09/movieshare_api/internal/interfaces/handler"
)

func main() {
	engine := gin.Default()

	// Set CORS config.
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		"http://localhost",
	}
	corsConfig.AllowMethods = []string{
		"GET",
		"POST",
		"PUT",
		"OPTIONS",
	}
	engine.Use(cors.New(corsConfig))

	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello MovieShare",
		})
	})
	engine.GET("/movies/random", handler.GetMovieAtRandom)
	engine.GET("/movies", handler.GetMovieList)

	engine.Run(":8000")
}
