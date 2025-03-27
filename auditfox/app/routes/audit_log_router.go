/*
AuditFox - Open Source Audit Logging Solution
Copyright (C) 2025 Danzopen by Daniel Barbosa
*/

package routes

import (
	"auditfox/models"

	"github.com/gin-gonic/gin"
)

const auditAddPath = "/auditlog/add"

func SetupRouters() *gin.Engine {
	router := gin.Default()
	postAuditLog(router)
	healthCheck(router)
	return router
}

func postAuditLog(router *gin.Engine) *gin.Engine {
	router.POST(auditAddPath, func(c *gin.Context) {
		var auditLog models.AuditLogRecord
		c.BindJSON(&auditLog)
		c.JSON(200, auditLog)
	})
	return router
}
