package routes

import (
	"gateway/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler handlers.UserHandler) {
	routes := api.Group("/user")
	routes.POST("/signup", userHandler.UserSignup)
	routes.POST("/otp", userHandler.OtpRequest)
	routes.POST("/login", userHandler.UserLogin)
	routes.POST("/valid-name", userHandler.ValidName)
	routes.POST("/logout", userHandler.LogoutUser)
	routes.POST("/resend-otp", userHandler.ResendOtp)
	routes.POST("/forgot-pass-otp", userHandler.ForgotPasswordOtp)
	routes.POST("/forgot-pass-validate", userHandler.ForgotPasswordValidateOtp)
	routes.PATCH("/forgot-pass-reset", userHandler.ForgotPasswordChangePassword)
	routes.POST("/validate-user", userHandler.ValidateUser)
	routes.POST("/google-login", userHandler.GoogleLogin)
	routes.PATCH("/change-user-name", userHandler.ChangeUserName)
	routes.PATCH("/change-email", userHandler.ChangeEmail)
	routes.PATCH("/change-password", userHandler.ChangePassword)
	routes.POST("/change-email-verify-otp", userHandler.ChangeEmailVerifyOtp)
	routes.POST("/change-phone-number-verify-otp", userHandler.ChangePhoneNumberOtp)
	routes.PATCH("/change-phone-number", userHandler.ChangePhoneNumber)
}
