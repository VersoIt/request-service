package dto

import (
	"RequestService/internal/domain/model"
)

type Request struct {
	ID      int64               `json:"ID"`
	Type    model.RequestType   `json:"request_type"`
	Payload []byte              `json:"payload"`
	Status  model.RequestStatus `json:"request_status"`
	UserID  int64               `json:"user_id"`
}

func (r *Request) Domain() model.Request {
	return model.Request{
		ID:      r.ID,
		Type:    r.Type,
		Payload: r.Payload,
		Status:  r.Status,
	}
}
