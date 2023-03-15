package handler

import (
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
	claimsTemp, _ := c.Get("user")
	var claims model.UserClaims
	if claimsTemp != nil {
		claims = claimsTemp.(model.UserClaims)
	}

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

	var identities model.Identities
	if claimsTemp == nil {
		identities = model.Identities{
			ID:    0,
			Nama:  "",
			Email: "",
		}
	} else {
		customer, err := h.Repository.GetCustomerByID(claims.ID)
		if err != nil {
			response.FailOrError(c, http.StatusNotFound, "customer not found", err)
			return
		}
		identities = model.Identities{
			ID:    customer.ID,
			Nama:  customer.Nama,
			Email: customer.Email,
		}
	}

	response.Success(c, http.StatusOK, "Spaces found", gin.H{
		"iden":   identities,
		"spaces": spaces,
	}, &spaceParam)
}

func (h *spaceHandler) GetSpaceByParam(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	var claims model.UserClaims
	if claimsTemp != nil {
		claims = claimsTemp.(model.UserClaims)
	}

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

	var identities model.Identities
	if claimsTemp == nil {
		identities = model.Identities{
			ID:    0,
			Nama:  "",
			Email: "",
		}
	} else {
		customer, err := h.Repository.GetCustomerByID(claims.ID)
		if err != nil {
			response.FailOrError(c, http.StatusNotFound, "customer not found", err)
			return
		}
		identities = model.Identities{
			ID:    customer.ID,
			Nama:  customer.Nama,
			Email: customer.Email,
		}
	}

	response.Success(c, http.StatusOK, "Spaces found", gin.H{
		"iden":   identities,
		"spaces": spaces,
	}, &spaceParam)
}

func (h *spaceHandler) GetSpaceByID(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	var claims model.UserClaims
	if claimsTemp != nil {
		claims = claimsTemp.(model.UserClaims)
	}

	request := model.GetByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid request", err)
		return
	}
	space, err := h.Repository.GetSpaceByID(request.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "space not found", err)
		return
	}

	var identities model.Identities
	if claimsTemp == nil {
		identities = model.Identities{
			ID:    0,
			Nama:  "",
			Email: "",
		}
	} else {
		customer, err := h.Repository.GetCustomerByID(claims.ID)
		if err != nil {
			response.FailOrError(c, http.StatusNotFound, "customer not found", err)
			return
		}
		identities = model.Identities{
			ID:    customer.ID,
			Nama:  customer.Nama,
			Email: customer.Email,
		}
	}

	response.Success(c, http.StatusOK, "space found", gin.H{
		"iden":   identities,
		"spaces": space,
	}, nil)
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
