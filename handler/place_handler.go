package handler

import (
	"intern_BCC/entity"
	"intern_BCC/model"
	"intern_BCC/repository"
	"intern_BCC/sdk/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type placeHandler struct {
	Repository repository.PlaceRepository
}

func NewPlaceHandler(repo *repository.PlaceRepository) placeHandler {
	return placeHandler{*repo}
}

func (h *placeHandler) CreatePlace(c *gin.Context) {
	request := model.CreatePlaceRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "Create place failed", err)
		return
	}
	place := entity.Place{
		Nama:    request.Nama,
		Alamat:  request.Alamat,
		OwnerID: request.OwnerID,
	}
	err := h.Repository.CreatePlace(&place)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Create place failed", err)
		return
	}
	response.Success(c, http.StatusCreated, "Place creation succeeded", request)
}

func (h *placeHandler) GetAllPlace(c *gin.Context) {
	places, err := h.Repository.GetAllPlace()
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Places not found", err)
		return
	}
	response.Success(c, http.StatusOK, "Places found", places)
}

func (h *placeHandler) GetPlaceByID(c *gin.Context) {
	request := model.GetPlaceByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Failed getting place", err)
		return
	}

	place, err := h.Repository.GetPlaceByID(request.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Place not found", err)
		return
	}

	response.Success(c, http.StatusOK, "Place found", place)
}

func (h *placeHandler) DeletePlaceByID(c *gin.Context) {
	ID := c.Param("id")
	parsedID, _ := strconv.ParseUint(ID, 10, 64)
	err := h.Repository.DeletePlaceByID(uint(parsedID))
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Place deleting failed", err)
		return
	}
	response.Success(c, http.StatusOK, "Successfully deleted place", nil)
}
