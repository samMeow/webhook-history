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

// DummyResponse just dummy yes
type DummyResponse struct {
	Success bool
}

type MainRouter struct {
	DB Storage
}

func (m MainRouter) QueryHandler(c *gin.Context) {
	m.DB.Add(
		c.Request.Host,
		c.Request.URL.Query(),
	)
	res := &DummyResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, res)
}

func (m MainRouter) GetHistoryHandler(c *gin.Context) {
	history := m.DB.GetAll()
	c.JSON(http.StatusOK, history)
}

func (m MainRouter) ClearHandler(c *gin.Context) {
	m.DB.Clear()
	res := &DummyResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, res)
}

func main() {
	store := []RequestHistory{}
	mainRoute := &MainRouter{
		DB: &StorageImpl{
			Store: store,
		},
	}
	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/", mainRoute.QueryHandler)
	router.GET("/history", mainRoute.GetHistoryHandler)
	router.DELETE("/", mainRoute.ClearHandler)
	router.Run(":" + port)
}
