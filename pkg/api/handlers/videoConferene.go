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
	body := models.UploadVideo{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse JSON data",
			"error":   err.Error(),
		})
		return
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
