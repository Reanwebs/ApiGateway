package middleware

import (
	"fmt"
	"gateway/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// user Auth
func AuthenticateUser(ctx *gin.Context) {
	authHelper(ctx, "user")
}

// user Admin
func AuthenticateAdmin(ctx *gin.Context) {
	authHelper(ctx, "admin")
}

// helper to validate cookie and expiry
func authHelper(ctx *gin.Context, user string) {

	tokenString, err := ctx.Cookie(user + "-auth")
	if err != nil || tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized, Please Login",
		})
		return
	}

	claims, err := ValidateToken(tokenString)
	if err != nil || tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized, Please Login token not valid",
		})
		return
	}

	if time.Now().Unix() > claims.ExpiresAt {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Need Re-Login time expired",
		})
		return
	}
	traceId := utils.GenerateTraceID()
	ctx.Set("userId", fmt.Sprint(claims.Id))
	ctx.Set("traceId", traceId)

}
