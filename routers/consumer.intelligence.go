package routers

import "github.com/gin-gonic/gin"

type ConsumerIntelligenceRoutes struct {}

func (r *ConsumerIntelligenceRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("/intelligence").Group("/consumer")
	{
		withMiddleware.GET("", )
	}
}