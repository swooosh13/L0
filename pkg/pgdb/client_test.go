package pgdb

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/swooosh13/L0/inetrnal/config"
	"os"
	"testing"
)

func TestPoolConfig(t *testing.T) {
	req := require.New(t)
	os.Chdir("../../")
	cfg, _ := config.GetConfig()

	t.Run("Check postgres config", func(t *testing.T) {
		pgCfg := cfg.PostgresDB
		poolCfg, _ := NewPoolConfig(pgCfg.Host, pgCfg.Port, pgCfg.Username, pgCfg.Password, pgCfg.Database, pgCfg.Timeout)
		req.NotNil(poolCfg)
	})
}

func TestConnPool(t *testing.T) {
	req := require.New(t)
	os.Chdir("../../")
	cfg, _ := config.GetConfig()

	t.Run("Check postgres connection", func(t *testing.T) {
		pgCfg := cfg.PostgresDB
		poolCfg, _ := NewPoolConfig(pgCfg.Host, pgCfg.Port, pgCfg.Username, pgCfg.Password, pgCfg.Database, pgCfg.Timeout)
		poolCfg.MaxConns = int32(pgCfg.MaxConns)

		conn, _ := NewConn(context.Background(), poolCfg)
		req.NotNil(conn)
	})
}