package jobuc

import (
	"context"

	"pandudpn/api/src/presenter/jobpresent"
	"pandudpn/api/src/utils/logger"
	"pandudpn/api/src/utils/response"
)

func (ju *jobUseCase) FindAllJobs(ctx context.Context) response.OutputResponseInterface {
	jobs, err := ju.jobRepo.FindAllJobs()
	if err != nil {
		return jobpresent.Response(ctx, err)
	}

	logger.Log.Debug(ctx, "success")
	return jobpresent.Response(ctx, jobs)
}
