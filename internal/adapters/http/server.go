package http

import (
	"log"
	"net/http"

	"github.com/Marketen/lido-index/internal/core/ports"
	"github.com/gorilla/mux"
)

// Server holds the dependencies for the HTTP server
type Server struct {
    Storer  ports.EventStore
}

// NewServer creates a new HTTP server with necessary dependencies
func NewServer(storer ports.EventStore) *Server {
    return &Server{
        Storer:  storer,
    }
}

// Start initializes the HTTP server and routes
func (s *Server) Start(address string) error {
    r := mux.NewRouter()
    r.HandleFunc("/events", s.handleGetEvents).Methods(http.MethodGet)
    http.Handle("/", r)

    log.Printf("Starting server on %s", address)
    return http.ListenAndServe(address, nil)
}

func (s *Server) handleGetEvents(w http.ResponseWriter, r *http.Request) {
    GetEventsHandler(s.Storer, w, r)
}
