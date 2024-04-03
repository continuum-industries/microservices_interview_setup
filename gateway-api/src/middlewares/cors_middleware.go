package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	AllowedOrigins map[string]bool
)

func init() {
	UseDefaultCORSSettings()
}

func UseDefaultCORSSettings() {
	appHostName := os.Getenv("APP_HOSTNAME")
	AllowedOrigins = map[string]bool{
		fmt.Sprintf("https://%s", appHostName): true,
		appHostName:                            true,
	}
	if strings.Contains(appHostName, "dev") {
		AllowedOrigins[fmt.Sprintf("http://%s", appHostName)] = true
		AllowedOrigins["http://localhost:8080"] = true
		AllowedOrigins["http://localhost:8081"] = true
		AllowedOrigins["http://localhost:8090"] = true
		AllowedOrigins["https://localhost:8080"] = true
		AllowedOrigins["https://localhost:8081"] = true
		AllowedOrigins["https://localhost:8090"] = true
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		origin := g.Request.Header.Get("Origin")
		if _, ok := AllowedOrigins[origin]; !ok {
			g.AbortWithError(403, fmt.Errorf("Origin %s is not an allowed origin", origin))
			return
		}

		g.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		g.Writer.Header().Set("Access-Control-Max-Age", "86400")
		g.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
		g.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		g.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		g.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if g.Request.Method == "OPTIONS" {
			g.AbortWithStatus(200)
		} else {
			g.Next()
		}
	}
}
