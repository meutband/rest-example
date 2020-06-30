package main

import (
	"fmt"
	"log"
	"net/http"
)

const listenPort = "8000"

type handler struct {
	router *http.ServeMux
}

func newHandler() *handler {
	h := &handler{router: http.DefaultServeMux}
	h.routes()
	return s
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func main() {

	// http server type
	server := http.Server{Addr: fmt.Sprintf("localhost:%s", listenPort)}

	// create custom handler type for http.server.handler
	hand := newHandler()
	server.Handler = hand

	log.Printf("Add server listening on port %s\n", listenPort)
	err := server.ListenAndServe()
	fmt.Println(err)
}
