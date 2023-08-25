package middleware

import (
	"errors"
	"fmt"
	"gateway/pkg/common/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// JwtCookieSetup sets up a JWT cookie
func JwtCookieSetup(c *gin.Context, name string, userId string) bool {
	// Get the configuration instance
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load configuration:", err)
		return false
	}

	// Calculate the token expiration time
	cookieTime := time.Now().Add(10 * time.Hour).Unix()

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        fmt.Sprint(userId),
		ExpiresAt: cookieTime,
	})

	// Generate a signed JWT token using the retrieved secret key
	if tokenString, err := token.SignedString([]byte(cfg.GetJWTSecretKey())); err == nil {
		// Set the cookie with the signed token string, valid for 10 hours
		c.SetCookie(name, tokenString, 10*3600, "", "", false, true)
		fmt.Println("JWT sign & set Cookie successful")
		return true
	}

	fmt.Println("Failed JWT setup")
	return false
}

func ValidateToken(tokenString string) (jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			cfg, err := config.LoadConfig()
			if err != nil {
				return nil, fmt.Errorf("error loading configuration: %w", err)
			}
			return []byte(cfg.GetJWTSecretKey()), nil
		},
	)

	if err != nil || !token.Valid {
		fmt.Println("Not a valid token")
		return jwt.StandardClaims{}, errors.New("not a valid token")
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		fmt.Println("Can't parse the claims")
		return jwt.StandardClaims{}, errors.New("can't parse the claims")
	}

	return *claims, nil
}
