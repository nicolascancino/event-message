package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Router interface {
	SetupRouter() *mux.Router
}

type Server struct {
	router Router
}

func NewServer(router Router) *Server {
	return &Server{router: router}
}

func (s Server) Start() {
	routerHandler := s.router.SetupRouter()

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), routerHandler); err != nil {

	}

}
