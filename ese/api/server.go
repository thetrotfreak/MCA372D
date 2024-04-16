package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

// VoteData represents a single vote cast by a voter
type VoteData struct {
	VoterID   string `json:"voter_id"`
	Candidate string `json:"candidate"`
	TimeStamp int64  `json:"timestamp"`
}

// VoteTally keeps track of the vote count for each candidate
type VoteTally struct {
	// Mutex is required to prevent concurrent modifications
	sync.Mutex
	Votes map[string]int
}

var voteTally VoteTally
var voteChan chan VoteData

const maxVotes = 100

func init() {
	// initalize appropriates
	voteTally = VoteTally{
		Votes: make(map[string]int),
	}
	voteChan = make(chan VoteData, maxVotes)

	go handleVotes()
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Election is now LIVE 🔴")
}

func voteHandler(w http.ResponseWriter, r *http.Request) {
	var vd VoteData
	err := json.NewDecoder(r.Body).Decode(&vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send the vote data to the channel
	select {
	case voteChan <- vd:
		fmt.Fprintf(w, "Vote received for %s\n", vd.Candidate)
	default:
		http.Error(w, "Maximum vote count reached", http.StatusTooManyRequests)
	}
}

func winnerHandler(w http.ResponseWriter, r *http.Request) {
	winner := getWinner()
	jsonData, err := json.Marshal(winner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func handleVotes() {
	for vd := range voteChan {
		voteTally.ProcessVote(vd)
	}
	close(voteChan)
}

func (vt *VoteTally) ProcessVote(vd VoteData) {
	vt.Lock()
	defer vt.Unlock()
	vt.Votes[vd.Candidate]++
}

func getWinner() string {
	voteTally.Lock()
	defer voteTally.Unlock()
	maxVotes := 0
	winners := make([]string, 0)
	for candidate, count := range voteTally.Votes {
		if count > maxVotes {
			maxVotes = count
			winners = []string{candidate}
		} else if count == maxVotes {
			winners = append(winners, candidate)
		}
	}

	if len(winners) == 1 {
		return winners[0]
	} else {
		return "Tie between: " + strings.Join(winners, ", ")
	}
}

func StartServer() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/vote", voteHandler)
	http.HandleFunc("/winner", winnerHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func main() {
	StartServer()
}
