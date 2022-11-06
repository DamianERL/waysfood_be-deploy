package routes

import (
	"waysfood/handlers"
	"waysfood/pkg/middleware"
	"waysfood/pkg/mysql"
	"waysfood/repositories"

	"github.com/gorilla/mux"
)


func CartRoutes(r *mux.Router)  {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h:= handlers.HandlerCart(cartRepository)
	
		r.HandleFunc("/carts",h.FindCart).Methods("GET")
	r.HandleFunc("/cart-id",h.GEtCart).Methods("GET")
	r.HandleFunc("/cart",middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/cartID",middleware.Auth(h.UpdateCart)).Methods("PATCH")
	r.HandleFunc("/cart/{id}",middleware.Auth(h.DeleteCart)).Methods("DELETE")
	r.HandleFunc("/cart-status", middleware.Auth(h.FindbyIDCart)).Methods("GET")

}