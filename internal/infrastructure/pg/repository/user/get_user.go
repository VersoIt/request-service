package user

import (
	"RequestService/internal/domain/model"
	"context"
)

func (r *Repository) GetUser(
	ctx context.Context,
	id int64,
) (model.User, error) {
	return model.User{}, nil
}
