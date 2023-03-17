package handler

import (
	"intern_BCC/model"
	"intern_BCC/repository"
	"intern_BCC/sdk/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type spaceHandler struct {
	Repository repository.SpaceRepository
}

func NewSpaceHandler(repo *repository.SpaceRepository) spaceHandler {
	return spaceHandler{*repo}
}

func (h *spaceHandler) CreateSpace(c *gin.Context) {
	request := model.CreateSpaceRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "invalid request", err)
		return
	}
	space := model.Space{
		Nama:     request.Nama,
		Kategori: model.Category[request.CategoryID-1],
		Alamat:   request.Alamat,
		Harga:    request.Harga,
		Periode:  request.Periode,
		OwnerID:  request.OwnerID,
	}
	err := h.Repository.CreateSpace(&space)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "create space failed", err)
		return
	}
	response.Success(c, http.StatusCreated, "space creation succeeded", request, nil)
}

func (h *spaceHandler) GetAllSpace(c *gin.Context) {
	var spaceParam model.PaginParam
	if err := h.Repository.BindParam(c, &spaceParam); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid request body", err)
		return
	}
	spaceParam.FormatPagin()
	spaces, totalElements, err := h.Repository.GetAllSpace(&spaceParam)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Spaces not found", err)
		return
	}
	spaceParam.ProcessPagin(totalElements)

	location := model.UserLocation{
		Lat: -7.9537341,
		Lon: 112.609102,
	}
	err = h.Repository.UpdateDistance(spaces, location)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "update distance failed", err)
		return
	}

	response.Success(c, http.StatusOK, "Spaces found", spaces, &spaceParam)
}

func (h *spaceHandler) GetSpaceByParam(c *gin.Context) {
	var request model.CategoryRequest
	if err := h.Repository.BindParam(c, &request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid request body", err)
		return
	}
	var spaceParam model.PaginParam
	if err := h.Repository.BindParam(c, &spaceParam); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid request body", err)
		return
	}
	spaceParam.FormatPagin()
	spaces, totalElements, err := h.Repository.GetSpaceByParam(&spaceParam, &request)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Spaces not found", err)
		return
	}
	spaceParam.ProcessPagin(totalElements)

	location := model.UserLocation{
		Lat: -7.9537341,
		Lon: 112.609102,
	}
	err = h.Repository.UpdateDistance(spaces, location)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "update distance failed", err)
		return
	}

	response.Success(c, http.StatusOK, "Spaces found", spaces, &spaceParam)
}

func (h *spaceHandler) GetSpaceByID(c *gin.Context) {
	request := model.GetByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid request", err)
		return
	}
	space, options, err := h.Repository.GetSpaceByID(request.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "space not found", err)
		return
	}

	pics, err := h.Repository.GetPicturesByOwnerID(space.OwnerID)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "failed getting pictures", err)
		return
	}

	reviews, _, err := h.Repository.GetReviewsBySpaceID(space.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "reviews not found", err)
		return
	}

	response.Success(c, http.StatusOK, "space found", gin.H{
		"pictures": pics,
		"space":    space,
		"options":  options,
		"reviews":  reviews,
	}, nil)
}

func (h *spaceHandler) AlterCreateReview(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)

	request := model.CreateReviewRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "create review failed", err)
		return
	}
	rate, err := strconv.ParseUint(request.Rating, 10, 64)
	id := model.GetByIDRequest{}
	if err := c.ShouldBindUri(&id); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid request", err)
		return
	}
	customer, err := h.Repository.GetCustomerByID(claims.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "customer not found", err)
		return
	}
	review := model.Review{
		CustomerID: claims.ID,
		SpaceID:    id.ID,
		OrderID:    1,
		Nama:       customer.Nama,
		Ulasan:     request.Ulasan,
		Rating:     int(rate),
	}
	err = h.Repository.CreateReview(&review)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "review creation failed", err)
		return
	}

	reviews, count, err := h.Repository.GetReviewsBySpaceID(id.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "reviews not found", err)
		return
	}
	var newRating float64
	for _, rating := range reviews {
		newRating += float64(rating.Rating)
	}
	newRating /= float64(count)
	err = h.Repository.UpdateRating(id.ID, newRating)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "update rating failed", err)
		return
	}

	response.Success(c, http.StatusCreated, "review creation succeeded", review, nil)
}

func (h *spaceHandler) DeleteSpaceByID(c *gin.Context) {
	request := model.GetByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid request", err)
		return
	}
	err := h.Repository.DeleteSpaceByID(request.ID)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "delete space failed", err)
		return
	}
	response.Success(c, http.StatusOK, "delete space success", nil, nil)
}
