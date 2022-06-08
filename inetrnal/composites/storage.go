package composites

import (
	"github.com/swooosh13/L0/inetrnal/repository/order"
	"github.com/swooosh13/L0/pkg/logger"
)

type Storage struct {
	Cache *order.CacheRepository
	Repo  *order.DbRepository
}

func NewStorage(pgComposite *PostgresDBComposite) *Storage {
	c := order.NewCache()
	d := order.NewDbRepository(pgComposite.Client)
	c.Recover(d)
	logger.Info("Data has been recovered from db to in-memory cache")

	return &Storage{
		Cache: c,
		Repo:  d,
	}
}
