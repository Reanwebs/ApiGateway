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

	// setup JWT
	ok := middleware.JwtCookieSetup(ctx, "admin-auth", res.Uid, res.Error)
	if !ok {
		res := errors.New("Generate JWT failure")
		ctx.JSON(http.StatusInternalServerError, res)
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

func (h *AdminHandler) ManageInterest(ctx *gin.Context) {
	body := models.ManageInterestRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ManageInterest(context.Background(), body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "action failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *AdminHandler) ManageCommunity(ctx *gin.Context) {
	body := models.ManageCommunityRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ManageCommunity(context.Background(), body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "action failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *AdminHandler) LogoutAdmin(ctx *gin.Context) {
	ctx.SetCookie("admin-auth", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Log out successful",
	})
}

func (h *AdminHandler) GetAllCommunity(ctx *gin.Context) {

	res, err := h.Client.GetAllCommunity(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to fetch all community",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
