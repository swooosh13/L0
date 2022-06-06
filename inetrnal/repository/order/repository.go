package order

import "github.com/swooosh13/L0/inetrnal/models/order"

type Repository interface {
	Load(string) (order.Order, bool)
	LoadAll() []order.Order
	Store(order2 order.Order)
}
