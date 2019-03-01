package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		log.Panic("$PORT must be set")
	}
}

type DummyResponse struct {
	Success bool
}

func handler(c *gin.Context) {
	res := &DummyResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, res)
}

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/", handler)
	router.Run(":" + port)
}
