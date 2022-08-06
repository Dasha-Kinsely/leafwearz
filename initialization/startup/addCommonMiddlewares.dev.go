package startup

import (
	"leafwearz/globals"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
)

func AddCommonMiddleware() {
	// Cors
	AllowOriginList := strings.Split(globals.CorsGlobalSetting.AllowOrigins, ",")
	AllowMethodsList := strings.Split(globals.CorsGlobalSetting.AllowMethods, ",")
	AllowHeadersList := strings.Split(globals.CorsGlobalSetting.AllowHeaders, ",")
	ExposeHeadersList := strings.Split(globals.CorsGlobalSetting.ExposeHeaders, ",")
	globals.DefaultServer.Use(cors.New(cors.Config{
		AllowOrigins:     AllowOriginList,
		AllowMethods:     AllowMethodsList,
		AllowHeaders:     AllowHeadersList,
		ExposeHeaders:    ExposeHeadersList,
		AllowCredentials: globals.CorsGlobalSetting.AllowCredentials,
		MaxAge: 12 * time.Hour,
  	}))
	// TODO: Swagger
	return
}