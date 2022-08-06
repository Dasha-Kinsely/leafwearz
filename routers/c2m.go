package routers

import "github.com/gin-gonic/gin"

type C2MCommunicationRoutes struct {}

func (r *C2MCommunicationRoutes) InitRoutes(Router *gin.RouterGroup) {
	withMiddleware := Router.Group("/communication")
	{
		withMiddleware.GET("", )
	}
}