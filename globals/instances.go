package globals

import (
	"leafwearz/controllers/repos"
	"leafwearz/controllers/serializers"
	"leafwearz/controllers/services"
	"leafwearz/controllers/validators"
)

var (
	SerializerService serializers.Serializers
	AuthValidatorService validators.AuthValidators
	JwtService validators.JWTService
	InsertionRepo repos.Creation
	FetchingRepo repos.Reading
	EditionRepo repos.Updation
	DeletionRepo repos.Deletion
)
var (
	AuthService services.AuthServices
)