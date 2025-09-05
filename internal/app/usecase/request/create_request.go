package request

import (
	"RequestService/internal/domain/model"
	"context"
	"encoding/json"
)

func (u *UC) CreateRequest(
	ctx context.Context,
	request model.Request,
	userID int64,
) (int64, error) {
	err := u.txManager.Do(ctx, func(ctx context.Context) error {
		id, err := u.requestService.CreateRequest(ctx, request, userID)
		if err != nil {
			return err
		}

		request.ID = id

		requestBytes, err := json.Marshal(request)
		if err != nil {
			return err
		}

		err = u.kafkaProducer.SendMessage(
			ctx,
			u.cfg.KafkaProducer.Topic,
			[]byte(request.String()),
			requestBytes,
		)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return request.ID, nil
}
