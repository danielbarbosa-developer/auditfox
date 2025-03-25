/*
AuditFox - Open Source Audit Logging Solution
Copyright (C) 2025 Danzopen by Daniel Barbosa
*/

package routes

import (
	"auditfox/models"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	return gin.Default()
}

func postAuditLog(router *gin.Engine) *gin.Engine {
	router.POST("/auditlog/add", func(c *gin.Context) {
		var auditLog models.AuditLogRecord
		c.BindJSON(&auditLog)
		c.JSON(200, auditLog)
	})
	return router
}
