package models

import "github.com/dgrijalva/jwt-go"

type TokenResponse struct {
	ID    int    `json:"user_id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type Claims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`

	jwt.StandardClaims
}
