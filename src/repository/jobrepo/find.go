package jobrepo

import "pandudpn/api/src/model"

const (
	queryFindAllJobs = "select * from job_experiences"
)

func (jr *jobRepository) FindAllJobs() ([]*model.JobExperiences, error) {
	var jobs = make([]*model.JobExperiences, 0)

	err := jr.db.Select(&jobs, queryFindAllJobs)
	return jobs, err
}
