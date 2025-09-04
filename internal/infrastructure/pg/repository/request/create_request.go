package request

import (
	"RequestService/internal/domain/model"
	"RequestService/internal/infrastructure/pg/repository/request/entity"
	"context"
)

func (r *Repository) CreateRequest(
	ctx context.Context,
	request model.Request,
	userID int64,
) (int64, error) {
	var entityRequest entity.Request
	entityRequest.FromDomain(request)

	var id int64

	err := r.txGetter.DefaultTrOrDB(ctx, r.db).GetContext(
		ctx,
		&id,
		queryCreateRequest,
		userID,
		entityRequest.Type,
		entityRequest.Payload,
		entityRequest.Status,
		entityRequest.CreatedAt,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}
