package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// VoteData represents a single vote cast by a voter
type VoteData struct {
	VoterID   string `json:"voter_id"`
	Candidate string `json:"candidate"`
	TimeStamp int64  `json:"timestamp"`
}

func main() {
	// Create a buffered channel for vote data
	voteChan := make(chan VoteData, 100)

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start a goroutine to send votes to the server
	wg.Add(1)
	go func() {
		defer wg.Done()
		sendVotesToServer(voteChan)
	}()

	// Simple menu for casting votes
	var voterID, candidate string
	for voterID != "exit" {
		fmt.Print("Enter Voter ID (or 'exit' to poll): ")
		fmt.Scanln(&voterID)
		if voterID == "exit" {
			break
		}
		fmt.Print("Enter Candidate: ")
		fmt.Scanln(&candidate)

		// Send the vote data to the channel
		vd := VoteData{VoterID: voterID, Candidate: candidate, TimeStamp: time.Now().Unix()}
		voteChan <- vd
	}

	// Close the channel to signal workers to stop
	close(voteChan)

	// Wait for all goroutines to finish
	wg.Wait()

	// Get the winner from the server
	winner, err := getWinner()
	if err != nil {
		fmt.Println("Error getting winner:", err)
		return
	}
	fmt.Println("Winner:", winner)
}

func sendVotesToServer(voteChan <-chan VoteData) {
	for vd := range voteChan {
		jsonData, err := json.Marshal(vd)
		if err != nil {
			fmt.Println("Error marshalling vote data:", err)
			continue
		}

		_, err = http.Post("http://localhost:8080/vote", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error sending vote to server:", err)
			continue
		}
	}
}

func getWinner() (string, error) {
	resp, err := http.Get("http://localhost:8080/winner")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var winner string
	err = json.NewDecoder(resp.Body).Decode(&winner)
	if err != nil {
		return "", err
	}

	return winner, nil
}
