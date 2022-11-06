package routes

import (
	"waysfood/handlers"
	"waysfood/pkg/middleware"
	"waysfood/repositories"
	"waysfood/pkg/mysql"


	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router){
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h:= handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions",middleware.Auth(h.FindTransactions)).Methods("GET")
	r.HandleFunc("/incomes",middleware.Auth(h.FindIncomes)).Methods("GET")
	r.HandleFunc("/transaction",middleware.Auth(h.CreateTransaction)).Methods("POST")
	//midtrans
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}