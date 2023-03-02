package model

type CreateCustomerRequest struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetCustomerByIDRequest struct {
	ID uint `uri:"id" binding:"required"`
}
