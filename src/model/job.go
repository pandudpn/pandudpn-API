package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Job struct {
	Id           uuid.UUID    `db:"id"`
	Office       string       `db:"office"`
	StartAt      time.Time    `db:"start_at"`
	EndAt        sql.NullTime `db:"end_at"`
	Description  string       `db:"description"`
	StillWorking bool         `db:"still_working"`
	As           string       `db:"as"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}

func NewJob() *Job {
	return &Job{
		CreatedAt: time.Now().UTC(),
	}
}

func (j *Job) IsNil() bool {
	return j.Id == uuid.Nil
}
