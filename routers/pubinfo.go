package routers

import "github.com/gin-gonic/gin"

type PubInfoRoutes struct {}

func (r *PubInfoRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("pub-info")
	{
		withMiddleware.GET("", )
	}
}