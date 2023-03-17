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
	response.Success(c, http.StatusCreated, "owner creation success", result, nil)
}

func (h *ownerHandler) Login(c *gin.Context) {
	var request model.Login
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid body", err)
		return
	}
	owner, err := h.Repository.FindByEmail(request.Email)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "email not found", err)
		return
	}
	err = crypto.ValidateHash(request.Password, owner.Password)
	if err != nil {
		msg := "Password salah"
		response.FailOrError(c, http.StatusBadRequest, "wrong password", errors.New(msg))
		return
	}
	tokenJwt, err := sdk_jwt.GenerateOwnerToken(owner)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "create token failed", err)
		return
	}
	response.Success(c, http.StatusOK, "login success", gin.H{
		"token": tokenJwt,
	}, nil)
}

func (h *ownerHandler) GetOwnerSpaces(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)

	spaces, err := h.Repository.GetOwnerSpaces(claims.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "spaces not found", err)
		return
	}
	var spaceIDs []uint
	for _, space := range spaces {
		spaceIDs = append(spaceIDs, space.ID)
	}
	reviews, err := h.Repository.GetOwnerReviews(spaceIDs)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "reviews not found", err)
		return
	}

	response.Success(c, http.StatusOK, "spaces found", gin.H{
		"spaces":  spaces,
		"reviews": reviews,
	}, nil)
}

func (h *ownerHandler) GetOwnerSpaceByCat(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)

	request := model.GetByCatRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	space, err := h.Repository.GetOwnerSpaceByCat(claims.ID, request.Kategori)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "space not found", err)
		return
	}

	reviews, err := h.Repository.GetReviewsBySpaceID(space.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "reviews not found", err)
		return
	}

	response.Success(c, http.StatusOK, "space found", gin.H{
		"space":   space,
		"reviews": reviews,
	}, nil)
}

func (h *ownerHandler) UpdateDescription(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)

	request := model.DescriptionRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid body", err)
		return
	}

	err = h.Repository.UpdateDescription(claims.ID, request.Deskripsi)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "update description failed", err)
		return
	}
	response.Success(c, http.StatusOK, "description updated", request, nil)
}

func (h *ownerHandler) UpdateCapacity(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)

	request := model.CapacityRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid body", err)
		return
	}

	err = h.Repository.UpdateCapacity(claims.ID, request.Kapasitas)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "update capacity failed", err)
		return
	}
	response.Success(c, http.StatusOK, "capacity updated", request, nil)
}

func (h *ownerHandler) UpdatePrice(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)

	request := model.PriceRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid body", err)
		return
	}
	cat := model.GetByCatRequest{}
	err = c.ShouldBindUri(&cat)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid request", err)
		return
	}
	space, err := h.Repository.GetOwnerSpaceByCat(claims.ID, cat.Kategori)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "space not found", err)
		return
	}
	err = h.Repository.UpdatePrice(space.ID, request.Harga)

	response.Success(c, http.StatusOK, "price updated", request, nil)
}

func (h *ownerHandler) AddFacilities(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)

	request := model.GetByCatRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	var facils model.AddFacilRequest
	err := c.ShouldBindJSON(&facils)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid body", err)
		return
	}

	space, err := h.Repository.GetOwnerSpaceByCat(claims.ID, request.Kategori)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "space not found", err)
		return
	}

	err = h.Repository.AddFacilities(space.ID, facils.Fasil)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "update description failed", err)
		return
	}
	response.Success(c, http.StatusOK, "add facilities succeeded", facils, nil)
}

func (h *ownerHandler) AddGeneralFacility(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)

	var facils model.AddFacilRequest
	err := c.ShouldBindJSON(&facils)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid body", err)
		return
	}

	err = h.Repository.AddGeneralFacility(claims.ID, facils.Fasil)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "update description failed", err)
		return
	}
	response.Success(c, http.StatusOK, "add facilities succeeded", facils, nil)
}

func (h *ownerHandler) SwitchAvailability(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)

	request := model.GetByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "failed getting owner", err)
		return
	}

	date, err := h.Repository.GetDateByID(request.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "date not found", err)
		return
	}
	option, err := h.Repository.GetOptionByID(date.OptionID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "option not found", err)
		return
	}
	space, err := h.Repository.GetSpaceByID(option.SpaceID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "space not found", err)
		return
	}
	owner, err := h.Repository.GetOwnerByID(space.OwnerID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "owner not found", err)
		return
	}
	if claims.ID != owner.ID {
		response.FailOrError(c, http.StatusForbidden, "access denied", err)
		return
	}

	statusNow, err := h.Repository.SwitchAvailability(date.ID)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "switch failed", err)
		return
	}
	response.Success(c, http.StatusOK, "switch availability succeeded", statusNow, nil)
}

func (h *ownerHandler) AddPicture(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)
	if claims.Role != "owner" {
		msg := "access denied"
		response.FailOrError(c, http.StatusForbidden, msg, errors.New(msg))
		return
	}
	request := model.GetByCatRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid request", err)
		return
	}
	space, err := h.Repository.GetOwnerSpaceByCat(claims.ID, request.Kategori)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "space not found", err)
		return
	}
	if space.OwnerID != claims.ID {
		response.FailOrError(c, http.StatusForbidden, "access denied", err)
		return
	}
	link, err := h.Repository.Upload(c)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "file not accepted", err)
		return
	}
	err = h.Repository.AddPicture(space.ID, link)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "upload file failed", err)
		return
	}
	response.Success(c, http.StatusOK, "file uploaded", link, nil)
}

func (h *ownerHandler) AddGalleryPicture(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)

	link, err := h.Repository.Upload(c)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "file not accepted", err)
		return
	}
	galleryPic := model.Picture{
		Link:    link,
		OwnerID: claims.ID,
	}
	err = h.Repository.AddGalleryPicture(&galleryPic)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "add gallery failed", err)
		return
	}
	response.Success(c, http.StatusCreated, "gallery added", galleryPic, nil)
}

func (h *ownerHandler) GetAllPictures(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)

	links, err := h.Repository.GetAllPictures(claims.ID)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "get pictures failed", err)
		return
	}
	response.Success(c, http.StatusOK, "records found", links, nil)
}

func (h *ownerHandler) GetAllOwner(c *gin.Context) {
	owners, err := h.Repository.GetAllOwner()
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "owners not found", err)
		return
	}
	response.Success(c, http.StatusOK, "owners found", owners, nil)
}

func (h *ownerHandler) GetOwnerByID(c *gin.Context) {
	request := model.GetByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "failed getting owner", err)
		return
	}
	owner, err := h.Repository.GetOwnerByID(request.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "owner not found", err)
		return
	}
	response.Success(c, http.StatusOK, "owner found", owner, nil)
}

func (h *ownerHandler) DeleteOwnerByID(c *gin.Context) {
	ID := c.Param("id")
	parsedID, _ := strconv.ParseUint(ID, 10, 64)
	err := h.Repository.DeleteOwnerByID(uint(parsedID))
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "delete owner failed", err)
		return
	}
	response.Success(c, http.StatusOK, "deleting success", nil, nil)
}
