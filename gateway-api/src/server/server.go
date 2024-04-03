package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	connection "github.com/continuum-industries/microservices_interview/gateway-api/grpc_connection"
	mw "github.com/continuum-industries/microservices_interview/gateway-api/middlewares"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

var (
	port = flag.Int("port", 80, "the port for the server to listen to")
)

func StartServer() {
	log.Println("Starting the FrontendGatewayAPI")

	r := setupRouter()

	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
		Handler: r,
	}

	log.Printf("Starting serving HTTP on %s\n", srv.Addr)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	connection.CloseConnections()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	r.Use(mw.CorsMiddleware())

	r.GET("/healthz/", func(g *gin.Context) { g.Status(http.StatusOK) })

	return r
}
