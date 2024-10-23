package ports

import "github.com/Marketen/lido-index/internal/core/domain"

// Fetch all events
type EventFetcher interface {
    FetchEvents() ([]domain.Event, error)
}

// Given an event, saves it on db
type EventStore interface {
    SaveEvents(events []domain.Event) error
    LoadEvents() ([]domain.Event, error)
}
