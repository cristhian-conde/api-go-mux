package api

import (
	"awesomeProject/service/user"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *Server {
	return &Server{addr, db}
}

func (s *Server) Run() error {

	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)
	log.Println("api server listening on " + s.addr)
	return http.ListenAndServe(s.addr, router)
}
