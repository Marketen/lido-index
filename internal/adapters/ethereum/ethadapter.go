package ethereum

import (
	"github.com/Marketen/lido-index/internal/core/domain"
)

type EthAdapter struct {
    // Ethereum client details. 
	// Add ethereum rpc api URL, etc.
}

func (e *EthAdapter) FetchEvents() ([]domain.Event, error) {
    // Logic to subscribe & fetch events from Ethereum
    return []domain.Event{}, nil
}
