package dto

import (
	"RequestService/internal/domain/model"
	"github.com/samber/lo"
)

type Request struct {
	Type    model.RequestType `json:"request_type"`
	Payload []byte            `json:"payload"`
	UserID  int64             `json:"user_id"`
}

func (r *Request) Domain() (model.Request, error) {
	if !lo.Contains(model.RequestTypes, r.Type) {
		return model.Request{}, model.ErrUnknownRequestType
	}

	return model.Request{
		Type:    r.Type,
		Payload: r.Payload,
		Status:  model.StatusPending,
	}, nil
}
