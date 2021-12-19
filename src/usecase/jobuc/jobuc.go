package jobuc

import "pandudpn/api/src/repository"

type jobUseCase struct {
	jobRepo repository.JobRepositoryInterface
}

func NewJobUseCase(jobRepo repository.JobRepositoryInterface) *jobUseCase {
	return &jobUseCase{
		jobRepo: jobRepo,
	}
}
