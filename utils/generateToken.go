package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string
	Role  string
	jwt.StandardClaims
}

func GenerateToken(email, role string, cfg string) (string, error) {
	expireTime := time.Now().Add(time.Minute * 20).Unix()
	fmt.Println(email, role)
	claims := &Claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Subject:   email,
			IssuedAt:  time.Now().Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString([]byte(cfg))
	if err != nil {
		log.Printf("unable to generate jwt token for user %v, err: %v", email, err.Error())
		return "", err
	}

	return signedToken, nil
}
