package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendVotesToServer(t *testing.T) {
	// Create a channel for testing
	voteChan := make(chan VoteData, 2)
	defer close(voteChan)

	// Send some test vote data
	voteChan <- VoteData{VoterID: "2347111", Candidate: "Bivas"}
	voteChan <- VoteData{VoterID: "2347161", Candidate: "Rohit"}

	// Call the function under test
	go sendVotesToServer(voteChan)
}

func TestGetWinner(t *testing.T) {
	// Create a test server that returns a dummy winner
	expected := "anonymous"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DummyWinner"))
	}))
	defer server.Close()

	winner, err := getWinner()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if winner != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, winner)
	}
}
