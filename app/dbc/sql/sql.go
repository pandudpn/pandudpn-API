package sql

import (
	"github.com/jmoiron/sqlx"
)

type Db struct {
	DB *sqlx.DB
}

type Tx struct {
	DB *sqlx.Tx
}

func (s *Db) Preparex(query string) (*sqlx.Stmt, error) {
	return s.DB.Preparex(query)
}

func (s *Db) Get(dest interface{}, query string, args ...interface{}) error {
	return s.DB.Get(dest, query, args...)
}

func (s *Db) Select(dest interface{}, query string, args ...interface{}) error {
	return s.DB.Select(dest, query, args...)
}

func (s *Db) Begin() *Tx {
	tx := s.DB.MustBegin()

	return &Tx{
		DB: tx,
	}
}

func (s *Tx) Preparex(query string) (*sqlx.Stmt, error) {
	return s.DB.Preparex(query)
}

func (s *Tx) Get(dest interface{}, query string, args ...interface{}) error {
	return s.DB.Get(dest, query, args...)
}

func (s *Tx) Select(dest interface{}, query string, args ...interface{}) error {
	return s.DB.Select(dest, query, args...)
}

func (s *Tx) Begin() *Tx {
	return nil
}
