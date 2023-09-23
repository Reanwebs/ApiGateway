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

	confernce.POST("/schedule-private-conference", conferenceHandler.SchedulePrivateConference)
	confernce.POST("/schedule-group-conference", conferenceHandler.ScheduleGroupConference)
	confernce.POST("/schedule-public-conference", conferenceHandler.SchedulePublicConference)
	confernce.POST("/start-private-conference", conferenceHandler.StartPrivateConference)
	confernce.POST("/start-group-conference", conferenceHandler.StartGroupConference)
	confernce.POST("/start-public-conference", conferenceHandler.StartPublicConference)
	confernce.PATCH("join-private-conference", conferenceHandler.JoinPrivateConference)
	confernce.PATCH("join-group-confernce", conferenceHandler.JoinGroupConfernce)
	confernce.PATCH("join-public-conference", conferenceHandler.JoinPublicConference)
	confernce.PATCH("/accept-joining", conferenceHandler.AcceptJoining)
	confernce.DELETE("/decline-joining", conferenceHandler.DeclineJoining)
	confernce.POST("/leave-private-conference", conferenceHandler.LeavePrivateConference)
	confernce.POST("/leave-group-conference", conferenceHandler.LeaveGroupConference)
	confernce.POST("/leave-public-conference", conferenceHandler.LeavePublicConference)
	confernce.DELETE("/remove-private-participant", conferenceHandler.RemovePrivateParticipant)
	confernce.DELETE("/remove-group-participant", conferenceHandler.RemoveGroupParticipant)
	confernce.DELETE("/remove-public-participant", conferenceHandler.RemovePublicParticipant)
	confernce.POST("/toggle-camera", conferenceHandler.ToggleCamera)
	confernce.POST("/toggle-mic", conferenceHandler.ToggleMic)
	confernce.POST("/toggle-participant-camera", conferenceHandler.ToggleParticipantCamera)
	confernce.POST("/toggle-participant-mic", conferenceHandler.ToggleParticipantMic)
	confernce.POST("/end-private-conference", conferenceHandler.EndPrivateConference)
	confernce.POST("/end-group-conference", conferenceHandler.EndGroupConference)
	confernce.POST("/end-public-conference", conferenceHandler.EndPublicConference)
	confernce.GET("/completed-schedules", conferenceHandler.CompletedSchedules)
	confernce.GET("/scheduled-conference", conferenceHandler.ScheduledConference)
	confernce.POST("/start-stream", conferenceHandler.StartStream)
	confernce.POST("/join-stream", conferenceHandler.JoinStream)
	confernce.PATCH("/leave-stream", conferenceHandler.LeaveStream)
	confernce.PATCH("/stop-stream", conferenceHandler.StopStream)
	confernce.GET("/get-stream-by-id", conferenceHandler.GetStream)
	confernce.GET("/grt-stream", conferenceHandler.GetOngoingStreams)
}
