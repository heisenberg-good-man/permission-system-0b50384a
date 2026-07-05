package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"recruitment-platform/internal/model"
	"recruitment-platform/internal/store"
)

func RegisterCandidateRoutes(api *gin.RouterGroup) {
	candidates := api.Group("/candidates")
	{
		candidates.GET("", ListCandidates)
		candidates.GET("/:id", GetCandidate)
		candidates.POST("", CreateCandidate)
		candidates.PUT("/:id", UpdateCandidate)
	}
}

func ListCandidates(c *gin.Context) {
	candidates := store.ListCandidates()
	c.JSON(http.StatusOK, gin.H{"data": candidates})
}

func GetCandidate(c *gin.Context) {
	id := c.Param("id")
	candidate, ok := store.GetCandidate(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "候选人不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": candidate})
}

func CreateCandidate(c *gin.Context) {
	var candidate model.Candidate
	if err := c.ShouldBindJSON(&candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	if candidate.Name == "" || candidate.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "姓名、邮箱为必填项"})
		return
	}
	id := store.CreateCandidate(candidate)
	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "候选人创建成功"})
}

func UpdateCandidate(c *gin.Context) {
	id := c.Param("id")
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	if !store.UpdateCandidatePartial(id, updates) {
		c.JSON(http.StatusNotFound, gin.H{"error": "候选人不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "候选人更新成功"})
}