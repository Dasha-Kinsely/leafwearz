package routers

import "github.com/gin-gonic/gin"

type PromoRoutes struct {}

func (r *PromoRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("promo")
	{
		withMiddleware.GET("", )
	}
}