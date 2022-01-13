package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Visitor struct {
	Id         uuid.UUID    `db:"id"`
	Ip         string       `db:"ip"`
	UserAgent  string       `db:"user_agent"`
	TotalVisit int          `db:"total_visit"`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
}

type VisitorRequest struct {
	Ip        string `json:"ip"`
	UserAgent string `json:"userAgent"`
}
