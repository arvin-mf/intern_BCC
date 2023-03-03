package model

type CreatePlaceRequest struct {
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	OwnerID uint   `json:"owner_id"`
}

type GetPlaceByIDRequest struct {
	ID uint `uri:"id" binding:"required"`
}
