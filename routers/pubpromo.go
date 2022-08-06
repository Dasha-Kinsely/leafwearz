package routers

import "github.com/gin-gonic/gin"

type PubPromoRoutes struct {}

func (r *PubPromoRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("/promo")
	{
		withMiddleware.GET("/rules/:id", )
		withMiddleware.GET("", )
	}
}