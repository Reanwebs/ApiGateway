package handlers

import (
	"context"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type conferenceHandler struct {
	Client interfaces.ConferenceClient
}

func NewConferenceHandler(client interfaces.ConferenceClient) conferenceHandler {
	return conferenceHandler{
		Client: client,
	}
}

func (h *conferenceHandler) StartConference(ctx *gin.Context) {
	body := models.StartConferenceRequest{}
	// traceId := utils.GenerateTraceID()
	// ctx.Set("traceId", traceId)

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.StartConference(context.Background(), body)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.ConferenceID), &res)
}

func (h *conferenceHandler) JoinConference(ctx *gin.Context) {
	body := models.JoinConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.JoinConference(context.Background(), body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)

}

func (h *conferenceHandler) AcceptJoining(ctx *gin.Context) {
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

func (h *conferenceHandler) DeclineJoining(ctx *gin.Context) {
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

func (h *conferenceHandler) LeaveConference(ctx *gin.Context) {
	body := models.LeaveConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.LeaveConference(context.Background(), body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

func (h *conferenceHandler) RemoveParticipant(ctx *gin.Context) {
	body := models.RemoveParticipantRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.RemoveParticipant(context.Background(), body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *conferenceHandler) ToggleCamera(ctx *gin.Context) {
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

func (h *conferenceHandler) ToggleMic(ctx *gin.Context) {
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

func (h *conferenceHandler) ToggleParticipantCamera(ctx *gin.Context) {
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

func (h *conferenceHandler) ToggleParticipantMic(ctx *gin.Context) {
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

func (h *conferenceHandler) EndConference(ctx *gin.Context) {
	body := models.EndConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.EndConference(context.Background(), body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)

}
