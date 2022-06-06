package composites

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/swooosh13/L0/pkg/pgdb"
	"net/url"
)

type PostgresDBComposite struct {
	Client pgdb.Client
}

func NewPostgresDBComposite(ctx context.Context, host, port, username, password, dbname string, timeout, maxConns int) (*PostgresDBComposite, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d", "postgres", url.QueryEscape(username), url.QueryEscape(password), host, port, dbname, timeout)
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}
	poolConfig.MaxConns = int32(maxConns)

	conn, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	return &PostgresDBComposite{
		Client: conn,
	}, nil
}
