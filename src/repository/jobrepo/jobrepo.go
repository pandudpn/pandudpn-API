package jobrepo

import (
	"pandudpn/api/app/dbc"
)

type jobRepository struct {
	db dbc.SqlDbc
}

func NewJobRepository(db dbc.SqlDbc) *jobRepository {
	return &jobRepository{
		db: db,
	}
}
