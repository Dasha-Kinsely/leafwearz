package startup

import (
	"leafwearz/controllers/repos"
	"leafwearz/controllers/serializers"
	"leafwearz/controllers/services"
	"leafwearz/controllers/validators"
	"leafwearz/globals"
	"sync"
)

func InitResources(wg *sync.WaitGroup) {
	defer wg.Done()
	// serializers
	globals.SerializerService = serializers.NewSerializers()
	// validators
	globals.AuthValidatorService = validators.NewAuthValidators(globals.SerializerService)
	globals.JwtService = validators.NewJWTService(globals.ServerGlobalSetting.JwtSecret)
	// repos
	globals.InsertionRepo = repos.NewCreationRepo(globals.MySqlClient)
	globals.FetchingRepo = repos.NewReadingRepo(globals.MySqlClient)
	globals.EditionRepo = repos.NewUpdationRepo(globals.MySqlClient)
	globals.DeletionRepo = repos.NewDeletionRepo(globals.MySqlClient)
	// services
	globals.AuthService = services.NewAuthServices(globals.InsertionRepo, globals.FetchingRepo, globals.EditionRepo)
}