package handlers

import (
	"errors"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/models"
	"gateway/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	Client interfaces.VideoClient
}

func NewVideoHandler(client interfaces.VideoClient) VideoHandler {
	return VideoHandler{
		Client: client,
	}
}

func (h *VideoHandler) UploadVideo(ctx *gin.Context) {

	body, err := utils.ParseUploadVideoForm(ctx)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}

	file, err := ctx.FormFile("video")
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}

	res, err := h.Client.UploadVideo(ctx, file, body)
	if err != nil {
		log.Println(err)
		errMsg := utils.ExtractErrorMessage(err.Error())
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Success": res,
	})
}

func (h *VideoHandler) FetchUserVideo(ctx *gin.Context) {
	request := models.FetchVideosRequest{}
	request.UserName = ctx.Query("userName")

	if request.UserName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user name required",
		})
		return
	}

	res, err := h.Client.FetchVideos(ctx, request)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)

}

func (h *VideoHandler) FetchUserArchivedVideo(ctx *gin.Context) {

	request := models.FetchVideosRequest{}
	request.UserName = ctx.Query("userName")

	if request.UserName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user name required",
		})
		return
	}

	res, err := h.Client.FetchArchivedVideos(ctx, request)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) FetchAllVideo(ctx *gin.Context) {

	res, err := h.Client.FetchAllVideos(ctx)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) ArchivVideo(ctx *gin.Context) {

	body := models.ArchivedVideos{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ArchiveVideo(ctx, body)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) GetVideoById(ctx *gin.Context) {

	videoId := ctx.Query("id")
	username := ctx.Query("userName")

	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user name required",
		})
		return
	}

	if videoId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "video id required",
		})
		return
	}

	data := models.GetVideoById{
		VideoId:  videoId,
		UserNAme: username,
	}

	res, err := h.Client.GetVideoById(ctx, data)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) ToggleStar(ctx *gin.Context) {
	body := models.ToggleStarRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ToggleStar(ctx, body)
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

func (h *VideoHandler) BlockVideo(ctx *gin.Context) {
	body := models.BlockVideoRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.BlockVideo(ctx, body)
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

func (h *VideoHandler) ReportVideo(ctx *gin.Context) {
	body := models.ReportVideoRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		log.Println(errors.New("failed to parse JSON data from report-video request"))
		return
	}

	res, err := h.Client.ReportVideo(ctx, body)
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

func (h *VideoHandler) GetReportedVideos(ctx *gin.Context) {

	res, err := h.Client.GetReportedVideos(ctx)
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

func (h *VideoHandler) FetchExclusiveVideo(ctx *gin.Context) {
	res, err := h.Client.FetchExclusiveVideo(ctx)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)
}
