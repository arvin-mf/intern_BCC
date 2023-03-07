package handler

import (
	"errors"
	"intern_BCC/model"
	"intern_BCC/repository"
	"intern_BCC/sdk/crypto"
	sdk_jwt "intern_BCC/sdk/jwt"
	"intern_BCC/sdk/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	Repository repository.CustomerRepository
}

func NewCustomerHandler(repo *repository.CustomerRepository) customerHandler {
	return customerHandler{*repo}
}

func (h *customerHandler) CreateCustomer(c *gin.Context) {
	var customer model.CreateCustomerRequest
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}
	_, err = h.Repository.FindByEmail(customer.Email)
	if err == nil {
		msg := "Email sudah dipakai"
		response.FailOrError(c, http.StatusBadRequest, "customer creation failed", errors.New(msg))
		return
	}
	if len(customer.Password) < 8 {
		msg := "Password minimal 8 karakter"
		response.FailOrError(c, http.StatusBadRequest, "customer creation failed", errors.New(msg))
		return
	}
	if customer.Password != customer.Konfirmpw {
		msg := "Konfirmasi password gagal"
		response.FailOrError(c, http.StatusBadRequest, "customer creation failed", errors.New(msg))
		return
	}
	result, err := h.Repository.CreateCustomer(customer)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "customer creation failed", err)
		return
	}
	response.Success(c, http.StatusCreated, "customer creation success", result, nil)
}

func (h *customerHandler) Login(c *gin.Context) {
	var request model.Login
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}
	customer, err := h.Repository.FindByEmail(request.Email)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "email not found", err)
		return
	}
	err = crypto.ValidateHash(request.Password, customer.Password)
	if err != nil {
		msg := "Password salah"
		response.FailOrError(c, http.StatusBadRequest, "wrong password", errors.New(msg))
		return
	}
	tokenJwt, err := sdk_jwt.GenerateToken(customer)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "create token failed", err)
		return
	}
	response.Success(c, http.StatusOK, "login success", gin.H{
		"token": tokenJwt,
	}, nil)
}

func (h *customerHandler) GetAllCustomer(c *gin.Context) {
	customers, err := h.Repository.GetAllCustomer()
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "customers not found", err)
		return
	}
	response.Success(c, http.StatusOK, "customers found", customers, nil)
}

func (h *customerHandler) GetCustomerByID(c *gin.Context) {
	request := model.GetCustomerByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "failed getting customer", err)
		return
	}
	customer, err := h.Repository.GetCustomerByID(request.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "customer not found", err)
		return
	}
	response.Success(c, http.StatusOK, "customer found", customer, nil)
}

func (h *customerHandler) DeleteCustomerByID(c *gin.Context) {
	ID := c.Param("id")
	parsedID, _ := strconv.ParseUint(ID, 10, 64)
	err := h.Repository.DeleteCustomerByID(uint(parsedID))
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "delete customer failed", err)
		return
	}
	response.Success(c, http.StatusOK, "deleting success", nil, nil)
}
