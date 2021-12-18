package repository

import "pandudpn/api/src/model"

type JobRepositoryInterface interface {
	FindAllJobs() ([]*model.JobExperiences, error)
}
