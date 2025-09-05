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
	GetHandlingRequestsWithLock(
		ctx context.Context,
		userID int64,
	) (model.Requests, error)
}
