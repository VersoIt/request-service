package request

import (
	"RequestService/internal/domain/model"
	"context"
)

func (s *Service) CreateRequest(
	ctx context.Context,
	request model.Request,
	userID int64,
) (int64, error) {
	err := s.userRepo.LockUser(ctx, userID)
	if err != nil {
		return 0, err
	}

	requests, err := s.requestRepo.GetHandlingRequests(ctx, userID)
	if err != nil {
		return 0, err
	}

	if len(requests) > 0 {
		return 0, model.ErrRequestUnderConsideration
	}

	id, err := s.requestRepo.CreateRequest(ctx, request, userID)
	if err != nil {
		return 0, err
	}

	return id, nil
}
