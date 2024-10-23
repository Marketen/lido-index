// eventrepository.go
package ports

import "github.com/Marketen/lido-index/internal/core/domain"

type EventRepository interface {
    Save(event domain.Event) error
    Retrieve(id string) (domain.Event, error)
    RetrieveAll() ([]domain.Event, error)
}

