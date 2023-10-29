package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/K-Kizuku/Somniosus_back/apps/account/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Client : PostgreSQLクライアント
type (
	Client struct {
		con *sqlx.DB
	}

	Do interface {
		sqlx.ExtContext
		sqlx.QueryerContext
		NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
		SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
		GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	}

	dbConfigParams struct {
		DBName string
		DBUser string
		DBPass string
		DBHost string
	}
)

// NewClient : DB用クライアント生成
func NewClient() (*Client, error) {
	params := &dbConfigParams{
		DBName: config.DbName,
		DBUser: config.DbUser,
		DBPass: config.DbPass,
		DBHost: config.DbHost,
	}

	con, err := sqlx.Connect("postgres", connectDSN(params))
	if err != nil {
		return nil, fmt.Errorf("sqlx open error(Write): %w", err)
	}

	return &Client{
		con: con,
	}, nil
}

func (c *Client) Close() {
	_ = c.con.Close()
}

func connectDSN(params *dbConfigParams) string {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", params.DBHost, params.DBUser, params.DBPass, params.DBName)
	return connStr
}

func (c *Client) Do(ctx context.Context) Do {
	if tx, ok := GetTx(ctx); ok {
		return tx
	}
	return c.con
}
