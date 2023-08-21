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
