package handlers

import (
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/models"
	"gateway/pkg/utils"
	"log"
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
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.HealthCheck(ctx, body)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *MonitizationHandler) CreateWallet(ctx *gin.Context) {
	body := models.CreateWalletRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.CreateWallet(ctx, body)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *MonitizationHandler) UpdateWallet(ctx *gin.Context) {
	body := models.UpdateWalletRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.UpdateWallet(ctx, body)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *MonitizationHandler) ParticipationReward(ctx *gin.Context) {
	body := models.ParticipationRewardRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ParticipationReward(ctx, body)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *MonitizationHandler) UserRewardHistory(ctx *gin.Context) {
	body := models.UserRewardHistoryRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.UserRewardHistory(ctx, body)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *MonitizationHandler) GetWallet(ctx *gin.Context) {

	res, err := h.Client.GetWallet(ctx)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *MonitizationHandler) ExclusiveContent(ctx *gin.Context) {
	body := models.ExclusiveContentRequest{}
	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ExclusiveContent(ctx, body.VideoId)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)

}
