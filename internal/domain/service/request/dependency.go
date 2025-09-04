package request

import (
	"RequestService/internal/domain/model"
	"context"
)

type requestRepository interface {
	CreateRequest(
		ctx context.Context,
		request model.Request,
		userID int64,
	) (int64, error)
	GetRequest(
		ctx context.Context,
		id int64,
	) (model.Request, error)
	GetHandlingRequests(
		ctx context.Context,
		userID int64,
	) (model.Requests, error)
}

type userRepository interface {
	GetUser(
		ctx context.Context,
		id int64,
	) (model.User, error)
	LockUser(
		ctx context.Context,
		id int64,
	) error
}
