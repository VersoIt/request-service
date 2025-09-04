package request

type Service struct {
	requestRepo requestRepository
	userRepo    userRepository
}

func New(requestRepo requestRepository, userRepo userRepository) *Service {
	return &Service{
		requestRepo: requestRepo,
		userRepo:    userRepo,
	}
}
