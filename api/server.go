package api

import (
	"log"
	"net/http"
	"strconv"

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
	router.LoadHTMLGlob("templates/*")
  router.Static("/static", "./static")
  router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.GET("/", s.renderIndex)

	api := router.Group("/api")
	{
		api.GET("/posts", s.getAllPosts)
		api.POST("/posts", s.createPost)
		api.GET("/posts/:id", s.getPost)
		api.PATCH("/posts/:id", s.updatePost)
	}

	router.Run(s.listenAddr)
}

func (s *Server) renderIndex(c *gin.Context) {
	posts, err := s.store.GetAllPosts()
  if err != nil {
    log.Println(err)
    return
  }

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
    "posts": posts,
	})
}

// JSON API

func handleError(status int, message string, c *gin.Context) {
	c.JSON(status, gin.H{"status": status, "message": message})
}

func (s *Server) getAllPosts(c *gin.Context) {
	posts, err := s.store.GetAllPosts()
	if err != nil {
		handleError(http.StatusInternalServerError, "Internal Server Error", c)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (s *Server) getPost(c *gin.Context) {
	id := c.Param("id")
	post, err := s.store.GetPost(id)
	if err != nil {
		if err == storage.ErrNotExists {
			handleError(http.StatusNotFound, "Not Found", c)
		} else {
			handleError(http.StatusInternalServerError, "Internal Server Error", c)
		}
		return
	}

	c.JSON(http.StatusOK, post)
}

func (s *Server) createPost(c *gin.Context) {
	var post types.Post

	if err := c.BindJSON(&post); err != nil {
		handleError(http.StatusBadRequest, "Bad Request", c)
		return
	}

	newPost, err := s.store.CreatePost(post)
	if err != nil {
		handleError(http.StatusInternalServerError, "Internal Server Error", c)
		return
	}

	c.JSON(http.StatusOK, newPost)
}

func (s *Server) updatePost(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handleError(http.StatusBadRequest, "Invalid ID", c)
		return
	}

	post := types.Post{ID: id}

	if err := c.BindJSON(&post); err != nil {
		handleError(http.StatusBadRequest, "Bad Request", c)
		return
	}

	updatedPost, err := s.store.UpdatePost(post)
	if err != nil {
		handleError(http.StatusInternalServerError, "Internal Server Error", c)
		return
	}

	c.JSON(http.StatusOK, updatedPost)
}
