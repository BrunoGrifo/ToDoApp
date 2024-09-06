package api

import (
	"database/sql"
	"log"
	"net/http"
	"todo/service/task"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) APIServer {
	return APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	var mux *http.ServeMux = http.NewServeMux()
	task.NewHandler().RegisterRoutes(mux)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, mux)
}
