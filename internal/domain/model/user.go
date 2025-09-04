package model

type User struct {
	ID       int64
	Requests map[int64]Request
}

func NewUser(id int64, requests []Request) *User {
	requestsMap := make(map[int64]Request, len(requests))

	for _, request := range requests {
		requestsMap[request.ID] = request
	}

	return &User{
		ID:       id,
		Requests: requestsMap,
	}
}
