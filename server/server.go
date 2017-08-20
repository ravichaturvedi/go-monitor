package server

import (
	"net/http"
	"log"
	"fmt"
)


// Server is interface to expose the registry to the outside world.
type Server interface {
	Serve() error
}

func New(h http.Handler) Server {
	return httpServer{h}
}


type httpServer struct {
	h http.Handler
}


func (s httpServer) Serve() error {
	// Register the handler to respond to the queries.
	http.Handle("/", s.h)

	log.Println(fmt.Sprintf("Starting server: http://0.0.0.0:1234"))

	// Start the server to listen on all the interfaces on port `1234`
	return http.ListenAndServe(":1234", nil)
}