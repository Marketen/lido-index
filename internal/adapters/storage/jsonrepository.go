// jsonrepository.go
package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Marketen/lido-index/internal/core/domain"
	"github.com/Marketen/lido-index/internal/core/ports"
)

type JSONRepository struct {
    filePath string
}

func NewJSONRepository(filePath string) ports.EventRepository {
    return &JSONRepository{filePath: filePath}
}

func (jr *JSONRepository) Save(event domain.Event) error {
    events, err := jr.load()
    if err != nil {
        return err
    }

    events = append(events, event)
    return jr.save(events)
}

func (jr *JSONRepository) Retrieve(id string) (domain.Event, error) {
    events, err := jr.load()
    if err != nil {
        return domain.Event{}, err
    }

    for _, event := range events {
        if event.ID == id {
            return event, nil
        }
    }
    return domain.Event{}, os.ErrNotExist
}

func (jr *JSONRepository) RetrieveAll() ([]domain.Event, error) {
    return jr.load()
}

func (jr *JSONRepository) load() ([]domain.Event, error) {
    var events []domain.Event
    data, err := ioutil.ReadFile(jr.filePath)
    if err != nil {
        if os.IsNotExist(err) {
            return []domain.Event{}, nil
        }
        return nil, err
    }
    err = json.Unmarshal(data, &events)
    return events, err
}

func (jr *JSONRepository) save(events []domain.Event) error {
    data, err := json.Marshal(events)
    if err != nil {
        return err
    }
    return ioutil.WriteFile(jr.filePath, data, 0644)
}
