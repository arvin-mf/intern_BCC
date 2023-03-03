package handler

import (
	"intern_BCC/model"
	"intern_BCC/repository"
	"intern_BCC/sdk/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ownerHandler struct {
	Repository repository.OwnerRepository
}

func NewOwnerHandler(repo *repository.OwnerRepository) ownerHandler {
	return ownerHandler{*repo}
}

func (h *ownerHandler) CreateOwner(c *gin.Context) {
	var owner model.CreateOwnerRequest
	err := c.ShouldBindJSON(&owner)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}
	result, err := h.Repository.CreateOwner(owner)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "owner creation failed", err)
		return
	}
	response.Success(c, http.StatusCreated, "owner creation success", result)
}
