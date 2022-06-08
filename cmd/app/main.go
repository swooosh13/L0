package main

import (
	"context"
	"fmt"
	"github.com/swooosh13/L0/inetrnal/broker"
	"github.com/swooosh13/L0/inetrnal/composites"
	"github.com/swooosh13/L0/inetrnal/config"
	"github.com/swooosh13/L0/pkg/logger"
	"go.uber.org/zap"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nats-io/stan.go"
)

func main() {
	logger.Init()

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatal("Config loaded with error", zap.Error(err))
	}

	logger.Info("Config has been loaded")

	pgCfg := cfg.PostgresDB
	pgComposite, err := composites.NewPostgresDBComposite(context.Background(), pgCfg.Host, pgCfg.Port, pgCfg.Username, pgCfg.Password, pgCfg.Database, pgCfg.Timeout, pgCfg.MaxConns)
	if err != nil {
		logger.Fatal("Postgres composite created with error", zap.Error(err))
	}
	logger.Info("PostgresComposite has been created successfully")

	orderComposite, err := composites.NewOrderComposite(pgComposite)
	if err != nil {
		logger.Fatal("OrderComposite has been created with error", zap.Error(err))
	}
	logger.Info("OrderComposite has been created successfully")

	sc, err := stan.Connect(cfg.Stan.ClusterId, cfg.Stan.ClientId)
	if err != nil {
		logger.Fatal("Error connecting to nats-streaming", zap.Error(err))
	}
	logger.Info("Nats connection successfully established")

	sc.Subscribe("pub-1", func(msg *stan.Msg) {
		broker.ReceiveOrder(msg, orderComposite)
	})

	r := chi.NewRouter()
	orderComposite.Handler.Register(r)

	logger.Info("Server has been started")
	http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Listen.Host, cfg.Listen.Port), r)
}
