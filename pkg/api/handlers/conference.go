package handlers

import (
	"context"
	"fmt"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/models"
	"net/http"
	"time"

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

// CONFERENCE HEALTH CHECK
//
//	@Summary		API FOR CONFERENCE SERVICE HEALTH CHECK
//	@ID				CONFERENCE-HEALTH-CHECK
//	@Description	CONFERENCE SERVICE HEALTH CHECK
//	@Tags			CONFERENCE
//	@Accept			json
//	@Produce		json
//
//	@Param			HEALTH-CHECK	body		models.HealthCheck	false	"Request body for health check"
//
//	@Success		200				{object}	conference.Response
//	@Failure		400				{object}	conference.Response
//	@Failure		502				{object}	conference.Response
//	@Router			/api/conference/health-check [post]
func (h *ConferenceHandler) HealthCheck(ctx *gin.Context) {
	body := models.HealthCheck{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   5 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.HealthCheck(ctx, body.Data, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "conference service down",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) SchedulePrivateConference(ctx *gin.Context) {
	body := models.SchedulePrivateConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println("\n\n", body.Time, "\n\n", body)
	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   5 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.SchedulePrivateConference(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "can't schedule conference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) ScheduleGroupConference(ctx *gin.Context) {
	body := models.ScheduleGroupConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   5 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.ScheduleGroupConference(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "can't schedule conference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
func (h *ConferenceHandler) SchedulePublicConference(ctx *gin.Context) {
	body := models.SchedulePublicConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   5 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.SchedulePublicConference(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "can't schedule conference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// CONFERENCE START PRIVATE CONFERENCE
//
//	@Summary		API FOR START PRIVATE CONFERENCE
//	@ID				START-PRIVATE-CONFERENCE
//	@Description	START PRIVATE CONFERENCE
//	@Tags			CONFERENCE
//	@Accept			json
//	@Produce		json
//	@Param			START-PRIVATE-CONFERENCE	body		models.StartPrivateConferenceRequest false	"Request body for start private conference"
//	@Success		200							{object}	conference.StartPrivateConferenceResponse
//	@Failure		400							{object}	conference.StartPrivateConferenceResponse
//	@Failure		502							{object}	conference.StartPrivateConferenceResponse
//	@Router			/api/conference/start-private-conference [post]
func (h *ConferenceHandler) StartPrivateConference(ctx *gin.Context) {
	body := models.StartPrivateConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.StartPrivateConference(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in StartPrivateConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// CONFERENCE START PUBLIC CONFERENCE
//
//	@Summary		API FOR START PUBLIC CONFERENCE
//	@ID				START-PUBLIC-CONFERENCE
//	@Description	START PUBLIC CONFERENCE
//	@Tags			CONFERENCE
//	@Accept			json
//	@Produce		json
//	@Param			START-PUBLIC-CONFERENCE	body		models.StartPublicConferenceRequest false	"Request body for start public conference"
//	@Success		200						{object}	conference.StartPublicConferenceResponse
//	@Failure		400						{object}	conference.StartPublicConferenceResponse
//	@Failure		502						{object}	conference.StartPublicConferenceResponse
//	@Router			/api/conference/start-public-conference [post]
func (h *ConferenceHandler) StartPublicConference(ctx *gin.Context) {
	body := models.StartPublicConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.StartPublicConference(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in StartPublicConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// CONFERENCE JOIN PRIVATE CONFERENCE
//
//	@Summary		API FOR JOIN PRIVATE CONFERENCE
//	@ID				JOIN-PRIVATE-CONFERENCE
//	@Description	JOIN PRIVATE CONFERENCE
//	@Tags			CONFERENCE
//	@Accept			json
//	@Produce		json
//	@Param			JOIN-PRIVATE-CONFERENCE	body		models.JoinPrivateConferenceRequest false	"Request body for join private conference"
//	@Success		200						{object}	conference.JoinPrivateConferenceResponse
//	@Failure		400						{object}	conference.JoinPrivateConferenceResponse
//	@Failure		502						{object}	conference.JoinPrivateConferenceResponse
//	@Router			/api/conference/join-private-conference [patch]
func (h *ConferenceHandler) JoinPrivateConference(ctx *gin.Context) {
	body := models.JoinPrivateConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.JoinPrivateConference(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in JoinPrivateConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// CONFERENCE JOIN PUBLIC CONFERENCE
//
//	@Summary		API FOR JOIN PUBLIC CONFERENCE
//	@ID				JOIN-PUBLIC-CONFERENCE
//	@Description	JOIN PUBLIC CONFERENCE
//	@Tags			CONFERENCE
//	@Accept			json
//	@Produce		json
//	@Param			JOIN-PUBLIC-CONFERENCE	body		models.JoinPublicConferenceRequest false	"Request body for sjoin public conference"
//	@Success		200						{object}	conference.JoinPublicConferenceResponse
//	@Failure		400						{object}	conference.JoinPublicConferenceResponse
//	@Failure		502						{object}	conference.JoinPublicConferenceResponse
//	@Router			/api/conference/join-public-conference [patch]
func (h *ConferenceHandler) JoinPublicConference(ctx *gin.Context) {
	body := models.JoinPublicConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.JoinPublicConference(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error in JoinPublicConference",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// CONFERENCE JOIN GROUP CONFERENCE
//
//	@Summary		API FOR JOIN GROUP CONFERENCE
//	@ID				JOIN-GROUP-CONFERENCE
//	@Description	JOIN GROUP CONFERENCE
//	@Tags			CONFERENCE
//	@Accept			json
//	@Produce		json
//	@Param			JOIN-GROUP-CONFERENCE	body		models.JoinGroupConferenceRequest false	"Request body for join group conference"
//	@Success		200						{object}	conference.JoinGroupConferenceResponse
//	@Failure		400						{object}	conference.JoinGroupConferenceResponse
//	@Failure		502						{object}	conference.JoinGroupConferenceResponse
//	@Router			/api/conference/join-group-confernce [patch]
func (h *ConferenceHandler) JoinGroupConfernce(ctx *gin.Context) {
	body := models.JoinGroupConferenceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.JoinGroupConfernce(ctx, body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.LeavePrivateConference(ctx, body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.EndPublicConference(ctx, body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.EndGroupConference(ctx, body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.EndPrivateConference(ctx, body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.RemovePublicParticipant(ctx, body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.RemoveGroupParticipant(ctx, body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.RemovePrivateParticipant(ctx, body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.LeavePublicConference(ctx, body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.LeaveGroupConference(ctx, body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.StartGroupConference(ctx, body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.AcceptJoining(context.Background(), body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.DeclineJoining(context.Background(), body, retryConfig)
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

func (h *ConferenceHandler) ScheduledConference(ctx *gin.Context) {
	res, err := h.Client.ScheduledConference(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to fetch",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
func (h *ConferenceHandler) CompletedSchedules(ctx *gin.Context) {
	res, err := h.Client.CompletedSchedules(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to fetch",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
func (h *ConferenceHandler) StartStream(ctx *gin.Context) {
	body := models.StartStreamRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.StartStream(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to start stream",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) JoinStream(ctx *gin.Context) {
	body := models.JoinStreamRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.JoinStream(context.Background(), body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to join stream",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) LeaveStream(ctx *gin.Context) {
	body := models.LeaveStreamRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.LeaveStream(context.Background(), body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to leave stream",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) StopStream(ctx *gin.Context) {
	body := models.StopStreamRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.StopStream(context.Background(), body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to stop stream",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) GetStream(ctx *gin.Context) {
	body := models.GetStreamRequest{}
	streamID := ctx.Query("streamID")

	if streamID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "streamID is required",
		})
		return
	}
	body.StreamID = streamID

	res, err := h.Client.GetStream(context.Background(), body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to get stream",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *ConferenceHandler) GetOngoingStreams(ctx *gin.Context) {
	sortString := ctx.DefaultQuery("sort", "")
	res, err := h.Client.GetOngoingStreams(context.Background(), sortString)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to get stream",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
