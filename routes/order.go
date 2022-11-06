package routes

import (
	"waysfood/handlers"
	"waysfood/pkg/middleware"
	"waysfood/pkg/mysql"
	"waysfood/repositories"

	"github.com/gorilla/mux"
)

func OrderRoutes(r *mux.Router) {
	orderRepository := repositories.RepositoryOrder(mysql.DB)
	h:= handlers.HandlerOrder(orderRepository)

	r.HandleFunc("/orders",h.FindOrder).Methods("GET")
	r.HandleFunc("/order/{id}",h.GetOrder).Methods("GET")
	r.HandleFunc("/order",middleware.Auth(h.CreateOrder)).Methods("POST")
	r.HandleFunc("/order/{id}",h.UpdateOrder).Methods("PATCH")
	r.HandleFunc("/order/{id}",middleware.Auth(h.DeleteOrder)).Methods("DELETE")

}