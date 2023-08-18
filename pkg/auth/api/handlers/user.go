package handlers

import (
	"context"
	"gateway/pkg/auth/common/client/interfaces"
	"gateway/pkg/auth/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Client interfaces.AuthClient
}

func NewAuthHandler(client interfaces.AuthClient) UserHandler {
	return UserHandler{
		Client: client,
	}
}

func (h *UserHandler) UserSignup(ctx *gin.Context) {
	body := models.RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.UserSignup(context.Background(), body)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}

func (h *UserHandler) OtpValidation(ctx *gin.Context) {
	body := models.OtpValidation{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := h.Client.OtpValidation(ctx, body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)

}

func (h *UserHandler) UserLogin(ctx *gin.Context) {
	body := models.LoginRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.UserLogin(ctx, body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func (h *UserHandler) ValidName(ctx *gin.Context) {
	body := models.ValidName{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ValidName(ctx, body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)

}
