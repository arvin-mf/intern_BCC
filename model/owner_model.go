package model

type CreateOwnerRequest struct {
	Email    string `json:"email"`
	Whatsapp string `json:"whatsapp"`
}

type GetOwnerByIDRequest struct {
	ID uint `uri:"id" binding:"required"`
}
