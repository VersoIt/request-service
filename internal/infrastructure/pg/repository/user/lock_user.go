package user

import "context"

func (r *Repository) LockUser(
	ctx context.Context,
	id int64,
) error {
	return nil
}
