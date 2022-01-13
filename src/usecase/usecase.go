package usecase

import (
	"context"

	"pandudpn/api/src/model"
	"pandudpn/api/src/utils/response"
)

type JobUseCaseInterface interface {
	FindAllJobs(ctx context.Context) response.OutputResponseInterface
}

type VisitorUseCaseInterface interface {
	NewVisitor(ctx context.Context, req *model.VisitorRequest) response.OutputResponseInterface
}
