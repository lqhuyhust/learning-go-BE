package controllers

import (
	"httpServer/services"
	"httpServer/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// upload profile image
	profile, err := utils.UploadFile(c, "profile", "uploads")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image"})
		return
	}

	err = ctrl.AuthService.Register(username, password, profile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "register success"})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := ctrl.AuthService.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// save access token to Cookie
	c.SetCookie("access_token", accessToken, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (ctrl *AuthController) Logout(c *gin.Context) {
	userID := c.GetUint("user_id")

	// delete refresh token
	if err := ctrl.AuthService.Logout(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// delete access token
	c.SetCookie("access_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}

func (ctrl *AuthController) RefreshToken(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := ctrl.AuthService.RefreshToken(input.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// save access token to Cookie
	c.SetCookie("access_token", accessToken, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
