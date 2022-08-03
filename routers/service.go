package routers

import "github.com/gin-gonic/gin"

type ServiceRoutes struct {}

func (r *ServiceRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("service")
	{
		withMiddleware.GET("", )
	}
}