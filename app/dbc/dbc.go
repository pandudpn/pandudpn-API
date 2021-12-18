// Package dbc is (Database Connection) is used for create to represents low level database interfaces
// in order to have an unified way to access database handler
package dbc

import (
	"pandudpn/api/app/dbc/sql"

	"github.com/jmoiron/sqlx"
)

// SqlDbc (SQL Database Connection) is a wrapper for SQL Database handler (can be *sqlx.DB or *sqlx.Tx)
type SqlDbc interface {
	Preparex(query string) (*sqlx.Stmt, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	// create transaction
	Begin() *sql.Tx
	// If you want support transactional
	Transactioner
}

type Transactioner interface {
	// Rollback a transaction
	Rollback() error
	// Commit a transaction
	Commit() error
	// TxEnd commits a transaction if no errors, otherwise callback
	// txFunc is the operations wrapped in a transaction
	TxEnd(txFunc func() error) error
}
