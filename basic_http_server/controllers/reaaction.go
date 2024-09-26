package controllers

import (
	"httpServer/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReactionController struct {
	ReactionService *services.ReactionService
}

func NewReactionController(reactionService *services.ReactionService) *ReactionController {
	return &ReactionController{
		ReactionService: reactionService,
	}
}

// Make reaction for post
func (ctrl *ReactionController) MakeReaction(c *gin.Context) {
	var input struct {
		PostID         uint `json:"post_id" binding:"required"`
		ReactionTypeID uint `json:"reaction_type_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	err := ctrl.ReactionService.MakeReaction(input.ReactionTypeID, userID, input.PostID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "reaction success"})
}

// Delete reaction for post
func (ctrl *ReactionController) DeleteReaction(c *gin.Context) {
	userID := c.GetUint("user_id")
	postIDParam := c.Param("id")
	postID, err := strconv.ParseUint(postIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	err = ctrl.ReactionService.DeleteReaction(userID, uint(postID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
