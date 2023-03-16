package handler

import (
	"errors"
	"intern_BCC/model"
	"intern_BCC/repository"
	"intern_BCC/sdk/response"
	"net/http"
	"strconv"

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

	order := model.Order{
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

func (h *orderHandler) CreateReview(c *gin.Context) {
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
	order, _ := h.Repository.GetOrderByID(id.ID)
	customer, err := h.Repository.GetCustomerByID(order.CustomerID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "customer not found", err)
		return
	}

	review := model.Review{
		CustomerID: claims.ID,
		SpaceID:    order.SpaceID,
		OrderID:    id.ID,
		Nama:       customer.Nama,
		Ulasan:     request.Ulasan,
		Rating:     int(rate),
	}
	err = h.Repository.CreateReview(&review)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "create review failed", err)
		return
	}

	reviews, count, err := h.Repository.GetReviewsBySpaceID(order.SpaceID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "reviews not found", err)
		return
	}
	var newRating float64
	for _, rating := range reviews {
		newRating += float64(rating.Rating)
	}
	newRating /= float64(count)
	err = h.Repository.UpdateRating(order.SpaceID, newRating)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "update rating failed", err)
		return
	}

	response.Success(c, http.StatusCreated, "review creation succeeded", nil, nil)
}
