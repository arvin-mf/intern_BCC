package handler

import (
	"intern_BCC/model"
	"intern_BCC/repository"
	"intern_BCC/sdk/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type optionHandler struct {
	Repository repository.OptionRepository
}

func NewOptionHandler(repo *repository.OptionRepository) optionHandler {
	return optionHandler{*repo}
}

func (h *optionHandler) CreateOption(c *gin.Context) {
	request := model.CreateOptionRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "create option failed", err)
		return
	}
	option := model.Option{
		SpaceID: request.SpaceID,
		Rentang: request.Rentang,
	}
	err := h.Repository.CreateOption(&option)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "create option failed", err)
		return
	}
	response.Success(c, http.StatusCreated, "option created", nil, nil)
}
