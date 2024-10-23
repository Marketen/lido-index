// blockchainwatcher.go
package ports

import "github.com/Marketen/lido-index/internal/core/domain"

type BlockchainWatcher interface {
    FetchEvents() ([]domain.Event, error)
}
