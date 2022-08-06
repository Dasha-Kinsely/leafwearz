package handlers

import (
	"fmt"
	"leafwearz/controllers/serializers"
	"leafwearz/controllers/services"
	"leafwearz/controllers/validators"
	"leafwearz/models/dto"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Signup(c *gin.Context)
	Signin(c *gin.Context)
}

type AuthHandlerResources struct {
	Services services.AuthServices
	Serializer serializers.Serializers
	Validator validators.AuthValidators
	JwtService validators.JWTService
}

func NewAuthHandler(
	services services.AuthServices,
	serializer serializers.Serializers,
	validator validators.AuthValidators,
	jwt validators.JWTService) AuthHandler {
	return &AuthHandlerResources{
		Services: services,
		Serializer: serializer,
		Validator: validator,
		JwtService: jwt,
	}
}

func (handler *AuthHandlerResources) Signin(c *gin.Context) {
	// check if the incoming form request is in valid format
	var signinRequest dto.SigninRequest
	err := c.ShouldBind(&signinRequest)
	if handler.Validator.Validate(c, err, "signinForm") {
		// verify user credentials
		userID, err := handler.Services.SigninUser(signinRequest.Email, signinRequest.Password) 
		if handler.Validator.Validate(c, err, "signinCred") {
			token := handler.JwtService.GenerateToken(strconv.FormatUint(uint64(userID), 10))
			c.Set("authenticated", token)
			c.Set("user", signinRequest.Email)
			log.Println(c.Get("authenticated"))
			log.Println(token)
			//responses.NewSigninSerializer(c, signinRequest.Email)
		}
	}
	return
}

func (handler *AuthHandlerResources) Signup(c *gin.Context) {
	var signupRequest dto.SignupRequest
	err := c.ShouldBind(&signupRequest)
	if handler.Validator.Validate(c, err, "signupForm") {
		user, err := handler.Services.CreateUser(signupRequest)
		if handler.Validator.Validate(c, err, "createUser") {
			//handler.Serializer.NewSignupSerializer(c, user)
			fmt.Println(user.CreatedAt)
			return
		}
	}
	return
}