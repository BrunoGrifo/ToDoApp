package api

import (
	"database/sql"
	"log"
	"net/http"
	"todo/service/task"

	"github.com/gorilla/csrf"
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
	// API handlers
	var mux *http.ServeMux = http.NewServeMux()
	var taskRepository *task.Repository = task.NewRepository(s.db)
	task.NewHandler(taskRepository).RegisterRoutes(mux)

	// File server
	fileserver := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	log.Println("Listening on", s.addr)
	csrfMiddleware := csrf.Protect(
		[]byte("32-byte-long-auth-key"),
		csrf.Path("/task"),
		csrf.Secure(false), // Disable this for local development without HTTPS
	)
	return http.ListenAndServe(s.addr, csrfMiddleware(mux))
	// return http.ListenAndServe(s.addr, mux)
}
