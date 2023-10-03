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

func (h *VideoHandler) UploadVideo(c *gin.Context) {

	body, err := utils.ParseUploadVideoForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
			"error":   err.Error(),
		})
		return
	}

	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to find the file",
			"error":   err.Error(),
		})
		return
	}

	res, err := h.Client.UploadVideo(c, file, body)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to upload",
			"error":   errMsg,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Success": res,
	})
}

func (h *VideoHandler) FetchUserVideo(c *gin.Context) {
	request := models.FetchVideosRequest{}
	request.UserName = c.Query("userName")

	if request.UserName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user name required",
		})
		return
	}

	res, err := h.Client.FetchVideos(c, request)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())

		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to fetch videos",
			"error":   errMsg,
		})
		return
	}
	c.JSON(http.StatusOK, &res)

}

func (h *VideoHandler) FetchUserArchivedVideo(c *gin.Context) {

	request := models.FetchVideosRequest{}
	request.UserName = c.Query("userName")

	if request.UserName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user name required",
		})
		return
	}

	res, err := h.Client.FetchArchivedVideos(c, request)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())

		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to fetch videos",
			"error":   errMsg,
		})
		return
	}
	c.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) FetchAllVideo(c *gin.Context) {

	res, err := h.Client.FetchAllVideos(c)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())

		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to fetch videos",
			"error":   errMsg,
		})
		return
	}
	c.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) ArchivVideo(c *gin.Context) {

	body := models.ArchivedVideos{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ArchiveVideo(c, body)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())

		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to archive videos",
			"error":   errMsg,
		})
		return
	}
	c.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) GetVideoById(c *gin.Context) {

	videoId := c.Query("id")
	username := c.Query("userName")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user name required",
		})
		return
	}

	if videoId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "video id required",
		})
		return
	}

	data := models.GetVideoById{
		VideoId:  videoId,
		UserNAme: username,
	}

	res, err := h.Client.GetVideoById(c, data)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())

		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to fetch videos",
			"error":   errMsg,
		})
		return
	}
	c.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) ToggleStar(c *gin.Context) {
	body := models.ToggleStarRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ToggleStar(c, body)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   errMsg,
		})
		return
	}

	c.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) BlockVideo(c *gin.Context) {
	body := models.BlockVideoRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.BlockVideo(c, body)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   errMsg,
		})
		return
	}

	c.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) ReportVideo(ctx *gin.Context) {
	body := models.ReportVideoRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to report video",
			"error":   err,
		})
		log.Println(errors.New("fail to parse JSON data from report-video request"))
		return
	}

	res, err := h.Client.ReportVideo(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusProcessing, gin.H{
			"message": "failed to report video",
			"error":   err.Error(),
		})
		log.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) GetReportedVideos(ctx *gin.Context) {

	res, err := h.Client.GetReportedVideos(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to fetch reported video",
			"error":   err.Error(),
		})
		log.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

func (h *VideoHandler) FetchExclusiveVideo(ctx *gin.Context) {
	res, err := h.Client.FetchAllVideos(ctx)
	if err != nil {
		errMsg := utils.ExtractErrorMessage(err.Error())

		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to fetch videos",
			"error":   errMsg,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)
}
