package request

import (
	"RequestService/internal/domain/model"
	"context"
)

type UC interface {
	GetRequest(
		ctx context.Context,
		id int64,
	) (model.Request, error)
	CreateRequest(
		ctx context.Context,
		request model.Request,
		userID int64,
	) (int64, error)
}
