package controllers

import (
	"httpServer/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	PostService *services.PostService
}

func NewPostController(postService *services.PostService) *PostController {
	return &PostController{
		PostService: postService,
	}
}

// Create a new post
func (ctrl *PostController) CreatePost(c *gin.Context) {
	var input struct {
		PostTitle   string `json:"post_title" binding:"required"`
		PostContent string `json:"post_content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")
	post, err := ctrl.PostService.CreatePost(userID, input.PostTitle, input.PostContent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

// Show all posts
func (ctrl *PostController) ShowPosts(c *gin.Context) {
	posts, err := ctrl.PostService.ShowPosts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// Show post by ID
func (ctrl *PostController) ShowPostByID(c *gin.Context) {
	postIDParam := c.Param("id")
	postID, err := strconv.ParseUint(postIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := ctrl.PostService.ShowPostByID(uint(postID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

// Update post by ID
func (ctrl *PostController) UpdatePostByID(c *gin.Context) {
	postIDParam := c.Param("id")
	postID, err := strconv.ParseUint(postIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var input struct {
		Title   *string `json:"title"`
		Content *string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	post, err := ctrl.PostService.UpdatePost(userID, uint(postID), input.Title, input.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

// Delete post by ID
func (ctrl *PostController) DeletePostByID(c *gin.Context) {
	postIDParam := c.Param("id")
	postID, err := strconv.ParseUint(postIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	userID := c.GetUint("user_id")

	err = ctrl.PostService.DeletePostByID(uint(postID), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
