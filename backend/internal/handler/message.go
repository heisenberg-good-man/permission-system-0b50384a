package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"recruitment-platform/internal/model"
	"recruitment-platform/internal/store"
)

func RegisterMessageRoutes(api *gin.RouterGroup) {
	msgs := api.Group("/messages")
	{
		msgs.GET("", ListMessages)
		msgs.POST("", CreateMessage)
	}
}

func ListMessages(c *gin.Context) {
	applicationID := c.Query("application_id")
	if applicationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "application_id 为必填项"})
		return
	}
	msgs := store.ListMessages(applicationID)
	c.JSON(http.StatusOK, gin.H{"data": msgs})
}

func CreateMessage(c *gin.Context) {
	var msg model.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	if msg.ApplicationID == "" || msg.SenderID == "" || msg.SenderType == "" || msg.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "application_id、sender_id、sender_type、content 为必填项"})
		return
	}
	id := store.CreateMessage(msg)
	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "消息发送成功"})
}