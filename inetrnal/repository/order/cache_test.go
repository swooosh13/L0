package order

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/swooosh13/L0/inetrnal/config"
	"github.com/swooosh13/L0/inetrnal/models/order"
	"github.com/swooosh13/L0/pkg/pgdb"
	"os"
	"testing"
)

var (
	valid = `
{
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
`
)

func TestCacheRepository(t *testing.T) {
	req := require.New(t)

	var o order.Order
	_ = json.Unmarshal([]byte(valid), &o)

	cache := NewCache()
	t.Run("Test store data to cache", func(t *testing.T) {
		cache.Store(o)
		t.Run("Is exists", func(t *testing.T) {
			_, ok := cache.Load(o.OrderUID)
			req.True(ok)
		})

		t.Run("Laod all length", func(t *testing.T) {
			all := cache.LoadAll()
			req.Equal(1, len(all))
		})

		t.Run("Check non-dup", func(t *testing.T) {
			cache.Store(o)
			o := cache.LoadAll()
			req.Equal(1, len(o))
		})
	})
}

func TestRecoverFunction(t *testing.T) {
	req := require.New(t)
	os.Chdir("../../../")
	cfg, _ := config.GetConfig()

	pgCfg := cfg.PostgresDB
	poolCfg, _ := pgdb.NewPoolConfig(pgCfg.Host, pgCfg.Port, pgCfg.Username, pgCfg.Password, pgCfg.Database, pgCfg.Timeout)
	poolCfg.MaxConns = int32(pgCfg.MaxConns)

	conn, _ := pgdb.NewConn(context.Background(), poolCfg)
	dbRepo := NewDbRepository(conn)

	cache := NewCache()
	t.Run("Recover cache", func(t *testing.T) {
		req.Equal(0, len(cache.LoadAll()))
		cache.Recover(dbRepo)
		req.True(len(cache.LoadAll()) > 0)
	})
}
