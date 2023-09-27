package handlers

import (
	"fmt"
	"gateway/pkg/common/client/interfaces"
	"gateway/pkg/common/models"
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

func (cr *VideoHandler) UploadVideo(c *gin.Context) {

	userName := c.PostForm("userName")
	avatarId := c.PostForm("avatarId")
	title := c.PostForm("title")
	discription := c.PostForm("discription")
	interest := c.PostForm("interest")
	thumbnailId := c.PostForm("thumbnailId")

	// Create an instance of UploadVideo struct and populate it
	body := models.UploadVideo{
		UserName:    userName,
		AvatarId:    avatarId,
		Title:       title,
		Discription: discription,
		Interest:    interest,
		ThumbnailId: thumbnailId,
	}

	file, err := c.FormFile("video")
	if err != nil {
		fmt.Println(file)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to find the file",
			"error":   err.Error(),
		})
		return
	}
	res, err1 := cr.Client.UploadVideo(c.Request.Context(), file, body)
	if err1 != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to upload",
			"error":   err1.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": res,
	})
}

func (cr *VideoHandler) FetchUserVideo(c *gin.Context) {
	request := models.FetchVideosRequest{}
	request.UserName = c.Query("userName")

	if request.UserName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user name required",
		})
		return
	}

	res, err := cr.Client.FetchVideos(c, request)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to fetch videos",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &res)

}

func (cr *VideoHandler) FetchUserArchivedVideo(c *gin.Context) {

	request := models.FetchVideosRequest{}
	request.UserName = c.Query("userName")

	if request.UserName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user name required",
		})
		return
	}

	res, err := cr.Client.FetchArchivedVideos(c, request)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to fetch videos",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &res)
}

func (cr *VideoHandler) FetchAllVideo(c *gin.Context) {

	res, err := cr.Client.FetchAllVideos(c)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to fetch videos",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &res)
}

func (cr *VideoHandler) ArchivVideo(c *gin.Context) {

	body := models.ArchivedVideos{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := cr.Client.ArchiveVideo(c, body)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to archive videos",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &res)
}
