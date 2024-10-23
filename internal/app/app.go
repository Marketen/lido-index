package app

import (
	"github.com/Marketen/lido-index/internal/adapters/ethereum"
	"github.com/Marketen/lido-index/internal/adapters/storage"
	"github.com/Marketen/lido-index/internal/core/services"
)

func InitializeApp() {

    // Initialize the Ethereum adapter and JSON storage
    ethAdapter := &ethereum.EthAdapter{} // Needed to fetch events from Ethereum
    jsonStorage := &storage.JsonStorage{FilePath: "path/to/file.json"} // Needed to store events

    // With both adapters plugged in, the event service can fetch and store events.
    // This is possible because NewEventService expects interfaces that match the EventFetcher and EventStore interfaces.
    eventService := services.NewEventService(ethAdapter, jsonStorage)

    // Start the event processing in a separate goroutine or via a scheduled job. 
    // The event service could have another method to fire notifications or perform other actions based on the events.
    eventService.ProcessEvents() 
}
