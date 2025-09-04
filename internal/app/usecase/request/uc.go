package request

import "github.com/avito-tech/go-transaction-manager/trm/manager"

type UC struct {
	requestRepo    requestRepository
	requestService requestService
	manager        *manager.Manager
}

func New(
	requestRepo requestRepository,
	requestService requestService,
	txManager *manager.Manager,
) *UC {
	return &UC{
		requestService: requestService,
		requestRepo:    requestRepo,
		manager:        txManager,
	}
}
