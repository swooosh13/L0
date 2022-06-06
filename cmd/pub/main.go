package main

import (
	"github.com/nats-io/stan.go"
)

const (
	clusterId = "microservice"
	clientId  = "pub-1"
)

func main() {
	sc, _ := stan.Connect(clusterId, clientId)

	sc.Publish("receive-data", []byte("message1"))
}
