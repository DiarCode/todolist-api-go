package models

import "github.com/dgrijalva/jwt-go"

type Token struct {
	Token string `json:"token"`
}

type Claims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`

	jwt.StandardClaims
}
