// blockchaineventfetcher.go
package services

import (
	"github.com/Marketen/lido-index/internal/core/domain"
	"github.com/Marketen/lido-index/internal/core/ports"
)

type BlockchainEventFetcher struct {
    watcher ports.BlockchainWatcher
}

func NewBlockchainEventFetcher(watcher ports.BlockchainWatcher) *BlockchainEventFetcher {
    return &BlockchainEventFetcher{watcher: watcher}
}

func (bef *BlockchainEventFetcher) FetchAndProcessEvents() ([]domain.Event, error) {
    return bef.watcher.FetchEvents()
}