package routes

import (
	"gateway/pkg/api/handlers"
	middleware "gateway/pkg/common/midleware"

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

	routes.Use(middleware.AuthenticateUser)

	routes.PATCH("/change-user-name", userHandler.ChangeUserName)
	routes.PATCH("/change-email", userHandler.ChangeEmail)
	routes.PATCH("/change-password", userHandler.ChangePassword)
	routes.POST("/change-email-verify-otp", userHandler.ChangeEmailVerifyOtp)
	routes.POST("/change-phno-request", userHandler.ChangePhoneNumberOtp)
	routes.PATCH("/change-phno", userHandler.ChangePhoneNumber)
	routes.PATCH("/change-avatar", userHandler.ChangeAvatar)
	routes.PATCH("/delete-avatar", userHandler.RemoveAvatar)
	routes.POST("/create-community", userHandler.CreateCommunity)
	routes.PATCH("/join-community", userHandler.JoinCommunity)
	routes.PATCH("/leave-community", userHandler.LeaveCommunity)
	routes.PATCH("/accept-join-community", userHandler.AcceptJoinCommunity)
	routes.PATCH("/remove-member", userHandler.RemoveMember)
	routes.POST("/add-moderator", userHandler.AddModerator)
	routes.POST("/add-member", userHandler.AddMember)
	routes.PATCH("/change-community-join-type", userHandler.ChangeCommunityJoinType)
	routes.PATCH("/delete-community", userHandler.DeleteCommunity)
	routes.GET("/get-interests", userHandler.GetInterstsUser)
	routes.GET("/get-user-by-name", userHandler.GetUserByName)
	routes.GET("/get-active-community", userHandler.GetActiveCommunity)
	routes.GET("/get-community-by-id", userHandler.GetCommunityById)
	routes.GET("/validate-community-name", userHandler.ValidateCommunityName)
	routes.GET("/get-user-details", userHandler.GetUserDetails)
	routes.GET("/get-joined-community", userHandler.GetJoinedCommunity)
	routes.GET("/search-community", userHandler.SearchCommunity)
}
