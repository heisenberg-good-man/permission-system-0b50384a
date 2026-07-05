package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"recruitment-platform/internal/model"
	"recruitment-platform/internal/store"
)

func RegisterJobRoutes(api *gin.RouterGroup) {
	jobs := api.Group("/jobs")
	{
		jobs.GET("", ListJobs)
		jobs.GET("/:id", GetJob)
		jobs.POST("", CreateJob)
		jobs.PUT("/:id", UpdateJob)
		jobs.DELETE("/:id", DeleteJob)
	}
}

func ListJobs(c *gin.Context) {
	status := c.Query("status")
	jobs := store.ListJobs(status)
	c.JSON(http.StatusOK, gin.H{"data": jobs})
}

func GetJob(c *gin.Context) {
	id := c.Param("id")
	job, ok := store.GetJob(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "职位不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": job})
}

func CreateJob(c *gin.Context) {
	var job model.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	if job.Title == "" || job.Company == "" || job.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "职位名称、公司、描述为必填项"})
		return
	}
	if job.Status == "" {
		job.Status = "open"
	}
	if job.RecruiterID == "" {
		job.RecruiterID = "recruiter-1"
	}
	id := store.CreateJob(job)
	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "职位创建成功"})
}

func UpdateJob(c *gin.Context) {
	id := c.Param("id")
	var job model.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	if !store.UpdateJob(id, job) {
		c.JSON(http.StatusNotFound, gin.H{"error": "职位不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "职位更新成功"})
}

func DeleteJob(c *gin.Context) {
	id := c.Param("id")
	if !store.DeleteJob(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "职位不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "职位删除成功"})
}