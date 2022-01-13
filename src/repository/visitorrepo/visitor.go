package visitorrepo

import "pandudpn/api/app/dbc"

type VisitorRepository struct {
	db dbc.SqlDbc
}

func NewVisitorRepository(db dbc.SqlDbc) *VisitorRepository {
	return &VisitorRepository{
		db: db,
	}
}
