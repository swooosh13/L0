package composites

import "github.com/swooosh13/L0/inetrnal/repository/order"

type Storage struct {
	Cache *order.CacheRepository
	Repo  *PostgresDBComposite
}

func NewStorage(pgComposite *PostgresDBComposite) *Storage {
	c := order.NewCache()
	c.Recover(order.NewDbRepository(pgComposite.Client))

	return &Storage{
		Cache: c,
		Repo: pgComposite,
	}
}