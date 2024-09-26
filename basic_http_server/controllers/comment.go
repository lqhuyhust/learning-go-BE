package controllers

import (
	"httpServer/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	CommentService *services.CommentService
}

func NewCommentController(service *services.CommentService) *CommentController {
	return &CommentController{
		CommentService: service,
	}
}

// Create a new comment
func (ctrl *CommentController) CreateComment(c *gin.Context) {
	var input struct {
		PostID   uint   `json:"post_id" binding:"required"`
		Content  string `json:"content" binding:"required"`
		ParentID *uint  `json:"parent_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	comment, err := ctrl.CommentService.CreateComment(input.PostID, userID, input.Content, input.ParentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)
}

// Get all comment for a post
func (ctrl *CommentController) GetComments(c *gin.Context) {
	postIDStr := c.Param("id")

	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	comments, err := ctrl.CommentService.GetComments(uint(postID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// Update a comment (only author can update)
func (ctrl *CommentController) UpdateComment(c *gin.Context) {
	commentIDParam := c.Param("comment_id")

	// Chuyển đổi commentID từ chuỗi sang uint
	commentID, err := strconv.ParseUint(commentIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var input struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id") // Lấy ID người dùng từ JWT

	// Gọi service để sửa comment
	comment, err := ctrl.CommentService.UpdateComment(userID, uint(commentID), input.Content)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// Delete a comment (only author can delete)
func (ctrl *CommentController) DeleteComment(c *gin.Context) {
	commentIDParam := c.Param("comment_id")

	// Chuyển đổi commentID từ chuỗi sang uint
	commentID, err := strconv.ParseUint(commentIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	userID := c.GetUint("user_id") // Lấy ID người dùng từ JWT

	// Gọi service để xóa comment
	err = ctrl.CommentService.DeleteComment(userID, uint(commentID))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
