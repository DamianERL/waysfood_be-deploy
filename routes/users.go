package routes

import (
	"waysfood/handlers"
	"waysfood/pkg/middleware"
	"waysfood/pkg/mysql"
	"waysfood/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router){
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users",h.FindUsers).Methods("Get")
	r.HandleFunc("/user/{id}", middleware.Auth(h.GetUser)).Methods("Get")
	r.HandleFunc("/patners",middleware.Auth(h.FindPartners)).Methods("GET")
	r.HandleFunc("/user",middleware.Auth(middleware.UploadFile(h.UpdateUser))).Methods("PATCH")
	//
	r.HandleFunc("/user",h.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}",h.DeleteUser).Methods("DELETE")
}