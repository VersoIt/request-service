package http

type Handler struct {
	uc UC
}

func NewHandler(uc UC) *Handler {
	return &Handler{
		uc: uc,
	}
}
