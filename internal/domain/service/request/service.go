package request

type Service struct {
	requestRepo requestRepository
}

func New(requestRepo requestRepository) *Service {
	return &Service{
		requestRepo: requestRepo,
	}
}
