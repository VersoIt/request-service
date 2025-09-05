package request

import (
	"RequestService/config"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
)

type UC struct {
	requestRepo    requestRepository
	requestService requestService
	txManager      *manager.Manager
	kafkaProducer  kafkaProducer
	cfg            config.Config
}

func New(
	requestRepo requestRepository,
	requestService requestService,
	txManager *manager.Manager,
	kafkaProducer kafkaProducer,
	cfg config.Config,
) *UC {
	return &UC{
		requestService: requestService,
		requestRepo:    requestRepo,
		txManager:      txManager,
		kafkaProducer:  kafkaProducer,
		cfg:            cfg,
	}
}
