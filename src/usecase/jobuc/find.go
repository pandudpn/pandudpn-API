package jobuc

import (
	"context"
	"net/http"

	"pandudpn/api/src/presenter/jobpresent"
	"pandudpn/api/src/utils/logger"
	"pandudpn/api/src/utils/response"
)

func (ju *jobUseCase) FindAllJobs(ctx context.Context) response.OutputResponseInterface {
	ctx2, dl := logger.InitializeGRPC(context.Background(), nil)
	defer dl.Finalize(ctx2)
	jobs, err := ju.jobRepo.FindAllJobs()
	if err != nil {
		return jobpresent.Response(ctx, err)
	}

	logger.Response(ctx2, http.StatusOK, jobs, nil)

	logger.Log.Debug(ctx, jobs)
	return jobpresent.Response(ctx, jobs)
}
