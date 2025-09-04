package request

import (
	"RequestService/internal/domain/model"
	"context"
)

func (u *UC) CreateRequest(
	ctx context.Context,
	request model.Request,
	userID int64,
) (int64, error) {
	id, err := u.requestService.CreateRequest(ctx, request, userID)
	if err != nil {
		return 0, err
	}

	// kafka

	return id, nil
}
