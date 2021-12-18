package jobrepo

import "pandudpn/api/src/model"

func (jr *jobRepository) FindAllJobs() ([]*model.JobExperiences, error) {
	var jobs = make([]*model.JobExperiences, 0)

	return jobs, nil
}
