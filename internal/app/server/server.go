package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Route struct {
	V1    []RouteHandler
	Other []RouteHandler
}

type RouteHandler struct {
	Method  string
	Path    string
	Handler []gin.HandlerFunc
}

type Server struct {
	port     string
	timeout  time.Duration
	engine   *gin.Engine
	instance *http.Server
	env      string
}

func (s *Server) GetInstance() *http.Server {
	return s.instance
}

func (s *Server) GetPort() string {
	return s.port
}

func (s *Server) GetEngine() *gin.Engine {
	return s.engine
}

func (s *Server) GetTimeout() int {
	return int(s.timeout)
}

// WithPort options
func WithPort(port string) func(*Server) {
	return func(s *Server) {
		s.port = port
	}
}

func WithEnv(env string) func(*Server) {
	return func(s *Server) {
		s.env = env
	}
}

// WithTimeout options
func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMiddleware(middleware ...gin.HandlerFunc) func(*Server) {
	return func(s *Server) {
		s.engine.Use(middleware...)
	}
}
