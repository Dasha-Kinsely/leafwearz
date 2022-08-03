package routers

import "github.com/gin-gonic/gin"

type AuthRoutes struct {}

func (r *AuthRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("auth")
	{
		withMiddleware.GET("", )
	}
}