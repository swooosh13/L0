package main

import (
	"encoding/json"
	"github.com/jaswdr/faker"
	"github.com/nats-io/stan.go"
	"github.com/swooosh13/L0/inetrnal/models/order"
	"math/rand"
	"time"
)

const (
	clusterId = "microservice"
	clientId  = "pub-1"
)

var fake faker.Faker

func init() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	src := rand.NewSource(seed)
	fake = faker.NewWithSeed(src)
}

func main() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	sc, _ := stan.Connect(clusterId, clientId)

	ticker := time.NewTicker(3 * time.Second)
	errorFormatTicker := time.NewTicker(5 * time.Second)

	errorFormates := []string{"{}", "asdads", "{Price: 200}"}

	for {
		select {
		case <-ticker.C:
			data, _ := json.Marshal(GenerateFakeOrder())

			sc.Publish("pub-1", data)
			time.Sleep(3 * time.Second)
		case <-errorFormatTicker.C:
			data := []byte(errorFormates[rand.Intn(3)])
			sc.Publish("pub-1", data)
		}
	}
}

func GenerateFakeOrder() *order.Order {
	trackNumber := fake.Gamer().Tag()
	n := rand.Intn(10) + 1
	itms := make([]order.Item, n)
	for i := 0; i < n; i++ {
		itms[i] = order.Item{
			ChrtId:      rand.Intn(10000) + 1000,
			TrackNumber: trackNumber,
			Price:       rand.Intn(9000) + 1000,
			Rid:         fake.UUID().V4(),
			Name:        fake.Company().Suffix(),
			Sale:        rand.Intn(100) + 1,
			Size:        "0",
			TotalPrice:  rand.Intn(100000) + 1000,
			NmId:        rand.Intn(200000) + 100000,
			Brand:       fake.Company().Name(),
			Status:      202,
		}
	}

	uid := fake.UUID().V4()

	o := &order.Order{
		OrderUID:    uid,
		TrackNumber: trackNumber,
		Entry:       "WBIL",
		Delivery: order.Delivery{
			Name:    fake.Person().Name(),
			Phone:   fake.Phone().Number(),
			Zip:     fake.Address().PostCode(),
			City:    fake.Address().City(),
			Address: fake.Address().Address(),
			Region:  fake.Address().CityPrefix(),
			Email:   fake.Internet().Email(),
		},
		Payment: order.Payment{
			Transaction:  uid,
			RequestId:    "",
			Currency:     fake.Currency().Currency(),
			Provider:     fake.Payment().CreditCardType(),
			Amount:       rand.Intn(10000) + 100,
			PaymentDt:    rand.Intn(1637907727) + 100,
			Bank:         fake.Company().Name(),
			DeliveryCost: rand.Intn(1500) + 200,
			GoodsTotal:   rand.Intn(1000) - 100,
			CustomFee:    rand.Intn(1000),
		},
		Items:             itms,
		Locale:            "ru",
		InternalSignature: "",
		CustomerId:        "test",
		DeliveryService:   "meest",
		ShardKey:          "9",
		SmId:              rand.Intn(300) + 1,
		DateCreated:       randate().String(),
		OofShard:          "1",
	}

	return o
}

func randate() time.Time {
	min := time.Date(2015, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2022, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
