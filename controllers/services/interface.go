package services

import (
	"leafwearz/models/dto"
	"leafwearz/models/entities"
)

type AuthServices interface {
	CreateUser(signupRequest dto.SignupRequest) (entities.User, error)
	SigninUser(email, password string) (uint, error)
}