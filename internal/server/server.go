package server

import (
	"RequestService/cfg"
	"github.com/labstack/echo/v4"
	"sync/atomic"
	"time"
)

type Server struct {
	e                  *echo.Echo
	started            atomic.Int32
	cfg                cfg.Config
	shutdownTimeoutSec time.Duration
	handlers           []handler
}

func New(cfg cfg.Config, e *echo.Echo, handlers ...handler) *Server {
	return &Server{
		e:                  e,
		cfg:                cfg,
		shutdownTimeoutSec: time.Second * 5,
		handlers:           handlers,
	}
}
