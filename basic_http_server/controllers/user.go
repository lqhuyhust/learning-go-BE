package controllers

import (
	"httpServer/services"
	"httpServer/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	userID := c.GetUint("user_id")

	user, err := ctrl.UserService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	userID := c.GetUint("user_id")

	// upload profile image
	profile, err := utils.UploadFile(c, "profile", "uploads")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image"})
		return
	}

	err = ctrl.UserService.UpdateUser(userID, profile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "update success"})
}
