package routers

import "github.com/gin-gonic/gin"

type AuthRoutes struct {}

func (r *AuthRoutes) InitRoutes(server *gin.RouterGroup) *gin.RouterGroup {
	withMiddleware := server.Group("/auth")
	{
		withMiddleware.GET("/signup", )
		withMiddleware.GET("/signin", )
		//withMiddleware.GET("/signout", )
	}
	return withMiddleware
}