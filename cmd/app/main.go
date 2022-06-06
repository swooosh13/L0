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

const (
	clusterId = "microservice"
	clientId  = "sub-1"
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

	storage := composites.NewStorage(pgComposite)
	log.Println("storage has been created")
	os := storage.Cache.LoadAll()
	fmt.Println(len(os))
	// TODO
	// stan
	sc, _ := stan.Connect(clusterId, clientId)

	sc.Subscribe("receive-data", func(msg *stan.Msg) {

		fmt.Println(msg.Data)
	})

	r := chi.NewRouter()

	http.ListenAndServe(":8080", r)
}

func RegisterOrder(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {

		})
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		})
	})
}
