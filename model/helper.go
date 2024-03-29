package model

import (
	"math"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

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

type Identities struct {
	ID    uint   `json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
}

type GetByIDRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type PaginParam struct {
	Limit           int `json:"limit" form:"limit"`
	Page            int `json:"page" form:"page"`
	Offset          int `json:"offset"`
	TotalElements   int `json:"total_elements"`
	CurrentElements int `json:"current_elements"`
	TotalPages      int `json:"total_pages"`
	CurrentPage     int `json:"current_page"`
}

func (pp *PaginParam) FormatPagin() {
	if pp.Limit == 0 {
		pp.Limit = 9
	}
	if pp.Page == 0 {
		pp.Page = 1
	}
	pp.Offset = (pp.Page - 1) * pp.Limit
}

func (pp *PaginParam) ProcessPagin(totalElements int) {
	pp.TotalElements = totalElements
	pp.TotalPages = int(math.Ceil(float64(pp.TotalElements) / float64(pp.Limit)))
	pp.CurrentPage = pp.Page
	if pp.TotalPages > pp.TotalElements/pp.Limit {
		if pp.Page < pp.TotalPages {
			pp.CurrentElements = pp.Limit
		} else {
			pp.CurrentElements = pp.TotalElements - (pp.TotalPages-1)*pp.Limit
		}
	} else {
		pp.CurrentElements = pp.Limit
	}
}
