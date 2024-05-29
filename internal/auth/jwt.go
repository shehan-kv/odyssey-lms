package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte(os.Getenv("JWT_KEY"))

func NewJWTToken(id int64) (string, error) {

	// TODO: Update to send a refresh token
	// currently the JWT expires in 2 days (for development)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
		"iat": time.Now().Unix(),
	})

	s, err := t.SignedString(key)

	return s, err
}

func VerifyJWTToken(tokenEncoded string) (string, error) {

	tokenDecoded, err := jwt.ParseWithClaims(tokenEncoded, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return "", err
	}

	if !tokenDecoded.Valid {
		return "", errors.New("invalid JWT token")
	}

	subject, err := tokenDecoded.Claims.GetSubject()
	if err != nil {
		return "", err
	}

	return subject, nil
}
