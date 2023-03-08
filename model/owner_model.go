package model

type CreateOwnerRequest struct {
	Email    string `json:"email" binding:"required"`
	Whatsapp string `json:"whatsapp" binding:"required"`
}

type GetOwnerByIDRequest struct {
	ID uint `uri:"id" binding:"required"`
}
