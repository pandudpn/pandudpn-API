package jobrepo

import "pandudpn/api/src/model"

const queryFindAllJobs = `
SELECT id, office, start_at, end_at, description, created_at, updated_at, deleted_at
FROM jobs
WHERE deleted_at IS NULL
`

func (jr *jobRepository) FindAllJobs() ([]*model.Job, error) {
	var jobs = make([]*model.Job, 0)
	
	err := jr.db.Select(&jobs, queryFindAllJobs)
	return jobs, err
}
