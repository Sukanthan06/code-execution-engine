package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Sukanthan06/code-execution-engine/api/handlers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", handlers.HealthCheck)
	router.POST("/execute", handlers.ExecuteCode)
	router.GET("/docker-test", handlers.DockerTest)
}
