package routes

import (
	"gateway/pkg/api/handlers"
	middleware "gateway/pkg/common/midleware"

	"github.com/gin-gonic/gin"
)

func MonitaizationRoutes(api *gin.RouterGroup, monitaizationHandlers handlers.MonitizationHandler) {
	monitaization := api.Group("/monitaization")

	monitaization.POST("/healthCheck", monitaizationHandlers.MonitaizationHealthCheck)
	monitaization.POST("/create-wallet", monitaizationHandlers.CreateWallet)
	monitaization.PUT("/update-wallet", monitaizationHandlers.UpdateWallet)

	monitaization.Use(middleware.AuthenticateUser)

	monitaization.POST("/participent-reward", monitaizationHandlers.ParticipationReward)
	monitaization.POST("/user-reward-history", monitaizationHandlers.UserRewardHistory)
	monitaization.GET("/get-wallet", monitaizationHandlers.GetWallet)
}
