package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"recruitment-platform/internal/store"
)

func RegisterStatsRoutes(api *gin.RouterGroup) {
	api.GET("/stats", GetStats)
}

func GetStats(c *gin.Context) {
	stats := store.GetStats()
	c.JSON(http.StatusOK, gin.H{"data": stats})
}