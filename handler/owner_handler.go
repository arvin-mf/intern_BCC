package handler

import (
	"errors"
	"intern_BCC/model"
	"intern_BCC/repository"
	"intern_BCC/sdk/crypto"
	"intern_BCC/sdk/response"
	"net/http"
	"strconv"

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

func (h *ownerHandler) Login(c *gin.Context) {
	var request model.Login
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}
	owner, err := h.Repository.FindByEmail(request.Email)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "email not found", err)
		return
	}
	err = crypto.ValidateHash(request.Password, owner.Password)
	if err != nil {
		msg := "Passsword salah"
		response.FailOrError(c, http.StatusBadRequest, "wrong password", errors.New(msg))
		return
	}
	/*
		PerTOKENan
	*/
	response.Success(c, http.StatusOK, "login success", nil)
}

func (h *ownerHandler) GetAllOwner(c *gin.Context) {
	owners, err := h.Repository.GetAllOwner()
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "owners not found", err)
		return
	}
	response.Success(c, http.StatusOK, "owners found", owners)
}

func (h *ownerHandler) GetOwnerByID(c *gin.Context) {
	request := model.GetOwnerByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "failed getting owner", err)
		return
	}
	owner, err := h.Repository.GetOwnerByID(request.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "owner not found", err)
		return
	}
	response.Success(c, http.StatusOK, "owner found", owner)
}

func (h *ownerHandler) DeleteOwnerByID(c *gin.Context) {
	ID := c.Param("id")
	parsedID, _ := strconv.ParseUint(ID, 10, 64)
	err := h.Repository.DeleteOwnerByID(uint(parsedID))
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "delete owner failed", err)
		return
	}
	response.Success(c, http.StatusOK, "deleting success", nil)
}
