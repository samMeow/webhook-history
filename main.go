package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultPort string = "8080"

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = defaultPort
		log.Printf("$PORT must be set fallback to defaultport %s\n", defaultPort)
	}
}

// DummyResponse just dummy yes
type DummyResponse struct {
	Success bool
}

type MainRouter struct {
	RequestRepository Storage
}

func (m MainRouter) QueryHandler(c *gin.Context) {
	params, err := json.Marshal(c.Request.URL.Query())
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	m.RequestRepository.Add(
		"GET",
		c.Request.Host,
		string(params),
	)
	res := &DummyResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, res)
}

func (m MainRouter) BodyHandler(c *gin.Context) {
	buf := bytes.Buffer{}
	_, err := buf.ReadFrom(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	m.RequestRepository.Add(
		"POST",
		c.Request.Host,
		buf.String(),
	)

	res := &DummyResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, res)
}

func (m MainRouter) GetHistoryHandler(c *gin.Context) {
	history := m.RequestRepository.GetAll()
	c.JSON(http.StatusOK, history)
}

func (m MainRouter) ClearHandler(c *gin.Context) {
	m.RequestRepository.Clear()
	res := &DummyResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, res)
}

func NewMainRouter(re Storage) *gin.Engine {
	mainRoute := &MainRouter{
		RequestRepository: re,
	}
	router := gin.Default()
	router.GET("/", mainRoute.QueryHandler)
	router.POST("/", mainRoute.BodyHandler)
	router.GET("/history", mainRoute.GetHistoryHandler)
	router.DELETE("/", mainRoute.ClearHandler)
	return router
}

func main() {
	store := []RequestHistory{}
	repo := &StorageImpl{
		Store: store,
	}
	router := NewMainRouter(repo)
	router.Use(gin.Logger())
	router.Run(":" + port)
}
