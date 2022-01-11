package projectrepo

import "pandudpn/api/app/dbc"

type ProjectRepository struct {
	db dbc.SqlDbc
}

func NewProjectRepository(db dbc.SqlDbc) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}
