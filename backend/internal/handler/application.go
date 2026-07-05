package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"recruitment-platform/internal/model"
	"recruitment-platform/internal/store"
)

func RegisterApplicationRoutes(api *gin.RouterGroup) {
	apps := api.Group("/applications")
	{
		apps.GET("", ListApplications)
		apps.GET("/:id", GetApplication)
		apps.POST("", CreateApplication)
		apps.PUT("/:id/status", UpdateApplicationStatus)
	}
}

func ListApplications(c *gin.Context) {
	jobID := c.Query("job_id")
	candidateID := c.Query("candidate_id")
	status := c.Query("status")
	apps := store.ListApplications(jobID, candidateID, status)
	c.JSON(http.StatusOK, gin.H{"data": apps})
}

func GetApplication(c *gin.Context) {
	id := c.Param("id")
	app, ok := store.GetApplication(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "投递记录不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": app})
}

func CreateApplication(c *gin.Context) {
	var app model.Application
	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	if app.JobID == "" || app.CandidateID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "职位ID、候选人ID为必填项"})
		return
	}
	id, err := store.CreateApplication(app)
	if err != nil {
		if err == store.ErrDuplicateApplication {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else if err == store.ErrJobNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else if err == store.ErrJobClosed {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "投递成功"})
}

func UpdateApplicationStatus(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	if !store.UpdateApplicationStatus(id, req.Status) {
		c.JSON(http.StatusNotFound, gin.H{"error": "投递记录不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "状态更新成功"})
}