package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Router struct {
	Router *gin.Engine
}

func New(options ...func(*Server)) *Server {
	// gin.DefaultWriter = io.Discard
	router := gin.Default()
	// to ensure rate limiter to work, we need to pass actual client IP address
	router.ForwardedByClientIP = true
	svr := &Server{
		engine: router,
	}
	// router.Use(logger.GinLog())

	for _, opt := range options {
		opt(svr)
	}

	return svr
}

func (s *Server) Start() error {
	ctx, cancel := context.WithCancel(context.Background())

	srv := &http.Server{
		Addr:    s.GetPort(),
		Handler: s.GetEngine(),
	}

	idleConnection := make(chan struct{})
	go func(ctx context.Context) {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		cancel()

		log.Info("Server is shutting down")
		if err := srv.Shutdown(ctx); err != nil {
			log.Error("Fail to shutting down", err)
		}
		close(idleConnection)
	}(ctx)

	log.Info("server running")

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("cannot start server with error", err)
		return err
	}
	<-idleConnection
	return nil
}

func (s *Server) RegisterRoutes(route Route) {
	v1Group := s.engine.Group("/api/v1")
	for _, route := range route.V1 {
		v1Group.Handle(route.Method, route.Path, route.Handler...)
	}
}
