package model

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
