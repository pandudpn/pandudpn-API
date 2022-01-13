package visitoruc

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"pandudpn/api/src/model"
	"pandudpn/api/src/utils/logger"
	"pandudpn/api/src/utils/response"
)

func (v *VisitorUseCase) NewVisitor(ctx context.Context, req *model.VisitorRequest) response.OutputResponseInterface {
	visitor, err := v.visitorRepo.FindVisitorByIpAndUserAgent(req.Ip, req.UserAgent)
	if err != nil {
		logger.Log.Debug(ctx, "error find visitor", err)
		if errors.Is(err, sql.ErrNoRows) {
			visitor = &model.Visitor{
				UserAgent:  req.UserAgent,
				Ip:         req.Ip,
				CreatedAt:  time.Now().UTC(),
				TotalVisit: 1,
			}

			err = v.visitorRepo.NewVisitor(visitor)
			if err != nil {
				return response.Errors(ctx, errInsert.StatusCode(), errInsert.Message(), err)
			}

			return response.Success(ctx, successInsert.StatusCode(), successInsert.Message())
		}

		return response.Errors(ctx, errQuery.StatusCode(), errQuery.Message(), err)
	}

	visitor.TotalVisit += 1
	visitor.UpdatedAt = sql.NullTime{
		Valid: true,
		Time:  time.Now().UTC(),
	}

	err = v.visitorRepo.UpdateVisitor(visitor)
	if err != nil {
		logger.Log.Debugf(ctx, "error update visitor %v", err)
		return response.Errors(ctx, errUpdate.StatusCode(), errUpdate.Message(), err)
	}

	return response.Success(ctx, successUpdate.StatusCode(), successUpdate.Message())
}
