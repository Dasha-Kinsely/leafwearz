package validators

import (
	"fmt"
	"leafwearz/models/entities"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(uid string) string
	ValidateToken(uid string, c *gin.Context) *jwt.Token
}
// This is bound to the interface and must not be placed in a different package
type JWTServiceResources struct {
	secretKey string
	issuer string
}
// Need to initialize this interface before calling it
func NewJWTService(secret string) JWTService {
	return &JWTServiceResources{
		issuer: "admin",
		secretKey: secret,
	}
}

func (j *JWTServiceResources) GenerateToken(uid string) string {
	claims := &entities.JWTCustomClaim{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: j.issuer,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return token
}

func (j *JWTServiceResources) ValidateToken(token string, c *gin.Context) *jwt.Token {
	t, err := jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("JWT validation failed at parsing step... %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		//fmt.Println(t)
		return nil
	}
	//fmt.Println(t)
	return t
}