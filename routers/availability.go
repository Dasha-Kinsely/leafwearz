package routers

import "github.com/gin-gonic/gin"

type AvailabilityRoutes struct {}

func (r *AvailabilityRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("/availability")
	{
		withMiddleware.GET("", )
	}
}