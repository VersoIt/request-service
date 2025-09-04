package server

import (
	"context"
	"errors"
)

func (s *Server) shutdown(ctx context.Context) error {
	shutdownErr := s.e.Shutdown(ctx)
	if shutdownErr != nil {
		closeErr := s.e.Close()
		if closeErr != nil {
			shutdownErr = errors.Join(shutdownErr, closeErr)
			s.e.Logger.Error(closeErr)
		} else {
			s.e.Logger.Info("Server closed forcefully")
		}

		s.e.Logger.Error(shutdownErr)
	}

	s.e.Logger.Info("Server closed")

	return shutdownErr
}
