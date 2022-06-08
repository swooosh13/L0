package composites

import "github.com/swooosh13/L0/inetrnal/handlers"

type OrderComposite struct {
	Storage *Storage
	Handler handlers.Handler
}

func NewOrderComposite(pgComposite *PostgresDBComposite) (*OrderComposite, error) {
	storage := NewStorage(pgComposite)
	handler := handlers.NewOrderHandler(storage.Cache)

	return &OrderComposite{
		Storage: storage,
		Handler: handler,
	}, nil
}
