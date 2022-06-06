package order

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
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
      "sale": 31,
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
	invalid1 = `{}`
	invalid2 = `{"oof_shard": "1"}`
	invalid3 = `abcecmperf`
)

func TestValidModel(t *testing.T) {
	req := require.New(t)

	var o Order
	_ = json.Unmarshal([]byte(valid), &o)

	t.Run("Check main signature", func(t *testing.T) {
		req.Equal("1", o.OofShard)
	})

	t.Run("Check items signature", func(t *testing.T) {
		req.Equal(31, o.Items[0].Sale)
	})

	t.Run("Check payment signature", func(t *testing.T) {
		req.Equal(1637907727,o.Payment.PaymentDt)
	})

	t.Run("Correct validate", func(t *testing.T) {
		err := o.Validate()
		req.Nil(err)
	})
}

func TestInvalidSignature(t *testing.T) {
	req := require.New(t)

	var o Order
	t.Run("Invalid object", func(t *testing.T) {
		_ = json.Unmarshal([]byte(invalid1), &o)
		err := o.Validate()
		req.Error(err)
	})

	t.Run("Invalid object signature", func(t *testing.T) {
		_ = json.Unmarshal([]byte(invalid2), &o)
		err := o.Validate()
		req.Error(err)
	})

	t.Run("invalid unmarshalling", func(t *testing.T) {
		err := json.Unmarshal([]byte(invalid3), &o)
		req.Error(err)
	})
}

func TestMarshalling(t *testing.T) {
	req := require.New(t)

	o := Order{}
	_, err := o.Value()
	req.Nil(err)
}