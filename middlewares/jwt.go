package middlewares

import (
	"leafwearz/controllers/serializers"
	"leafwearz/controllers/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

// The argument must first be initialized at "setup/routes/routes.go" to initialize its struct values
func AuthorizeJWT(j validators.JWTService, serializer serializers.Serializers) gin.HandlerFunc {
	return func(c *gin.Context) {
		// fetch the auth header received from frontend requests
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			responseJSON := serializer.ErrorResponse("Failed to process jwt request", "no token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, responseJSON)
			return
		}
		// validate it to make sure that the user has signed-in before requesting the contents it is protecting
		// This is merely a function caller: 1) nil means there is an error parsing jwt, 2) invalid token means wrong credential
		token := j.ValidateToken(authHeader, c)
		if !token.Valid || token == nil {
			response := serializer.ErrorResponse("JWT error", "Your token is not valid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}