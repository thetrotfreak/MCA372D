package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVoteHandler(t *testing.T) {
	// Create a test vote data
	vd := VoteData{VoterID: "2347111", Candidate: "Bivas Kumar"}
	jsonData, err := json.Marshal(vd)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test request with the vote data
	req, err := http.NewRequest("POST", "/vote", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler
	voteHandler(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, status)
	}
}

func TestGetWinner(t *testing.T) {
	// Create some test vote data
	vd1 := VoteData{VoterID: "voter1", Candidate: "Alice"}
	vd2 := VoteData{VoterID: "voter2", Candidate: "Bob"}
	vd3 := VoteData{VoterID: "voter3", Candidate: "Alice"}

	// Process the test vote data
	voteTally.ProcessVote(vd1)
	voteTally.ProcessVote(vd2)
	voteTally.ProcessVote(vd3)

	// Call the function under test
	winner := getWinner()

	// Check the result
	expected := "Alice"
	if winner != expected {
		t.Errorf("Expected winner '%s', but got '%s'", expected, winner)
	}
}
