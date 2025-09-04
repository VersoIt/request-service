package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func (s *Server) Run(ctx context.Context) error {
	if !s.started.CompareAndSwap(0, 1) {
		return errServerStarted
	}

	for _, h := range s.handlers {
		err := h.InitRoutes(s.e)
		if err != nil {
			return err
		}
	}

	s.e.Logger.Info("Server started")

	errChan := make(chan error, 1)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		err := s.e.Start(fmt.Sprintf(":%d", s.cfg.Server.Port))
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	defer func() {
		wg.Wait()
		close(errChan)
	}()

	select {
	case <-ctx.Done():
		ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*s.shutdownTimeoutSec)
		defer cancel()

		return s.shutdown(ctxTimeout)
	case err := <-errChan:
		ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*s.shutdownTimeoutSec)
		defer cancel()

		shutdownErr := s.shutdown(ctxTimeout)
		if shutdownErr != nil {
			err = errors.Join(err, shutdownErr)
		}

		return err
	}
}
