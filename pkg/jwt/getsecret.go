package jwt

import "leafwearz/globals"

func GetSecret() string {
	return globals.ServerGlobalSetting.JwtSecret
}