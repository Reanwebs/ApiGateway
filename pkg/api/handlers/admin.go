package handlers

import (
	"context"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	Client interfaces.AdminClient
}

func NewAdminHandler(client interfaces.AdminClient) AdminHandler {
	return AdminHandler{
		Client: client,
	}
}

func (h *AdminHandler) AdminLogin(ctx *gin.Context) {
	body := models.AdminLoginRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.AdminLogin(context.Background(), body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "login failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func (h *AdminHandler) GetUsers(ctx *gin.Context) {

	res, err := h.Client.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "can't fetch users",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

func (h *AdminHandler) ManageUser(ctx *gin.Context) {
	body := models.ManageUserRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ManageUser(context.Background(), body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "action failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

func (h *AdminHandler) GetInterest(ctx *gin.Context) {

	res, err := h.Client.GetInterest(context.Background())
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "intrest fetching failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *AdminHandler) AddInterest(ctx *gin.Context) {
	body := models.AddInterestRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.AddInterest(context.Background(), body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "intrest adding failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &res)

}
