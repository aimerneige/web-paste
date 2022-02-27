package route

import (
	"github.com/aimerneige/web-paste/controller"
	"github.com/aimerneige/web-paste/middleware"
	"github.com/gin-gonic/gin"
)

func AllRouteCollection(r *gin.Engine) *gin.Engine {
	r.Use(middleware.BlackListMiddleware())
	record := r.Group("/record")
	RecordRouteCollection(record)
	r.NoRoute(controller.NotFound)

	return r
}
