package Utils

import (
	"SSO_BE_API/Config"
	"SSO_BE_API/Model/Entity"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte(Config.JWT_KEY)
var refreshTokenKey = []byte(Config.JWT_REFRESH_TOKEN_KEY)

type MyClaims struct {
	UserCredential int
	jwt.RegisteredClaims
}

func generate(user Entity.User, duration time.Time, key []byte) string {
	claims := MyClaims{
		UserCredential: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "awesomeProject",
			Subject:   "awesomeProject",
			ExpiresAt: jwt.NewNumericDate(duration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(key)
	return ss
}

func GenerateRefreshToken(user Entity.User) string {
	return generate(user, time.Now().Add(time.Hour*24*30), refreshTokenKey)
}

func GenerateJwtToken(User Entity.User) string {
	return generate(User, time.Now().Add(time.Minute*30), jwtKey)
}

func Claims(token string) (*MyClaims, error) {
	claims := &MyClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, err
	}

	return claims, nil
}
