package main

import (
	"fmt"

	"github.com/dhruvyadav14/go-stuff/logging" // Replace with the actual package path of your logging module
)

func main() {
	// Use NewDevelopmentLogger for development environment
	logger, err := logging.NewDevelopmentLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer logger.Sync() // Flush any buffered logs on exit

	// Log messages at various levels
	logger.Debug("Starting the application")
	logger.Info("Processing request for user ID: 123")
	logger.Warn("Encountered a potential issue: Database connection timed out")
	logger.Error(fmt.Errorf("An error occurred while processing the request"))
	logger.Fatal("Critical error, application shutting down")
}
