package jobcontroller

import (
	"pandudpn/api/src/usecase"
)

type jobController struct {
	jobUseCase usecase.JobUseCaseInterface
}

func NewJobController(jobUseCase usecase.JobUseCaseInterface) *jobController {
	return &jobController{
		jobUseCase: jobUseCase,
	}
}
