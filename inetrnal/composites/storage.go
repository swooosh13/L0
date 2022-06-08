package composites

import "github.com/swooosh13/L0/inetrnal/repository/order"

type Storage struct {
	Cache *order.CacheRepository
	Repo  *order.DbRepository
}

func NewStorage(pgComposite *PostgresDBComposite) *Storage {
	c := order.NewCache()
	d := order.NewDbRepository(pgComposite.Client)
	c.Recover(d)

	return &Storage{
		Cache: c,
		Repo:  d,
	}
}
