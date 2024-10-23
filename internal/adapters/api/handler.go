// handler.go
package api

import (
	"encoding/json"
	"net/http"

	"github.com/Marketen/lido-index/internal/core/ports"
)

// SetupRouter initializes the HTTP routes and handlers.
func SetupRouter(eventService ports.EventRepository) *http.ServeMux {
    router := http.NewServeMux()
    router.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            events, err := eventService.RetrieveAll() // Assuming a method to retrieve all events
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            json.NewEncoder(w).Encode(events)
        } else {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        }
    })

    return router
}
