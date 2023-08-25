package handlers

import (
	"context"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConferenceHandler struct {
	Client interfaces.ConferenceClient
}

func NewConferenceHandler(client interfaces.ConferenceClient) ConferenceHandler {
	return ConferenceHandler{
		Client: client,
	}
}

func (h *ConferenceHandler) HealthCheck(ctx *gin.Context) {
	body := models.HealthCheck{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.HealthCheck(ctx, body.Data)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "conference service down",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) ScheduleConference(ctx *gin.Context) {
	body := models.ScheduleConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ScheduleConference(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in scheduleConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

func (h *ConferenceHandler) StartPrivateConference(ctx *gin.Context) {
	body := models.StartPrivateConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.StartPrivateConference(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in StartPrivateConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) StartPublicConference(ctx *gin.Context) {
	body := models.StartPublicConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.StartPublicConference(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in StartPublicConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) JoinPrivateConference(ctx *gin.Context) {
	body := models.JoinPrivateConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.JoinPrivateConference(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in JoinPrivateConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) JoinPublicConference(ctx *gin.Context) {
	body := models.JoinPublicConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.JoinPublicConference(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in JoinPublicConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) JoinGroupConfernce(ctx *gin.Context) {
	body := models.JoinGroupConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.JoinGroupConfernce(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in JoinGroupConfernce",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) LeavePrivateConference(ctx *gin.Context) {
	body := models.LeavePrivateConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.LeavePrivateConference(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in LeavePrivateConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) EndPublicConference(ctx *gin.Context) {
	body := models.EndPublicConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.EndPublicConference(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in EndPublicConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) EndGroupConference(ctx *gin.Context) {
	body := models.EndGroupConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.EndGroupConference(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in EndGroupConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) EndPrivateConference(ctx *gin.Context) {
	body := models.EndPrivateConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.EndPrivateConference(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in EndPrivateConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) RemovePublicParticipant(ctx *gin.Context) {
	body := models.RemovePublicParticipantRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.RemovePublicParticipant(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in RemovePublicParticipant",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) RemoveGroupParticipant(ctx *gin.Context) {
	body := models.RemoveGroupParticipantRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.RemoveGroupParticipant(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in RemoveGroupParticipant",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) RemovePrivateParticipant(ctx *gin.Context) {
	body := models.RemovePrivateParticipantRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.RemovePrivateParticipant(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in RemovePrivateParticipant",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) LeavePublicConference(ctx *gin.Context) {
	body := models.LeavePublicConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.LeavePublicConference(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in LeavePublicConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
func (h *ConferenceHandler) LeaveGroupConference(ctx *gin.Context) {
	body := models.LeaveGroupConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.LeaveGroupConference(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in LeaveGroupConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) StartGroupConference(ctx *gin.Context) {
	body := models.StartGroupConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.StartGroupConference(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in StartGroupConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) AcceptJoining(ctx *gin.Context) {
	body := models.AcceptJoiningRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.AcceptJoining(context.Background(), body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

func (h *ConferenceHandler) DeclineJoining(ctx *gin.Context) {
	body := models.DeclineJoiningRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.DeclineJoining(context.Background(), body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

func (h *ConferenceHandler) ToggleCamera(ctx *gin.Context) {
	body := models.ToggleCameraRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ToggleCamera(context.Background(), body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) ToggleMic(ctx *gin.Context) {
	body := models.ToggleMicRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ToggleMic(context.Background(), body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) ToggleParticipantCamera(ctx *gin.Context) {
	body := models.ToggleParticipantCameraRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ToggleParticipantCamera(context.Background(), body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) ToggleParticipantMic(ctx *gin.Context) {
	body := models.ToggleParticipantMicRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ToggleParticipantMic(context.Background(), body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
