package crypto

import (
	"intern_BCC/entity"
	"intern_BCC/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(payload entity.Customer) (string, error) {
	expStr := os.Getenv("JWT_EXP")
	var exp time.Duration
	exp, err := time.ParseDuration(expStr)
	if expStr == "" || err != nil {
		exp = time.Hour * 1
	}
	tokenJwtSementara := jwt.NewWithClaims(jwt.SigningMethodHS256, model.NewUserClaims(payload.ID, exp))
	tokenJwt, err := tokenJwtSementara.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		return "", err
	}
	return tokenJwt, nil
}
