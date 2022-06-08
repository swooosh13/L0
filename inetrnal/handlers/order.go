package handlers

import (
	"github.com/swooosh13/L0/inetrnal/repository/order"
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
)

var ordersTpl = template.Must(template.ParseFiles("web/index.html"))
var orderTpl = template.Must(template.ParseFiles("web/order.html"))

type OrderHandler struct {
	storage   *order.CacheRepository
	pageCache map[string]interface{}
}

func NewOrderHandler(s *order.CacheRepository) Handler {
	return &OrderHandler{
		storage:   s,
		pageCache: make(map[string]interface{}),
	}
}

func (o *OrderHandler) Register(r *chi.Mux) {
	r.Route("/order", func(r chi.Router) {
		r.Get("/", o.OrdersPage)
		r.Get("/{id}", o.OrderPage)
	})
}

func (o *OrderHandler) OrderPage(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")

	order, _ := o.storage.Load(uid)
	o.pageCache["order"] = order
	o.pageCache["title"] = "order " + order.OrderUID

	orderTpl.Execute(w, o.pageCache)
}

func (o *OrderHandler) OrdersPage(w http.ResponseWriter, r *http.Request) {
	orders := o.storage.LoadAll()

	o.pageCache["orders"] = orders
	o.pageCache["title"] = "This is title!"

	ordersTpl.Execute(w, o.pageCache)
}
