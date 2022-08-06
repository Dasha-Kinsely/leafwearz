package startup

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
)

type Cors struct {
	Mode string `json:"mode" yaml:"mode"`
	AllowOrigins string `json:"allow_origin" yaml:"allow_origin"`
	AllowMethods string `json:"allow_methods" yaml:"allow_methods"`
	AllowHeaders string `json:"allow_headers" yaml:"allow_headers"`
	ExposeHeaders string `json:"expose_headers" yaml:"expose_headers"`
	AllowCredentials bool `json:"allow_credentials" yaml:"allow_credentials"`
}

var CorsGlobalSetting *Cors

func AddCommonMiddleware() {
	// Cors
	AllowOriginList := strings.Split(CorsGlobalSetting.AllowOrigins, ",")
	AllowMethodsList := strings.Split(CorsGlobalSetting.AllowMethods, ",")
	AllowHeadersList := strings.Split(CorsGlobalSetting.AllowHeaders, ",")
	ExposeHeadersList := strings.Split(CorsGlobalSetting.ExposeHeaders, ",")
	DefaultServer.Use(cors.New(cors.Config{
		AllowOrigins:     AllowOriginList,
		AllowMethods:     AllowMethodsList,
		AllowHeaders:     AllowHeadersList,
		ExposeHeaders:    ExposeHeadersList,
		AllowCredentials: CorsGlobalSetting.AllowCredentials,
		MaxAge: 12 * time.Hour,
  	}))
	// TODO: Swagger
	return
}