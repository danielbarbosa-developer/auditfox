package routes

import (
	"auditfox/models"

	"github.com/gin-gonic/gin"
)

const healthCheckPath = "/healthcheck"

func healthCheck(router *gin.Engine) *gin.Engine {

	router.GET(healthCheckPath, func(ctx *gin.Context) {
		healthCheck := checkStatus()
		ctx.JSON(200, healthCheck)
	})

	return router
}

func checkStatus() models.HealthCheck {
	return models.HealthCheck{
		Status: "RUNNING",
	}
}
