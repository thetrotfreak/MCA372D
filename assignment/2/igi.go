package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type OperationalStatus int

const (
	OP_ABORTED OperationalStatus = iota
	OP_LIVE
	OP_ACCOMPLISHED
)

type Mission struct {
	Name        string
	Objective   []string
	Status      OperationalStatus
	Codeword    string
	Confidential bool
}

var missionMap = make(map[string]Mission)

// Sensitive codeword that triggers a panic
const sensitiveCodeword = "ultraSecret"

var currentFile *os.File

func main() {
	for {
		fmt.Println("\nMission Control:")
		fmt.Println("1. Save Mission")
		fmt.Println("2. Load Mission")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Proceed with: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			saveMission()
		case 2:
			loadMission()
		case 3:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}

func saveMission() {
	defer cleanup()

	// Take input from the user for mission details
	var mission Mission
	fmt.Print("Enter Mission Name: ")
	fmt.Scan(&mission.Name)
	fmt.Print("Enter Mission Objectives (comma-separated): ")
	fmt.Scan(&mission.Objective)
	fmt.Print("Enter Mission Status (0 for Aborted, 1 for Live, 2 for Accomplished): ")
	fmt.Scan(&mission.Status)
	fmt.Print("Enter Mission Codeword: ")
	fmt.Scan(&mission.Codeword)
	fmt.Print("Is Mission Confidential? (true/false): ")
	fmt.Scan(&mission.Confidential)

	// Save mission data to the map
	missionMap[mission.Codeword] = mission

	// Save mission data to a file
	err := saveMissionsToFile("missions_data.json", missionMap)
	if err != nil {
		handleError(err)
		return
	}

	fmt.Println("Mission saved successfully!")
}

func loadMission() {
	defer cleanup()

	// Prompt the user for the codeword of the mission to access
	var codeword string
	fmt.Print("Enter Mission Codeword to Access: ")
	fmt.Scan(&codeword)

	// Check if the specified codeword triggers a panic
	if mission, ok := missionMap[codeword]; ok && mission.Confidential && codeword == sensitiveCodeword {
		panic("Unauthorized access to confidential mission!")
	}

	// Load missions data from the file
	loadedMission, err := loadMissionFromFile("missions_data.json", codeword)
	if err != nil {
		handleError(err)
		return
	}

	// Accessing loaded mission data
	fmt.Println("Loaded Mission Data:", loadedMission)
}

// Save missions data to a file
func saveMissionsToFile(filename string, missions map[string]Mission) error {
	// Convert missions data to JSON
	missionsJSON, err := json.Marshal(missions)
	if err != nil {
		return err
	}

	// Open the file for writing
	currentFile, err = os.Create(filename)
	if err != nil {
		return err
	}

	// Write the missions data to the file
	_, err = currentFile.Write(missionsJSON)
	if err != nil {
		return err
	}

	fmt.Println("Missions data saved to file:", filename)
	return nil
}

// Load missions data from a file
func loadMissionFromFile(filename string, codeword string) (Mission, error) {
	// Read the missions data from the file
	missionsJSON, err := ioutil.ReadFile(filename)
	if err != nil {
		return Mission{}, err
	}

	// Unmarshal the JSON data into map of missions
	var missions map[string]Mission
	err = json.Unmarshal(missionsJSON, &missions)
	if err != nil {
		return Mission{}, err
	}

	// Retrieve the mission with the specified codeword
	loadedMission, found := missions[codeword]
	if !found {
		return Mission{}, fmt.Errorf("Mission with codeword %s not found", codeword)
	}

	fmt.Println("Missions data loaded from file:", filename)
	return loadedMission, nil
}

// Cleanup resources using defer
func cleanup() {
	fmt.Println("Cleaning up resources...")
	if currentFile != nil {
		currentFile.Close()
		currentFile = nil
	}
}

// Handle errors gracefully
func handleError(err error) {
	fmt.Println("Error:", err)
}
