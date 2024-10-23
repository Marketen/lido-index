// eventservice.go
package services

import (
	"github.com/Marketen/lido-index/internal/core/domain"
	"github.com/Marketen/lido-index/internal/core/ports"
)

type EventService struct {
    repo ports.EventRepository
}

func NewEventService(repo ports.EventRepository) *EventService {
    return &EventService{repo: repo}
}

func (es *EventService) ProcessAndStoreEvents(events []domain.Event) error {
    for _, event := range events {
        // Add processing logic if needed
        err := es.repo.Save(event)
        if err != nil {
            return err
        }
    }
    return nil
}