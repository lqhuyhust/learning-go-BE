package controllers

import (
	"httpServer/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PingController struct {
	PingService *services.PingService
}

func NewPingController(pingService *services.PingService) *PingController {
	return &PingController{
		PingService: pingService,
	}
}

func (ctrl *PingController) Ping(c *gin.Context) {
	userID := c.GetUint("user_id")

	// uint to string
	userIDValue := strconv.FormatUint(uint64(userID), 10)

	err := ctrl.PingService.HandlePing(userIDValue)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (ctrl *PingController) GetPingCount(c *gin.Context) {
	userID := c.Param("id")

	count := ctrl.PingService.GetPingCount(userID)

	// Get approximate user count
	userCount, err := ctrl.PingService.GetApproximateUserCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":                userID,
		"count":                  count,
		"approximate_user_count": userCount,
	})
}

func (ctrl *PingController) GetTopUsers(c *gin.Context) {
	topUsers, err := ctrl.PingService.GetTopUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"top_users": topUsers,
	})
}
