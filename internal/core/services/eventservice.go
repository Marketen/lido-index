package services

import (
	"log"

	"github.com/Marketen/lido-index/internal/core/ports"
)

// EventService implements the business logic for event handling.
type EventService struct {
    fetcher ports.EventFetcher  // Port for fetching events
    storer  ports.EventStore    // Port for storing events
}

// NewEventService creates a new instance of EventService with necessary dependencies.
// Whatever is passed as fetcher and storer must implement the EventFetcher and EventStore interfaces respectively.
func NewEventService(fetcher ports.EventFetcher, storer ports.EventStore) *EventService {
    return &EventService{
        fetcher: fetcher,
        storer:  storer,
    }
}

// ProcessEvents handles the fetching and storing of events.
func (es *EventService) ProcessEvents() error {
    events, err := es.fetcher.FetchEvents()
    if err != nil {
        log.Printf("Error fetching events: %v", err)
        return err
    }

    if len(events) == 0 {
        log.Println("No events found to process.")
        return nil
    }

    err = es.storer.SaveEvents(events)
    if err != nil {
        log.Printf("Error storing events: %v", err)
        return err
    }

    log.Printf("Successfully processed %d events.", len(events))
    return nil
}