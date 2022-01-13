package jobpresent

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"time"

	"pandudpn/api/src/utils/nullhandler"

	"github.com/google/uuid"

	"pandudpn/api/src/model"
	"pandudpn/api/src/utils/response"
)

const layoutDateTime = "02 Jan 2006"

type jobResponse struct {
	Id           uuid.UUID    `json:"id"`
	Office       string       `json:"office"`
	As           string       `json:"as"`
	StartAt      time.Time    `json:"startAt"`
	EndAt        *time.Time   `json:"endAt"`
	Description  string       `json:"description"`
	StillWorking bool         `json:"stillWorking"`
	TotalWorking float64      `json:"totalWorking"`
	Formatted    jobFormatted `json:"formatted"`
}

type jobFormatted struct {
	StartAt      string `json:"startAt"`
	EndAt        string `json:"endAt"`
	TotalWorking string `json:"totalWorking"`
}

func Response(ctx context.Context, value interface{}) response.OutputResponseInterface {
	if err, ok := value.(error); ok {
		return response.Errors(ctx, http.StatusInternalServerError, err.Error(), err)
	}

	jobExp := value.([]*model.Job)
	jobs := createJobsResponse(jobExp)
	return response.Success(ctx, http.StatusOK, jobs)
}

func createJobsResponse(jobs []*model.Job) []*jobResponse {
	var jobsres = make([]*jobResponse, 0)

	for _, job := range jobs {
		end := time.Now().UTC()
		nullTime := nullhandler.Time(job.EndAt)

		if !nullTime.IsNil() {
			end = nullTime.ValueOrZero()
		}

		totalDuration := end.Sub(job.StartAt)
		total := math.Ceil(totalDuration.Seconds() / 2600640)

		jobres := &jobResponse{
			Id:           job.Id,
			Description:  job.Description,
			EndAt:        nullTime.ValueOrZeroPtr(),
			Office:       job.Office,
			As:           job.As,
			StartAt:      job.StartAt,
			StillWorking: job.StillWorking,
			TotalWorking: total,
			Formatted: jobFormatted{
				StartAt:      job.StartAt.Format(layoutDateTime),
				EndAt:        end.Format(layoutDateTime),
				TotalWorking: fmt.Sprintf("%.0f bulan", total),
			},
		}

		if nullTime.IsNil() {
			jobres.Formatted.EndAt = "Sekarang"
		}

		jobsres = append(jobsres, jobres)
	}

	return jobsres
}
