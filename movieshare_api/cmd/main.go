package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rikuhatano09/movieshare_api/internal/interfaces/handler"
)

func main() {
	engine := gin.Default()
	engine.GET("/movie", handler.GetMovie)

	engine.Run(":3000")
}
