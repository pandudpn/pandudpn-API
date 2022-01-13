package visitorcontroller

import "pandudpn/api/src/usecase"

type VisitorHandler struct {
	visitorUseCase usecase.VisitorUseCaseInterface
}

func NewVisitorHandler(vu usecase.VisitorUseCaseInterface) *VisitorHandler {
	return &VisitorHandler{
		visitorUseCase: vu,
	}
}
