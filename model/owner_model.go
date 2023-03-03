package model

type CreateOwnerRequest struct {
	Email    string `json:"email"`
	Whatsapp string `json:"whatsapp"`
}
