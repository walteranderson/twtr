package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/walteranderson/twtr/storage"
	"github.com/walteranderson/twtr/types"
)

func handleError(status int, message string, c *gin.Context) {
  c.JSON(status, gin.H{"status": status, "message": message})
}

func (s *Server) helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
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
