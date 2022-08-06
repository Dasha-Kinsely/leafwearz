package validators

import (
	"leafwearz/models/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthValidators interface {
	Validate(c *gin.Context, err error, step string) bool
}

func (validator *AuthValidatorsResources) Validate(c *gin.Context, err error, step string) bool {
	empty := entities.Empty{}
	switch step {
	case "signupForm":
		if err != nil {
			res := validator.Serializers.ErrorResponse("sign up form is in an invalid format", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return false
		} else {
			return true
		}
	case "createUser":
		if err != nil {
			res := validator.Serializers.ErrorResponse("problem occurred while signing up user", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return false
		} else {
			return true
		}
	case "signinForm":
		if err != nil {
			res := validator.Serializers.ErrorResponse("sign in form is in an invalid format", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return false
		} else {
			return true
		}
	case "signinCred" :
		if err != nil {
			res := validator.Serializers.ErrorResponse("problem occurred while signing in user", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return false
		} else {
			return true
		}
	default:
		if err != nil {
			res := validator.Serializers.ErrorResponse("this is an unknown error, please see error body for details", err.Error(), empty)
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return false
		} else {
			return true
		}
	}
}