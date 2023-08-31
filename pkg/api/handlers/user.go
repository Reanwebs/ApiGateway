package handlers

import (
	"context"
	"errors"
	"gateway/pkg/common/client/interfaces"
	middleware "gateway/pkg/common/midleware"
	"gateway/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Client interfaces.AuthClient
}

func NewAuthHandler(client interfaces.AuthClient) UserHandler {
	return UserHandler{
		Client: client,
	}
}

// USER SIGN-UP WITH VERIFICATION OF OTP
//
//	@Summary		API FOR NEW USER SIGN UP OTP VERIFICATION
//	@ID				SIGNUP-USER-OTP-VERIFY
//	@Description	VERIFY THE OTP AND DATA INSERT TO THE DATABASE
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			user_details	body		models.RegisterRequestBody	false	"otp"
//	@Success		201				{object}	pb.SignupResponse
//	@Failure		400				{object}	pb.SignupResponse
//	@Failure		502				{object}	pb.SignupResponse
//	@Router			/api/user/signup [post]
func (h *UserHandler) UserSignup(ctx *gin.Context) {
	body := models.RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.UserSignup(context.Background(), body)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "signup failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

// USER SIGN-UP OTP REQUEST
//
//	@Summary		API FOR NEW USER SIGN UP
//	@ID				SIGNUP-USER
//	@Description	CREATE A NEW USER WITH REQUIRED DETAILS
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			user_details	body		models.OtpValidation	false	"New user Details"
//	@Success		200				{object}	pb.OtpSignUpResponse
//	@Failure		400				{object}	pb.OtpSignUpResponse
//	@Failure		500				{object}	pb.OtpSignUpResponse
//	@Router			/api/user/otp [post]
func (h *UserHandler) OtpRequest(ctx *gin.Context) {
	body := models.OtpValidation{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := h.Client.OtpRequest(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "otp sending failed",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)

}

// USER USERLOGIN
//
//	@Summary		API FOR USER LOGIN
//	@ID				USER-LOGIN
//	@Description	VERIFY THE EMAIL,PASSWORD AND GENERATE A JWT TOKEN AND SET IT TO A COOKIE
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			login_details	body		models.LoginRequestBody	true	"Enter the email and password"
//	@Success		201				{object}	pb.LoginResponse
//	@Failure		401				{object}	pb.LoginResponse
//	@Failure		500				{object}	pb.LoginResponse
//	@Failure		502				{object}	pb.LoginResponse
//	@Router			/api/user/login [post]
func (h *UserHandler) UserLogin(ctx *gin.Context) {
	body := models.LoginRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.UserLogin(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to login",
			"error":   err.Error(),
		})
		return
	}

	// setup JWT
	ok := middleware.JwtCookieSetup(ctx, "user-auth", res.Uid)
	if !ok {
		res := errors.New("failed to login")
		ctx.JSON(http.StatusInternalServerError, res)
		return

	}

	ctx.JSON(http.StatusCreated, &res)
}

// USER USERNAME VALIDATION
//
//	@Summary		API FOR USERNAME VALIDATION
//	@ID				USERNAME-VALIDATION
//	@Description	VERIFY USERNAME IF IT UNIQUE OR NOT
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			USERNAME	body		models.ValidName	true	"Enter User Name"
//	@Success		200			{object}	pb.ValidNameResponse
//	@Failure		400			{object}	pb.ValidNameResponse
//	@Failure		400			{object}	pb.ValidNameResponse
//	@Router			/api/user/valid-name [post]
func (h *UserHandler) ValidName(ctx *gin.Context) {
	body := models.ValidName{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ValidName(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "User name not available",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

// USER RESEND OTP REQUEST
//
//	@Summary		API FOR RESEND OTP REQUEST
//	@ID				RESEND-OTP
//	@Description	RESEND OTP REQUEST
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			RESENDOTP	body		models.ResendOtp	false	"enter phone number"
//	@Success		200			{object}	pb.ResendOtpResponse
//	@Failure		400			{object}	pb.ResendOtpResponse
//	@Failure		400			{object}	pb.ResendOtpResponse
//	@Router			/api/user/resend-otp [post]
func (h *UserHandler) ResendOtp(ctx *gin.Context) {
	body := models.ResendOtp{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ResendOtp(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "resending otp failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

// USERLOGOUT
//
//	@Summary		API FOR USER LOGOUT
//	@ID				USER-LOGOUT
//	@Description	LOGOUT USER AND ALSO CLEAR COOKIES
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{string}	Log	out	successful
//
//	@Router			/api/user/logout [post]
func (h *UserHandler) LogoutUser(c *gin.Context) {
	c.SetCookie("user-auth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Log out successful",
	})
}

