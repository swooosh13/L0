package main

import (
	"context"
	"fmt"
	"github.com/swooosh13/L0/inetrnal/composites"
	"github.com/swooosh13/L0/inetrnal/config"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nats-io/stan.go"
)

func main() {
	cfg, _ := config.GetConfig()
	log.Println("load config")

	pgCfg := cfg.PostgresDB
	pgComposite, err := composites.NewPostgresDBComposite(context.Background(), pgCfg.Host, pgCfg.Port, pgCfg.Username, pgCfg.Password, pgCfg.Database, pgCfg.Timeout, pgCfg.MaxConns)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("pg connected - ok")

	orderComposite, err := composites.NewOrderComposite(pgComposite)
	if err != nil {
		log.Fatal(err)
	}

	sc, _ := stan.Connect(cfg.Stan.ClusterId, cfg.Stan.ClientId)
	sc.Subscribe("pub-1", func(msg *stan.Msg) {
		// TODO
		// 1. unmarshal to struct
		// 2. validate struct
		// 3. Store to repo & cache
		fmt.Println(string(msg.Data))
	})

	r := chi.NewRouter()
	orderComposite.Handler.Register(r)

	http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Listen.Host, cfg.Listen.Port), r)
}
