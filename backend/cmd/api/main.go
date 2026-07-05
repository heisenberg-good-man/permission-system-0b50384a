package main

import (
	"recruitment-platform/internal/handler"
	"recruitment-platform/internal/middleware"
	"recruitment-platform/internal/store"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	store.InitMockData()

	api := r.Group("/api")
	{
		handler.RegisterJobRoutes(api)
		handler.RegisterCandidateRoutes(api)
		handler.RegisterApplicationRoutes(api)
		handler.RegisterMessageRoutes(api)
		handler.RegisterStatsRoutes(api)
	}

	r.Run(":8080")
}