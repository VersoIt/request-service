package request

import (
	"RequestService/internal/domain/model"
	"context"
)

type requestRepository interface {
	GetRequest(
		ctx context.Context,
		id int64,
	) (model.Request, error)
}

type requestService interface {
	CreateRequest(
		ctx context.Context,
		request model.Request,
		userID int64,
	) (int64, error)
}

type kafkaProducer interface {
	SendMessage(ctx context.Context, topic string, key, value []byte) error
}
