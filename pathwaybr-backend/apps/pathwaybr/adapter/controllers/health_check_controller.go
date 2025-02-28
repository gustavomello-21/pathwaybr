package controllers

import (
	"github.com/gin-gonic/gin"
)

type HealthCheckController struct{}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

func (h *HealthCheckController) Get(httpContext *gin.Context) {
	httpContext.JSON(200, gin.H{
		"message": "Everything is fine!",
	})
}
