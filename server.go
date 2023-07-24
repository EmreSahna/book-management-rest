package main

import (
	"github.com/EmreSahna/go_mysql_book_management/middleware"
	"github.com/EmreSahna/go_mysql_book_management/routes"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Run() {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	routes.RegisterBookRoutes(r)
	routes.RegisterAuthorRoutes(r)
	routes.RegisterUserRoutes(r)
	http.ListenAndServe(s.listenAddr, r)
}
