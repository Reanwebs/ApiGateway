package routes

import (
	"gateway/pkg/api/handlers"
	middleware "gateway/pkg/common/midleware"

	"github.com/gin-gonic/gin"
)

func ConferenceRoutes(api *gin.RouterGroup, conferenceHandler handlers.ConferenceHandler) {

	confernce := api.Group("/conference")

	confernce.POST("/healthCheck")

	confernce.Use(middleware.AuthenticateUser)

	confernce.POST("/scheduleConference", conferenceHandler.ScheduleConference)
	confernce.POST("/startPrivateConference", conferenceHandler.StartPrivateConference)
	confernce.POST("/startGroupConference", conferenceHandler.StartGroupConference)
	confernce.POST("/startPublicConference", conferenceHandler.StartPublicConference)
	confernce.PATCH("/joinPrivateConference", conferenceHandler.JoinPrivateConference)
	confernce.PATCH("/joinGroupConfernce", conferenceHandler.JoinGroupConfernce)
	confernce.PATCH("/joinPublicConference", conferenceHandler.JoinPublicConference)
	confernce.PATCH("/acceptJoining", conferenceHandler.AcceptJoining)
	confernce.DELETE("/declineJoining", conferenceHandler.DeclineJoining)
	confernce.POST("/leavePrivateConference", conferenceHandler.LeavePrivateConference)
	confernce.POST("/leaveGroupConference", conferenceHandler.LeaveGroupConference)
	confernce.POST("/leavePublicConference", conferenceHandler.LeavePublicConference)
	confernce.DELETE("/removePrivateParticipant", conferenceHandler.RemovePrivateParticipant)
	confernce.DELETE("/removeGroupParticipant", conferenceHandler.RemoveGroupParticipant)
	confernce.DELETE("/removePublicParticipant", conferenceHandler.RemovePublicParticipant)
	confernce.POST("/toggleCamera", conferenceHandler.ToggleCamera)
	confernce.POST("/toggleMic", conferenceHandler.ToggleMic)
	confernce.POST("/toggleParticipantCamera", conferenceHandler.ToggleParticipantCamera)
	confernce.POST("/toggleParticipantMic", conferenceHandler.ToggleParticipantMic)
	confernce.POST("/endPrivateConference", conferenceHandler.EndPrivateConference)
	confernce.POST("/endGroupConference", conferenceHandler.EndGroupConference)
	confernce.POST("/endPublicConference", conferenceHandler.EndPublicConference)
}
