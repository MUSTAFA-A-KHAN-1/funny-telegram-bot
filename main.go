package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/MUSTAFA-A-KHAN/funny-telegram-bot/controller/joker"
)

func main() {
	botToken := "7484235929:AAHiFUWLo2nmqXyMe9gby7yc0SBUb8ZysE4"
	var wg sync.WaitGroup
	// Add two tasks to the WaitGroup
	wg.Add(2)
	go runJokerBot(botToken, &wg)
	go startHTTPServer()
	// Wait for goroutines to complete
	wg.Wait()

	log.Println("All bots stopped running.")

}

// Function to start a bot with error handling
func runJokerBot(botToken string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the wait group counter when this goroutine completes
	err := joker.StartBot(botToken)
	if err != nil {
		log.Panic("error starting bot: ", err)
	}

}

// startHTTPServer starts a simple HTTP server for health checks
func startHTTPServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bot is running!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
