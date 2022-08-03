package routers

import "github.com/gin-gonic/gin"

type ProductRoutes struct {}

func (r *ProductRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("product")
	{
		withMiddleware.GET("", )
	}
}