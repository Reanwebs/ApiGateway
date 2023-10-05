package handlers

import (
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MonitizationHandler struct {
	Client interfaces.MonitaizationClient
}

func NewMonitizationHandler(client interfaces.MonitaizationClient) MonitizationHandler {
	return MonitizationHandler{
		Client: client,
	}
}

func (h *MonitizationHandler) MonitaizationHealthCheck(ctx *gin.Context) {
	body := models.Request{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.HealthCheck(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "conference service down",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *MonitizationHandler) CreateWallet(ctx *gin.Context) {
	body := models.CreateWalletRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.CreateWallet(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to create wallet",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *MonitizationHandler) UpdateWallet(ctx *gin.Context) {
	body := models.UpdateWalletRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.UpdateWallet(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to update wallet",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *MonitizationHandler) ParticipationReward(ctx *gin.Context) {
	body := models.ParticipationRewardRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ParticipationReward(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to update wallet",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *MonitizationHandler) UserRewardHistory(ctx *gin.Context) {
	body := models.UserRewardHistoryRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.UserRewardHistory(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to fetch wallet",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *MonitizationHandler) GetWallet(ctx *gin.Context) {

	res, err := h.Client.GetWallet(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to fetch wallet",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *MonitizationHandler) ExclusiveContent(ctx *gin.Context) {
	var videoID string
	if err := ctx.BindJSON(&videoID); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ExclusiveContent(ctx, videoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to fetch wallet",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)

}
