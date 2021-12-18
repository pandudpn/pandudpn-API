package model

import "time"

type JobExperiences struct {
	Id           int        `db:"id"`
	Office       string     `db:"office"`
	StartAt      time.Time  `db:"start_at"`
	EndAt        *time.Time `db:"end_at"`
	Descrtiption string     `db:"description"`
	StillWorking bool       `db:"still_working"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
}
