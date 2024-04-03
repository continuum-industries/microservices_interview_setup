package proxies

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func GetgRPCContext(g *gin.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	return ctx, cancel
}
