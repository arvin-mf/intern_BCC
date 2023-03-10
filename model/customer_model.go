package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type CreateCustomerRequest struct {
	Nama      string `json:"nama" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Konfirmpw string `json:"konfirmpw" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserClaims struct {
	ID   uint   `json:"id" binding:"required"`
	Role string `json:"role" binding:"required"`
	jwt.RegisteredClaims
}

func NewUserClaims(id uint, role string, exp time.Duration) UserClaims {
	return UserClaims{
		ID:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
}

type GetCustomerByIDRequest struct {
	ID uint `uri:"id" binding:"required"`
}
