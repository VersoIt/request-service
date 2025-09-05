package request

import (
	"RequestService/internal/domain/model"
	"RequestService/internal/infrastructure/pg/repository/request/entity"
	"context"
	"database/sql"
	"errors"
)

func (r *Repository) GetRequest(
	ctx context.Context,
	id int64,
) (model.Request, error) {
	var request entity.Request

	err := r.txGetter.DefaultTrOrDB(ctx, r.db).GetContext(
		ctx,
		&request,
		queryGetRequest,
		id,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Request{}, model.ErrRequestNotFound
	}

	if err != nil {
		return model.Request{}, err
	}

	return request.Domain(), nil
}
