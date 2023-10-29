package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type (
	txKey struct{}

	Transaction interface {
		InTx(ctx context.Context, fn func(ctx context.Context) error) error
	}
	transaction struct {
		*sqlx.DB
	}
)

func NewTransaction(c *Client) Transaction {
	t := transaction{c.con}
	return &t
}

func (t transaction) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := t.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	ctx = context.WithValue(ctx, txKey{}, tx)

	defer func() {
		if p := recover(); p != nil {
			if e := tx.Rollback(); e != nil {
				err = fmt.Errorf("rollback error was occurred: %w", err)
			}
		} else if err != nil {
			if e := tx.Rollback(); e != nil {
				err = fmt.Errorf("rollback error was occurred: %w", err)
			}
		} else {
			err = tx.Commit()
		}
	}()

	err = func() error {
		return fn(ctx)
	}()
	return err
}

func GetTx(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(txKey{}).(*sqlx.Tx)
	return tx, ok
}
