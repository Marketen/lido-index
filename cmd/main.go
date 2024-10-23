package main

import (
	"net/http"

	"github.com/Marketen/lido-index/internal/adapters/api"
	"github.com/Marketen/lido-index/internal/adapters/storage"
	"github.com/Marketen/lido-index/internal/core/services"
	"github.com/Marketen/lido-index/internal/infrastructure/blockchain"
)

func main() {
    // Initialize infrastructure. These are adapters.
    ethereumWatcher := blockchain.NewEthereumWatcher() // Assume constructor for Ethereum watcher
    jsonRepo := storage.NewJSONRepository("path_to_json_storage.json")

    // Initialize core services.
    eventFetcher := services.NewBlockchainEventFetcher(ethereumWatcher)
    eventService := services.NewEventService(jsonRepo)

    // Fetch events from the blockchain
    events, err := eventFetcher.FetchAndProcessEvents()
    if err != nil {
        panic(err)
    }

    // Process and store events
    err = eventService.ProcessAndStoreEvents(events)
    if err != nil {
        panic(err)
    }

    // Setup and start the HTTP server
    router := api.SetupRouter(eventService)
    http.ListenAndServe(":8080", router) // Run HTTP server
}
