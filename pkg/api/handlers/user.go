package handlers

import (
	"context"
	"errors"
	"gateway/pkg/common/client/interfaces"
	middleware "gateway/pkg/common/midleware"
	"gateway/pkg/common/models"
	"net/http"
	"time"

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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.UserSignup(context.Background(), body, retryConfig)

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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.OtpRequest(ctx, body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.UserLogin(ctx, body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.ValidName(ctx, body, retryConfig)
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

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.ResendOtp(ctx, body, retryConfig)
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
func (h *UserHandler) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("user-auth", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Log out successful",
	})
}

// USER FORGOT PASSWORD OTP REQUEST
//
//	@Summary		API FOR FORGOT PASSWORD OTP REQUEST
//	@ID				FORGOT-PASSWORD-REQUEST
//	@Description	FORGOT PASSWORD OTP REQUEST
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			FORGOT-PASSWORD-REQUEST	body		models.ForgotPasswordOtpRequest	false	"enter phone number"
//	@Success		200						{object}	pb.ForgotPasswordOtpResponse
//	@Failure		400						{object}	pb.ForgotPasswordOtpResponse
//	@Failure		400						{object}	pb.ForgotPasswordOtpResponse
//	@Router			/api/user/forgot-pass-otp [post]
func (h *UserHandler) ForgotPasswordOtp(ctx *gin.Context) {
	body := models.ForgotPasswordOtpRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.ForgotPasswordOtp(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "sending otp failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

// USER FORGOT PASSWORD OTP VALIDATION
//
//	@Summary		API FOR FORGOT PASSWORD OTP VALIDATION
//	@ID				FORGOT-PASSWORD-VALIDATION
//	@Description	FORGOT PASSWORD OTP VALIDATION
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			FORGOT-PASSWORD-VALIDATION	body		models.ForgotPasswordValidateOtpRequest	false	"enter phone number"
//	@Success		200							{object}	pb.ForgotPasswordValidateOtpResponse
//	@Failure		400							{object}	pb.ForgotPasswordValidateOtpResponse
//	@Failure		400							{object}	pb.ForgotPasswordValidateOtpResponse
//	@Router			/api/user/forgot-pass-validate [post]
func (h *UserHandler) ForgotPasswordValidateOtp(ctx *gin.Context) {
	body := models.ForgotPasswordValidateOtpRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.ForgotPasswordValidateOtp(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "otp validation failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// USER FORGOT PASSWORD RESET PASSWORD
//
//	@Summary		API FOR RESETING PASSWORD AFTER VERIFICATION
//	@ID				RESET-PASSWORD
//	@Description	RESETING PASSWORD AFTER VERIFICATION
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			RESET-PASSWORD	body		models.ForgotPasswordChangePasswordRequest	false	"enter new password"
//	@Success		200				{object}	pb.ForgotPasswordOtpResponse
//	@Failure		400				{object}	pb.ForgotPasswordOtpResponse
//	@Failure		400				{object}	pb.ForgotPasswordOtpResponse
//	@Router			/api/user/forgot-pass-reset [patch]
func (h *UserHandler) ForgotPasswordChangePassword(ctx *gin.Context) {
	body := models.ForgotPasswordChangePasswordRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.ForgotPasswordChangePassword(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "otp verification failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// USER VALIDATE USER
//
//	@Summary		API FOR VALIDATE USER
//	@ID				VALIDATE USER
//	@Description	VALIDATING USER BLOCKED OR NOT
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			VALIDATE-USER	body		models.ValidateUserRequest	false	"enter mail"
//	@Success		200				{object}	pb.ValidateUserResponse
//	@Failure		400				{object}	pb.ValidateUserResponse
//	@Failure		400				{object}	pb.ValidateUserResponse
//	@Router			/api/user/validate-user [post]
func (h *UserHandler) ValidateUser(ctx *gin.Context) {
	body := models.ValidateUserRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.ValidateUser(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "validate user failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

// USER GOOGLE LOGIN
//
//	@Summary		API FOR GOOGLE LOGIN
//	@ID				GOOGLE LOGIN
//	@Description	USER CAN LOGIN USING GMAIL
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			GOOGLE-LOGIN	body		models.GoogleLoginRequest	false	"enter gmail"
//	@Success		200				{object}	pb.GoogleLoginResponse
//	@Failure		400				{object}	pb.GoogleLoginResponse
//	@Failure		400				{object}	pb.GoogleLoginResponse
//	@Router			/api/user/google-login [post]
func (h *UserHandler) GoogleLogin(ctx *gin.Context) {
	body := models.GoogleLoginRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.GoogleLogin(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "google login failed",
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

	ctx.JSON(http.StatusOK, &res)
}

// USER CHANGE USER NAME
//
//	@Summary		API FOR CHANGE USER NAME
//	@ID				CHANGE USER NAME
//	@Description	USER CAN CHANGE USER NAME
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			CHANGE-USER-NAME	body		models.ChangeUserNameRequest	false	"enter user name"
//	@Success		200					{object}	pb.ChangeUserNameResponse
//	@Failure		400					{object}	pb.ChangeUserNameResponse
//	@Failure		400					{object}	pb.ChangeUserNameResponse
//	@Router			/api/user/change-user-name [patch]
func (h *UserHandler) ChangeUserName(ctx *gin.Context) {
	body := models.ChangeUserNameRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.ChangeUserName(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to change user name",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// USER CHANGE EMAIL
//
//	@Summary		API FOR CHANGE EMAIL
//	@ID				CHANGE EMAIL
//	@Description	USER CAN CHANGE EMAIL
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			CHANGE-EMAIL	body		models.ChangeEmailRequest	false	"enter email"
//	@Success		200				{object}	pb.ChangeEmailResponse
//	@Failure		400				{object}	pb.ChangeEmailResponse
//	@Failure		400				{object}	pb.ChangeEmailResponse
//	@Router			/api/user/change-email[patch]
func (h *UserHandler) ChangeEmail(ctx *gin.Context) {
	body := models.ChangeEmailRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.ChangeEmail(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to change email",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// USER CHANGE PASSWORD
//
//	@Summary		API FOR CHANGE PASSWORD
//	@ID				CHANGE PASSWORD
//	@Description	USER CAN CHANGE PASSWORD
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			CHANGE-PASSWORD	body		models.ChangePasswordRequest	false	"enter password"
//	@Success		200				{object}	pb.ChangePasswordResponse
//	@Failure		400				{object}	pb.ChangePasswordResponse
//	@Failure		400				{object}	pb.ChangePasswordResponse
//	@Router			/api/user/change-password[patch]
func (h *UserHandler) ChangePassword(ctx *gin.Context) {
	body := models.ChangePasswordRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.ChangePassword(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to change password",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// USER CHANGE EMAIL OTP
//
//	@Summary		API FOR CHANGE EMAIL OTP
//	@ID				CHANGE EMAIL OTP
//	@Description	CHANGE EMAIL OTP VERIFICATION
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			EMAIL-OTP	body		models.ChangeEmailVerifyOtpRequest	false	"enter otp"
//	@Success		200			{object}	pb.ChangeEmailVerifyOtpResponse
//	@Failure		400			{object}	pb.ChangeEmailVerifyOtpResponse
//	@Failure		400			{object}	pb.ChangeEmailVerifyOtpResponse
//	@Router			/api/user/change-email-verify-otp[post]
func (h *UserHandler) ChangeEmailVerifyOtp(ctx *gin.Context) {
	body := models.ChangeEmailVerifyOtpRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.ChangeEmailVerifyOtp(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "otp validation failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// USER CHANGE PHONE NUMBER OTP
//
//	@Summary		API FOR CHANGE PHONE NUMBER OTP
//	@ID				CHANGE PHONE NUMBER OTP
//	@Description	CHANGE PHONE NUMBER OTP VERIFICATION
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			PHONE-NUMBER-OTP	body		models.ChangePhoneNumberOtpRequest	false	"enter otp"
//	@Success		200					{object}	pb.ChangePhoneNumberOtpResponse
//	@Failure		400					{object}	pb.ChangePhoneNumberOtpResponse
//	@Failure		400					{object}	pb.ChangePhoneNumberOtpResponse
//	@Router			/api/user/change-phone-number-verify-otp[post]
func (h *UserHandler) ChangePhoneNumberOtp(ctx *gin.Context) {
	body := models.ChangePhoneNumberOtpRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.ChangePhoneNumberOtp(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "otp validation failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// USER CHANGE PHONE NUMBER
//
//	@Summary		API FOR CHANGE PHONE NUMBER
//	@ID				CHANGE PHONE NUMBER
//	@Description	USER CAN CHANGE PHONE NUMBER
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			PHONE-NUMBER	body		models.ChangePhoneNumberOtpRequest	false	"enter phone number"
//	@Success		200				{object}	pb.ChangePhoneNumberRequest
//	@Failure		400				{object}	pb.ChangePhoneNumberRequest
//	@Failure		400				{object}	pb.ChangePhoneNumberRequest
//	@Router			/api/user/change-phone-number[patch]
func (h *UserHandler) ChangePhoneNumber(ctx *gin.Context) {
	body := models.ChangePhoneNumberRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.ChangePhoneNumber(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to change phone number",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// USER CHANGE AVATAR
//
//	@Summary		API FOR CHANGE AVATAR
//	@ID				CHANGE AVATAR
//	@Description	USER CAN CHANGE AVATAR IMAGE
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			CHANGE-AVATAR	body		models.ChangeAvatarRequest	false	"needed avatar id"
//	@Success		200				{object}	pb.ChangeAvatarResponse
//	@Failure		400				{object}	pb.ChangeAvatarResponse
//	@Failure		400				{object}	pb.ChangeAvatarResponse
//	@Router			/api/user/change-avatar[patch]
func (h *UserHandler) ChangeAvatar(ctx *gin.Context) {
	body := models.ChangeAvatarRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.ChangeAvatar(ctx, body, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't change avatar",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

// USER REMOVE AVATAR
//
//	@Summary		API FOR REMOVE AVATAR
//	@ID				REMOVE AVATAR
//	@Description	USER CAN REMOVE AVATAR IMAGE
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	pb.ChangeAvatarResponse
//	@Failure		400	{object}	pb.ChangeAvatarResponse
//	@Router			/api/user/delete-avatar[patch]
func (h *UserHandler) RemoveAvatar(ctx *gin.Context) {
	retryConfig := models.RetryConfig{
		MaxRetries:    5,
		MaxDuration:   3 * time.Second,
		RetryInterval: 1 * time.Second, // Wait 1 second between retries
	}

	res, err := h.Client.RemoveAvatar(ctx, retryConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "remove avatar failed",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *UserHandler) CreateCommunity(ctx *gin.Context) {
	body := models.CreateCommunityRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.CreateCommunity(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to create community",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)

}

func (h *UserHandler) JoinCommunity(ctx *gin.Context) {
	body := models.JoinCommunityRequest{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.JoinCommunity(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to join community",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *UserHandler) LeaveCommunity(ctx *gin.Context) {
	body := models.LeaveCommunityRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.LeaveCommunity(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to leave community",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *UserHandler) AcceptJoinCommunity(ctx *gin.Context) {
	body := models.AcceptJoinCommunityRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.AcceptJoinCommunity(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to accept joining community",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *UserHandler) RemoveMember(ctx *gin.Context) {
	body := models.RemoveMemberRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.RemoveMember(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to remove member",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *UserHandler) AddModerator(ctx *gin.Context) {
	body := models.AddModeratorRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.AddModerator(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to add moderator",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *UserHandler) AddMember(ctx *gin.Context) {
	body := models.AddMemberRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.AddMember(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to add member",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *UserHandler) ChangeCommunityJoinType(ctx *gin.Context) {
	body := models.ChangeCommunityJoinTypeRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.ChangeCommunityJoinType(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to change community join type",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *UserHandler) DeleteCommunity(ctx *gin.Context) {
	body := models.DeleteCommunityRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.DeleteCommunity(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to delete community",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (h *UserHandler) GetInterstsUser(ctx *gin.Context) {

	res, err := h.Client.GetInterstsUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to fetch intrests",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)

}
