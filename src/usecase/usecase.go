package usecase

import (
	"context"

	"pandudpn/api/src/utils/response"
)

type JobUseCaseInterface interface {
	FindAllJobs(ctx context.Context) response.OutputResponseInterface
}
