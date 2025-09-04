package request

import (
	"RequestService/internal/domain/model"
	"context"
)

func (u *UC) GetRequest(
	ctx context.Context,
	id int64,
) (model.Request, error) {
	req, err := u.requestRepo.GetRequest(ctx, id)
	if err != nil {
		return model.Request{}, err
	}

	return req, nil
}
