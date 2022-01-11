package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	Id          uuid.UUID      `db:"id"`
	Name        string         `db:"name"`
	Description string         `db:"description"`
	Slug        string         `db:"slug"`
	DemoLink    sql.NullString `db:"demo_link"`
	StartAt     time.Time      `db:"start_at"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at"`
}
