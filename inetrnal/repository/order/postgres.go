package order

import (
	"context"
	"fmt"
	"github.com/swooosh13/L0/inetrnal/models/order"
	"github.com/swooosh13/L0/pkg/pgdb"
)

type DbRepository struct {
	client pgdb.Client
}

func NewDbRepository(client pgdb.Client) *DbRepository {
	return &DbRepository{
		client: client,
	}
}

func (r *DbRepository) Load(key string) (order.Order, bool) {
	q := `
		select data from orders where order_uid = $1; 
	`

	var data order.Order
	err := r.client.QueryRow(context.Background(), q, key).Scan(&data)
	if err != nil {
		return order.Order{}, false
	}

	return data, true
}

func (r *DbRepository) LoadAll() []order.Order {
	q := `
		select data from orders;
	`

	orders := make([]order.Order, 0)
	rows, err := r.client.Query(context.Background(),q)
	if err != nil {
		fmt.Println("error 2")
		return orders
	}

	for rows.Next() {
		var o order.Order
		err = rows.Scan(&o)
		if err != nil {
			fmt.Println("error parsing")
			return orders
		}

		orders = append(orders, o)
	}

	return orders
}

func (r *DbRepository) Store(o order.Order) {
	q := `
		insert into orders (order_uid, data) values ($1, $2);
	`

	_, err := r.client.Exec(context.Background(), q, o.OrderUID, o)
	if err != nil {
		return
	}
}



