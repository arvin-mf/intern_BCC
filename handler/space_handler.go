package handler

import (
	"errors"
	"intern_BCC/entity"
	"intern_BCC/model"
	"intern_BCC/repository"
	"intern_BCC/sdk/response"
	"net/http"

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
		response.FailOrError(c, http.StatusUnprocessableEntity, "create space failed", err)
		return
	}
	space := entity.Space{
		Nama:     request.Nama,
		Kategori: entity.Category[request.CategoryID-1],
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
	response.Success(c, http.StatusOK, "Spaces found", spaces, &spaceParam)
}

func (h *spaceHandler) GetSpaceByID(c *gin.Context) {
	request := model.GetByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "failed getting owner", err)
		return
	}
	owner, err := h.Repository.GetSpaceByID(request.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "owner not found", err)
		return
	}
	response.Success(c, http.StatusOK, "owner found", owner, nil)
}

func (h *spaceHandler) AddPicture(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)
	if claims.Role != "owner" {
		msg := "access denied"
		response.FailOrError(c, http.StatusForbidden, msg, errors.New(msg))
		return
	}
	request := model.GetByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "failed getting owner", err)
		return
	}
	space, err := h.Repository.GetSpaceByID(request.ID)
	if space.OwnerID != claims.ID {
		response.FailOrError(c, http.StatusForbidden, "access denied", err)
		return
	}
	link, err := h.Repository.Upload(c)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "file not accepted", err)
		return
	}
	err = h.Repository.AddPicture(request.ID, link)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "upload file failed", err)
		return
	}
	response.Success(c, http.StatusOK, "file uploaded", nil, nil)
}
