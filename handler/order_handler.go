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

type orderHandler struct {
	Repository repository.OrderRepository
}

func NewOrderHandler(repo *repository.OrderRepository) orderHandler {
	return orderHandler{*repo}
}

func (h *orderHandler) CreateOrder(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)
	if claims.Role != "customer" {
		msg := "access denied"
		response.FailOrError(c, http.StatusForbidden, msg, errors.New(msg))
		return
	}

	request := model.GetByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "failed getting owner", err)
		return
	}

	order := entity.Order{
		CustomerID: claims.ID,
		SpaceID:    request.ID,
	}

	err := h.Repository.CreateOrder(&order)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "create order failed", err)
		return
	}
	response.Success(c, http.StatusCreated, "order creation succeeded", request, nil)
}

func (h *orderHandler) GetAllOrder(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)

	orders, err := h.Repository.GetAllOrder(claims.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "orders not found", err)
		return
	}
	response.Success(c, http.StatusOK, "orders found", orders, nil)
}

func (h *orderHandler) GetOrderByID(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)

	request := model.GetByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "failed getting order", err)
		return
	}
	order, err := h.Repository.GetOrderByID(request.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "order not found", err)
		return
	}
	if order.CustomerID != claims.ID {
		response.FailOrError(c, http.StatusForbidden, "access denied", err)
	}
	response.Success(c, http.StatusOK, "order found", order, nil)
}
