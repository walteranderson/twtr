package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/walteranderson/twtr/storage"
	"github.com/walteranderson/twtr/types"
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

	router.Run(s.listenAddr)
}

func (s *Server) helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}

func (s *Server) getAllPosts(c *gin.Context) {
	posts, err := s.store.GetAllPosts()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(posts)
	c.JSON(http.StatusOK, posts)
}

func (s *Server) createPost(c *gin.Context) {
	var post types.Post

	if err := c.BindJSON(&post); err != nil {
		log.Fatal(err)
	}

  newPost, err := s.store.CreatePost(post)
  if err != nil {
    log.Fatal(err)
  }

  log.Println(newPost)
}
