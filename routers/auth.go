package routers

import (
	"leafwearz/globals"
	"leafwearz/handlers"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {}

func (r *AuthRoutes) InitRoutes(server *gin.RouterGroup) *gin.RouterGroup {
	handler := handlers.NewAuthHandler(globals.AuthService, globals.SerializerService, globals.AuthValidatorService, globals.JwtService)
	withMiddleware := server.Group("/auth")
	{
		withMiddleware.POST("/signup", handler.Signin)
		withMiddleware.POST("/signin", handler.Signup)
		//withMiddleware.GET("/signout", )
	}
	return withMiddleware
}