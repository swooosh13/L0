package order

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/swooosh13/L0/inetrnal/config"
	"github.com/swooosh13/L0/inetrnal/models/order"
	"github.com/swooosh13/L0/pkg/pgdb"
	"os"
	"testing"
)

func TestDbRepository(t *testing.T) {
	req := require.New(t)

	os.Chdir("../../../")
	cfg, _ := config.GetConfig()

	pgCfg := cfg.PostgresDB
	poolCfg, _ := pgdb.NewPoolConfig(pgCfg.Host, pgCfg.Port, pgCfg.Username, pgCfg.Password, pgCfg.Database, pgCfg.Timeout)
	poolCfg.MaxConns = int32(pgCfg.MaxConns)

	conn, _ := pgdb.NewConn(context.Background(), poolCfg)
	dbRepo := NewDbRepository(conn)

	var o order.Order
	_ = json.Unmarshal([]byte(valid), &o)

	t.Run("Test store data to db", func(t *testing.T) {
		t.Run("Is exists", func(t *testing.T) {
			dbRepo.Store(o)
			_, ok := dbRepo.Load(o.OrderUID)
			req.True(ok)
		})

	//	add one more
		o.OrderUID = "b563feb7b2b84b6test2"
		dbRepo.Store(o)
		t.Run("Check length", func(t *testing.T) {
			orders := dbRepo.LoadAll()
			fmt.Println(len(orders))
			req.True(len(orders) > 0)
		})
	})
}
