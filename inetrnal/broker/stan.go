package broker

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/swooosh13/L0/inetrnal/composites"
	order2 "github.com/swooosh13/L0/inetrnal/models/order"
	"github.com/swooosh13/L0/pkg/logger"
	"go.uber.org/zap"
)

func ReceiveOrder(msg *stan.Msg, orderComposite *composites.OrderComposite) {
	var o order2.Order
	err := json.Unmarshal(msg.Data, &o)
	if err != nil {
		logger.Error("Received error json format", zap.Error(err))
		return
	}

	err = o.Validate()
	if err != nil {
		logger.Error("Invalid order validation format", zap.Error(err))
		return
	}

	orderComposite.Storage.Cache.Store(o)
	orderComposite.Storage.Repo.Store(o)
	logger.Info("The order was successfully written to the DB and in-memory cache")
}
