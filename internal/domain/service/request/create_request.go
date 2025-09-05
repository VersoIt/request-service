package request

import (
	"RequestService/internal/domain/model"
	"context"
	"time"
)

func (s *Service) CreateRequest(
	ctx context.Context,
	request model.Request,
	userID int64,
) (int64, error) {
	requests, err := s.requestRepo.GetHandlingRequestsWithLock(ctx, userID)
	if err != nil {
		return 0, err
	}

	if len(requests) > 0 {
		return 0, model.ErrRequestUnderConsideration
	}

	request.CreatedAt = time.Now()

	id, err := s.requestRepo.CreateRequest(ctx, request, userID)
	if err != nil {
		return 0, err
	}

	return id, nil
}
