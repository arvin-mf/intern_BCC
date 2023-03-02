package handler

import (
	"intern_BCC/model"
	"intern_BCC/repository"
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
	result, err := h.Repository.CreateCustomer(customer)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Customer creation failed", err)
		return
	}
	response.Success(c, http.StatusCreated, "Customer creation success", result)
}

func (h *customerHandler) GetAllCustomer(c *gin.Context) {
	customers, err := h.Repository.GetAllCustomer()
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Customers not found", err)
		return
	}
	response.Success(c, http.StatusOK, "Customers found", customers)
}

func (h *customerHandler) GetCustomerByID(c *gin.Context) {
	request := model.GetCustomerByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Failed getting customer", err)
		return
	}
	customer, err := h.Repository.GetCustomerByID(request.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Customer not found", err)
		return
	}
	response.Success(c, http.StatusOK, "Customer found", customer)
}

func (h *customerHandler) DeleteCustomerByID(c *gin.Context) {
	ID := c.Param("id")
	parsedID, _ := strconv.ParseUint(ID, 10, 64)
	err := h.Repository.DeleteCustomerByID(uint(parsedID))
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Delete customer failed", err)
		return
	}
	response.Success(c, http.StatusOK, "Deleting success", nil)
}
