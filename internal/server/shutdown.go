package server

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
)

func (s *Server) shutdown(ctx context.Context) error {
	shutdownErr := s.e.Shutdown(ctx)
	if shutdownErr != nil {
		closeErr := s.e.Close()
		if closeErr != nil {
			shutdownErr = errors.Join(shutdownErr, closeErr)
			log.Error(closeErr)
		} else {
			log.Info("Server closed forcefully")
		}

		log.Error(shutdownErr)
	}

	log.Info("Server closed")

	return shutdownErr
}
