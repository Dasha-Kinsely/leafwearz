package routers

import "github.com/gin-gonic/gin"

type UserProfileRoutes struct {}

func (r *UserProfileRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("/user")
	{
		withMiddleware.GET("", )
	}
}