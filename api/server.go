package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/walteranderson/twtr/storage"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() {
	router := gin.Default()
	router.GET("/", s.helloWorld)

	router.Run(s.listenAddr)
}

func (s *Server) helloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World")
}
