package entity

import (
	"RequestService/internal/domain/model"
	"time"
)

type Request struct {
	ID        int64               `db:"id"`
	Type      model.RequestType   `db:"type"`
	Payload   []byte              `db:"payload"`
	Status    model.RequestStatus `db:"status"`
	UserID    int64               `db:"user_id"`
	CreatedAt time.Time           `db:"created_at"`
}

func (r *Request) FromDomain(req model.Request) {
	r.ID = req.ID
	r.Type = req.Type
	r.Payload = req.Payload
	r.Status = req.Status
	r.CreatedAt = req.CreatedAt
}

func (r *Request) Domain() model.Request {
	return model.Request{
		ID:        r.ID,
		Type:      r.Type,
		Payload:   r.Payload,
		Status:    r.Status,
		CreatedAt: r.CreatedAt,
	}
}

type Requests []Request

func (r Requests) Domain() model.Requests {
	domainRequests := make([]model.Request, 0, len(r))
	for _, req := range r {
		domainRequests = append(domainRequests, req.Domain())
	}

	return domainRequests
}
