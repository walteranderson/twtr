package api

import (
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
	router.GET("/posts", s.getAllPosts)
	router.POST("/posts", s.createPost)
	router.GET("/posts/:id", s.getPost)
	router.PATCH("/posts/:id", s.updatePost)

	router.Run(s.listenAddr)
}
