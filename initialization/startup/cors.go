package startup

type Cors struct {
	Mode string `json:"mode" yaml:"mode"`
	Whitelist []CorsWhitelist `json:"whitelist" yaml:"whitelist"`
}

type CorsWhitelist struct {
	AllowOrigins string `json:"allow_origin" yaml:"allow_origin"`
	AllowMethods string `json:"allow_methods" yaml:"allow_methods"`
	AllowHeaders string `json:"allow_headers" yaml:"allow_headers"`
	ExposeHeaders string `json:"expose_headers" yaml:"expose_headers"`
	AllowCredentials bool `json:"allow_credentials" yaml:"allow_credentials"`
}

var CorsGlobalSetting *Cors