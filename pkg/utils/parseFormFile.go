package utils

import (
	"gateway/pkg/common/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseUploadVideoForm(c *gin.Context) (models.UploadVideo, error) {
	body := models.UploadVideo{
		UserName:       c.PostForm("userName"),
		AvatarId:       c.PostForm("avatarId"),
		Title:          c.PostForm("title"),
		Discription:    c.PostForm("discription"),
		Interest:       c.PostForm("interest"),
		ThumbnailId:    c.PostForm("thumbnailId"),
		Exclusive:      false,
		Coin_for_watch: 0,
	}

	exclusiveStr := c.PostForm("exclusive")
	if exclusiveStr == "true" {
		body.Exclusive = true
	}

	coinTowatchStr := c.PostForm("coinTowatch")
	if coinTowatchStr != "" {
		coinTowatch, err := strconv.ParseUint(coinTowatchStr, 10, 32)
		if err != nil {
			log.Println("can't parse in PostForm coinTowatch")
			return models.UploadVideo{}, err
		}
		body.Coin_for_watch = uint32(coinTowatch)
	}

	return body, nil
}
