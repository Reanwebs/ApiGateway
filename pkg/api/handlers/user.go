package handlers

import (
	"context"
	"errors"
	"gateway/pkg/common/client/interfaces"
	middleware "gateway/pkg/common/midleware"
	"gateway/pkg/common/models"
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

func (h *UserHandler) OtpRequest(ctx *gin.Context) {
	body := models.OtpValidation{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := h.Client.OtpRequest(ctx, body)
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
		ctx.JSON(http.StatusBadGateway, err)
		return
	}

	// setup JWT
	ok := middleware.JwtCookieSetup(ctx, "user-auth", uint(res.Token))
	if !ok {
		res := errors.New("failed to login")
		ctx.JSON(http.StatusInternalServerError, res)
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "User name not available",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)

}
