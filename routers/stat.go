package routers

import "github.com/gin-gonic/gin"

type StatRoutes struct {}

func (r *StatRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("/stat")
	{
		withMiddleware.GET("", )
	}
}