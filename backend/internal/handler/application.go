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
		apps.PUT("/:id/offer", SendOffer)
		apps.PUT("/:id/offer/accept", AcceptOffer)
		apps.PUT("/:id/offer/reject", RejectOffer)
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

func SendOffer(c *gin.Context) {
	id := c.Param("id")
	var offer model.OfferDetail
	if err := c.ShouldBindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	if offer.OfferSalaryMin <= 0 || offer.OfferSalaryMax <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "薪资范围为必填项"})
		return
	}
	if offer.StartDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "到岗时间为必填项"})
		return
	}
	if !store.SendOffer(id, offer) {
		c.JSON(http.StatusNotFound, gin.H{"error": "投递记录不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Offer发送成功"})
}

func AcceptOffer(c *gin.Context) {
	id := c.Param("id")
	if !store.AcceptOffer(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法接受Offer，状态不允许"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Offer接受成功"})
}

func RejectOffer(c *gin.Context) {
	id := c.Param("id")
	if !store.RejectOffer(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法拒绝Offer，状态不允许"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Offer已拒绝"})
}