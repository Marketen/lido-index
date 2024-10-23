package http

import (
	"encoding/json"
	"net/http"

	"github.com/Marketen/lido-index/internal/core/ports"
)

// GetEventsHandler handles the GET request for fetching events
func GetEventsHandler(storer ports.EventStore, w http.ResponseWriter, r *http.Request) {
    events, err := storer.LoadEvents()
    if err != nil {
        http.Error(w, "Failed to retrieve events", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(events)
}
