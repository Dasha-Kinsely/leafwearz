package routers

import "github.com/gin-gonic/gin"

type UserRoutes struct {}

func (r *UserRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("user")
	{
		withMiddleware.GET("", )
	}
}