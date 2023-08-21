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

	tokenString, err := ctx.Cookie(user + "-auth") // get cookie for user or admin with name

	if err != nil || tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized User Please Login",
		})
		return
	}

	claims, err := ValidateToken(tokenString) // auth function validate the token and return claims with error
	if err != nil || tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized User Please Login token not valid",
		})
		return
	}

	// check the cliams expire time
	if time.Now().Unix() > claims.ExpiresAt {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "User Need Re-Login time expired",
		})
		return
	}
	traceId := utils.GenerateTraceID()
	// claim the" userId and set it on context
	ctx.Set("userId", fmt.Sprint(claims.Id))
	ctx.Set("traceId", traceId)

}
