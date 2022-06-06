package order

import (
	"github.com/swooosh13/L0/inetrnal/models/order"
	"sync"
)

type CacheRepository struct {
	mx     *sync.RWMutex
	orders map[string]order.Order
}

func NewCache() *CacheRepository {
	return &CacheRepository{
		orders: make(map[string]order.Order),
		mx:     &sync.RWMutex{},
	}
}

func (r *CacheRepository) Load(key string) (order.Order, bool) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	v, ok := r.orders[key]
	return v, ok
}

func (r *CacheRepository) LoadAll() []order.Order {
	r.mx.RLock()
	defer r.mx.RUnlock()

	orders := make([]order.Order, 0)
	for _, v := range r.orders {
		orders = append(orders, v)
	}

	return orders
}
func (r *CacheRepository) Store(o order.Order) {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.orders[o.OrderUID] = o
}

func (r *CacheRepository) Recover(repoClient *DbRepository) {
	orders := repoClient.LoadAll()
	for _, v := range orders {
		r.Store(v)
	}
}
