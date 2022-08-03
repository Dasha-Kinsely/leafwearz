package routers

import "github.com/gin-gonic/gin"

type RecommendationSettingRoutes struct {}

func (r *RecommendationSettingRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("recommendation")
	{
		withMiddleware.GET("", )
	}
}