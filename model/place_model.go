package model

type CreatePlaceRequest struct {
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
}

type GetPlaceByIDRequest struct {
	ID uint `uri:"id" binding:"required"`
}
