package config

import (
	"github.com/gin-gonic/gin"
	handler "github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers"
	activity "github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/activity/v1"
	auth "github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/auth/v1"
	intinerary "github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/intinerary/v1"
	trip "github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/trip/v1"
)

func Routes(controllers []interface{}) *gin.Engine {
	r := gin.Default()

	r.GET("/health-check", handler.NewHealthCheckController().Get)

	v1Api := r.Group("/api/v1")
	{
		for _, controller := range controllers {
			switch c := controller.(type) {
			case *auth.SessionController:
				v1Api.POST("/login", c.Create)
			case *auth.RegisterController:
				v1Api.POST("/register", c.Create)
			case *trip.TripController:
				v1Api.POST("/trip", c.Create)
				v1Api.GET("/users/:user_id/trips", c.Index)
				v1Api.GET("/trips/:trip_id", c.Show)
			case *intinerary.IntineraryController:
				v1Api.POST("/trips/:trip_id/itineraries", c.Create)
			case *activity.ActivityController:
				v1Api.POST("/intineraries/:intinerary_id/activities", c.Create)
			}
		}
	}
	return r
}
