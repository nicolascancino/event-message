package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler interface {
	PublishMessageHandler(w http.ResponseWriter, r *http.Request)
}
type Router struct {
	handler Handler
	mux     *mux.Router
}

func NewRouter(handler Handler, mux *mux.Router) *Router {
	return &Router{handler: handler, mux: mux}
}

func (r Router) SetupRouter() *mux.Router {
	r.mux.HandleFunc("/publish", r.handler.PublishMessageHandler).Methods(http.MethodPost)
	return r.mux
}
