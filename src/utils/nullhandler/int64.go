package nullhandler

import "database/sql"

type NullInt64 struct {
	sql.NullInt64
}

// NewIntFrom creates a new Int64 that will never be blank.
func NewIntFrom(f int64) NullInt64 {
	if f == 0 {
		return NullInt64{
			sql.NullInt64{
				Valid: false,
				Int64: 0,
			},
		}
	}

	return NullInt64{
		sql.NullInt64{
			Valid: true,
			Int64: f,
		},
	}
}

// NewIntFromPtr creates a new Int64 that be null if s is nil.
func NewIntFromPtr(f *int64) NullInt64 {
	if f == nil {
		return NullInt64{
			sql.NullInt64{
				Valid: false,
				Int64: 0,
			},
		}
	}

	return NullInt64{
		sql.NullInt64{
			Valid: true,
			Int64: *f,
		},
	}
}

// ValueOrZero returns the inner value if valid, otherwise zero.
func (n *NullInt64) ValueOrZero() int64 {
	if n.Valid {
		return n.Int64
	}
	return 0
}

// ValueOrZeroPtr returns the inner value if valid, otherwise nil.
func (n *NullInt64) ValueOrZeroPtr() *int64 {
	if n.Valid {
		return &n.Int64
	}
	return nil
}

// IsNil returns true for null Integer, for potential future omitempty support.
func (n *NullInt64) IsNil() bool {
	return !n.Valid
}
