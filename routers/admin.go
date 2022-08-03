package routers

import "github.com/gin-gonic/gin"

type AdminRoutes struct {}

func (r *AdminRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("admin")
	{
		withMiddleware.GET("", )
	}
}