package api

import (
	"database/sql"
	"log"
	"net/http"
	"trackify/service/payments"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewServer(addr string, db *sql.DB) APIServer {
	return APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	paymentsHandler := payments.NewHandler()
	paymentsHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
