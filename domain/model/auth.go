package model

import "github.com/golang-jwt/jwt/v4"

type MyClaims struct {
	jwt.RegisteredClaims
	Username string `json:"Username"`
}
