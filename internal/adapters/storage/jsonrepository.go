package storage

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Marketen/lido-index/internal/core/domain"
)

type JsonStorage struct {
    FilePath string
}

func (js *JsonStorage) SaveEvents(events []domain.Event) error {
    fileData, err := json.Marshal(events)
    if err != nil {
        return err
    }
    return ioutil.WriteFile(js.FilePath, fileData, 0644)
}

// Load Events
func (js *JsonStorage) LoadEvents() ([]domain.Event, error) {
    fileData, err := ioutil.ReadFile(js.FilePath)
    if err != nil {
        return nil, err
    }
    var events []domain.Event
    err = json.Unmarshal(fileData, &events)
    if err != nil {
        return nil, err
    }
    return events, nil
}