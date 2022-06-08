package broker

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/swooosh13/L0/inetrnal/composites"
	order2 "github.com/swooosh13/L0/inetrnal/models/order"
)

func ReceiveOrder(msg *stan.Msg, orderComposite *composites.OrderComposite) {
	var o order2.Order
	err := json.Unmarshal(msg.Data, &o)
	if err != nil {
		fmt.Println("received error json format")
		return
	}

	err = o.Validate()
	if err != nil {
		fmt.Println("invalid order validation format")
		return
	}

	orderComposite.Storage.Cache.Store(o)
	orderComposite.Storage.Repo.Store(o)

	fmt.Println("received correct order")
}
