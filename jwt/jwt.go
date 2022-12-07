package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

// JWT is a struct that contains the JWT token
type JWT struct {
	Token string
}

// NewJWT creates a new JWT struct
func NewJWT(token string) *JWT {
	return &JWT{
		Token: token,
	}
}

// Validate validates the JWT token
func (j *JWT) Validate() (bool, error) {
	_, err := jwt.Parse(j.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
