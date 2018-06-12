package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

var (
	_ Q = (*sqlx.DB)(nil)
	_ Q = (*sqlx.Tx)(nil)
)

// Q sqlx.DB or sqlx.Tx
type Q interface {
	sqlx.ExtContext
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

var conn *sqlx.DB

// Setup setup db connection
func Setup(db *sqlx.DB) {
	conn = db
}

// Invoke run statement
func Invoke(f func(db *sqlx.DB) error) error {
	return f(conn)
}

// InvokeTx run statement in transaction
func InvokeTx(f func(tx *sqlx.Tx) error) error {
	tx, err := conn.Beginx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = f(tx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
