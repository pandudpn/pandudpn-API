package nullhandler

import "database/sql"

type NullString struct {
	sql.NullString
}

// NewStringFrom creates a new String that will never be blank.
func NewStringFrom(s string) NullString {
	if s == "" {
		return NullString{
			sql.NullString{
				Valid:  false,
				String: "",
			},
		}
	}

	return NullString{
		sql.NullString{
			Valid:  true,
			String: s,
		},
	}
}

// NewStringFromPtr creates a new String that be null if s is nil.
func NewStringFromPtr(s *string) NullString {
	if s == nil {
		return NullString{
			sql.NullString{
				Valid:  false,
				String: "",
			},
		}
	}

	return NullString{
		sql.NullString{
			Valid:  true,
			String: *s,
		},
	}
}

// ValueOrZero returns the inner value if valid, otherwise zero.
func (n *NullString) ValueOrZero() string {
	if n.Valid {
		return n.String
	}
	return ""
}

// ValueOrZeroPtr returns the inner value if valid, otherwise nil.
func (n *NullString) ValueOrZeroPtr() *string {
	if n.Valid {
		return &n.String
	}
	return nil
}

// IsNil returns true for null String, for potential future omitempty support.
func (n *NullString) IsNil() bool {
	return !n.Valid
}
