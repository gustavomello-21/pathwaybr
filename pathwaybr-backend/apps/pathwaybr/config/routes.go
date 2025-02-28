package config

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.GET("/health-check", controllers.NewHealthCheckController().Get)

	v1Api := r.Group("/api/v1")
	{
		v1Api.GET("/login")
	}
	return r
}
