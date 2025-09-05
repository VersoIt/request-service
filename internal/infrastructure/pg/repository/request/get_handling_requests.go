package request

import (
	"RequestService/internal/domain/model"
	"RequestService/internal/infrastructure/pg/repository/request/entity"
	"context"
)

func (r *Repository) GetHandlingRequestsWithLock(
	ctx context.Context,
	userID int64,
) (model.Requests, error) {

	var requests entity.Requests

	err := r.txGetter.DefaultTrOrDB(ctx, r.db).SelectContext(
		ctx,
		&requests,
		queryGetHandlingRequestsForUpdate,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return requests.Domain(), nil
}
