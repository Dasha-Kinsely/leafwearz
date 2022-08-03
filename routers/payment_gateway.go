package routers

import "github.com/gin-gonic/gin"

type PaymentGatewayRoutes struct {}

func (r *PaymentGatewayRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("payment-gateway")
	{
		withMiddleware.GET("", )
	}
}