package handlers

import (
	"context"
	"gateway/pkg/auth/common/client"
	"gateway/pkg/auth/common/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	Client client.ServiceClient
}

func NewAuthHandler(client client.ServiceClient) UserHandler {
	return UserHandler{
		Client: client,
	}
}

func UserSignup(ctx *gin.Context, c pb.AutharizationClient) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UserSignup(context.Background(), &pb.SignupRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}

func UserLogin(ctx *gin.Context, c pb.AutharizationClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UserLogin(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
