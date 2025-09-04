package request

import (
	"RequestService/internal/domain/model"
	"RequestService/internal/infrastructure/pg/repository/request/entity"
	"context"
)

func (r *Repository) GetHandlingRequests(
	ctx context.Context,
	userID int64,
) (model.Requests, error) {

	var requests entity.Requests

	err := r.txGetter.DefaultTrOrDB(ctx, r.db).GetContext(
		ctx,
		&requests,
		queryGetHandlingRequests,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return requests.Domain(), nil
}
