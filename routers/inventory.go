package routers

import "github.com/gin-gonic/gin"

type InventoryRoutes struct {}

func (r *InventoryRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("inventory")
	{
		withMiddleware.GET("", )
	}
}