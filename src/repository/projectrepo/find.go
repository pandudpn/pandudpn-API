package projectrepo

import (
	"pandudpn/api/src/model"
)

const findAllProjectActive = `
SELECT id, name, description, slug, demo_link, start_at, created_at, updated_at, deleted_at
FROM projects
WHERE deleted_at IS NULL
ORDER BY start_at DESC
`

func (pr *ProjectRepository) FindAll() ([]*model.Project, error) {
	var projects = make([]*model.Project, 0)

	err := pr.db.Select(&projects, findAllProjectActive)
	return projects, err
}
