package api

import (
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
	router.GET("/posts/:id", s.getPost)
	router.POST("/posts", s.createPost)

	router.Run(s.listenAddr)
}

func (s *Server) helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}

func (s *Server) getAllPosts(c *gin.Context) {
	posts, err := s.store.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "404", "message": "Internal Server Error"})
	}

	c.JSON(http.StatusOK, posts)
}

func (s *Server) getPost(c *gin.Context) {
	id := c.Param("id")
	post, err := s.store.GetPost(id)
	if err != nil {
		if err == storage.ErrNotExists {
			c.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Not Found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "404", "message": "Internal Server Error"})
			return
		}
	}

	c.JSON(http.StatusOK, post)
}

func (s *Server) createPost(c *gin.Context) {
	var post types.Post

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "404", "message": "Internal Server Error"})
	}

	newPost, err := s.store.CreatePost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "404", "message": "Internal Server Error"})
	}

	c.JSON(http.StatusOK, newPost)
}
