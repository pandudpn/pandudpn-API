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
	FullTime  []*jobData `json:"full_time"`
	Freelance []*jobData `json:"freelance"`
}

type jobData struct {
	Id           uuid.UUID    `json:"id"`
	Office       string       `json:"office"`
	As           string       `json:"as"`
	StartAt      time.Time    `json:"start_at"`
	EndAt        *time.Time   `json:"end_at"`
	Description  string       `json:"description"`
	StillWorking bool         `json:"still_working"`
	TotalWorking float64      `json:"total_working"`
	Formatted    jobFormatted `json:"formatted"`
}

type jobFormatted struct {
	StartAt      string `json:"start_at"`
	EndAt        string `json:"end_at"`
	TotalWorking string `json:"total_working"`
}

func Response(ctx context.Context, value interface{}) response.OutputResponseInterface {
	if err, ok := value.(error); ok {
		return response.Errors(ctx, http.StatusInternalServerError, err.Error(), err)
	}

	jobExp := value.([]*model.Job)
	jobs := createJobsResponse(jobExp)
	return response.Success(ctx, http.StatusOK, jobs)
}

func createJobsResponse(jobs []*model.Job) *jobResponse {
	var jobsres = new(jobResponse)

	for _, job := range jobs {
		end := time.Now().UTC()
		nullTime := nullhandler.Time(job.EndAt)

		if !nullTime.IsNil() {
			end = nullTime.ValueOrZero()
		}

		totalDuration := end.Sub(job.StartAt)
		total := math.Ceil(totalDuration.Seconds() / 2600640)

		data := &jobData{
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
			data.Formatted.EndAt = "Sekarang"
		}

		if job.IsFullTime {
			jobsres.FullTime = append(jobsres.FullTime, data)
		} else {
			jobsres.Freelance = append(jobsres.Freelance, data)
		}
	}

	return jobsres
}
