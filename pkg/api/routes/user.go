package routes

import (
	"gateway/pkg/api/handlers"
	middleware "gateway/pkg/common/midleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler handlers.UserHandler, conferenceHandler handlers.ConferenceHandler) {
	routes := api.Group("/user")
	routes.POST("/signup", userHandler.UserSignup)
	routes.POST("/otp", userHandler.OtpRequest)
	routes.POST("/login", userHandler.UserLogin)
	routes.POST("/valid-name", userHandler.ValidName)

	routes.POST("/healthCheck")
	confernce := api.Group("/conference")
	confernce.Use(middleware.AuthenticateUser)
	confernce.POST("/startConference", conferenceHandler.StartConference)
	confernce.POST("/joinConference", conferenceHandler.JoinConference)
	confernce.POST("/acceptJoining", conferenceHandler.AcceptJoining)
	confernce.POST("/declineJoining", conferenceHandler.DeclineJoining)
	confernce.POST("/leaveConference", conferenceHandler.LeaveConference)
	confernce.POST("/removeParticipant", conferenceHandler.RemoveParticipant)
	confernce.POST("/toggleCamera", conferenceHandler.ToggleCamera)
	confernce.POST("/toggleMic", conferenceHandler.ToggleMic)
	confernce.POST("/toggleParticipantCamera", conferenceHandler.ToggleParticipantCamera)
	confernce.POST("/toggleParticipantMic", conferenceHandler.ToggleParticipantMic)
	confernce.POST("/endConference", conferenceHandler.EndConference)
}
