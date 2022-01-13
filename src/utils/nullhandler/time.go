package nullhandler

import (
	"database/sql"
	"time"
)

type NullTime struct {
	sql.NullTime
}

func Time(t sql.NullTime) NullTime {
	return NullTime{
		t,
	}
}

// NewTimeFrom creates a new Int64 that will never be blank.
func NewTimeFrom(t time.Time) NullTime {
	if t.IsZero() {
		return NullTime{
			sql.NullTime{
				Valid: false,
				Time:  time.Time{},
			},
		}
	}

	return NullTime{
		sql.NullTime{
			Valid: true,
			Time:  t,
		},
	}
}

// NewTimeFromPtr creates a new Time that be null if s is nil.
func NewTimeFromPtr(t *time.Time) NullTime {
	if t == nil {
		return NullTime{
			sql.NullTime{
				Valid: false,
				Time:  time.Time{},
			},
		}
	}

	return NullTime{
		sql.NullTime{
			Valid: true,
			Time:  *t,
		},
	}
}

// ValueOrZero returns the inner value if valid, otherwise zero.
func (n NullTime) ValueOrZero() time.Time {
	if n.Valid {
		return n.Time
	}
	return time.Time{}
}

// ValueOrZeroPtr returns the inner value if valid, otherwise nil.
func (n NullTime) ValueOrZeroPtr() *time.Time {
	if n.Valid {
		return &n.Time
	}
	return nil
}

// IsNil returns true for null time.Time, for potential future omitempty support.
func (n NullTime) IsNil() bool {
	return !n.Valid
}
