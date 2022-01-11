package nullhandler

import "database/sql"

type NullFloat64 struct {
	sql.NullFloat64
}

// NewFloatFrom creates a new Float64 that will never be blank.
func NewFloatFrom(f float64) NullFloat64 {
	if f == 0 {
		return NullFloat64{
			sql.NullFloat64{
				Valid:   false,
				Float64: 0,
			},
		}
	}

	return NullFloat64{
		sql.NullFloat64{
			Valid:   true,
			Float64: f,
		},
	}
}

// NewFloatFromPtr creates a new Float64 that be null if s is nil.
func NewFloatFromPtr(f *float64) NullFloat64 {
	if f == nil {
		return NullFloat64{
			sql.NullFloat64{
				Valid:   false,
				Float64: 0,
			},
		}
	}

	return NullFloat64{
		sql.NullFloat64{
			Valid:   true,
			Float64: *f,
		},
	}
}

// ValueOrZero returns the inner value if valid, otherwise zero.
func (n *NullFloat64) ValueOrZero() float64 {
	if n.Valid {
		return n.Float64
	}
	return 0
}

// ValueOrZeroPtr returns the inner value if valid, otherwise nil.
func (n *NullFloat64) ValueOrZeroPtr() *float64 {
	if n.Valid {
		return &n.Float64
	}
	return nil
}

// IsNil returns true for null Float, for potential future omitempty support.
func (n *NullFloat64) IsNil() bool {
	return !n.Valid
}
