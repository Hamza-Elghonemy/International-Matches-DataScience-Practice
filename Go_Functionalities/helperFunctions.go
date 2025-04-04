package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Read csv file
	file, err := os.Open("../Data/results.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 
	}
	defer file.Close() // Ensure the file is closed after reading
	fmt.Println("File opened successfully")

	egyptMatches := getEgyptMatches()

	fmt.Printf("Win percentage of Egypt: %.2f%%", getWinPercentage("Egypt", egyptMatches)) // Print the win percentage of Egypt
}

// This function will read the csv file and filter the matches related to Egypt 
func getEgyptMatches() [][]string {
	file, err := os.Open("../Data/results.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return [][]string{}
	}
	defer file.Close() 
	reader := csv.NewReader(file) // Create a new CSV reader
	// Read all records from the CSV file
	records, err := reader.ReadAll() 
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return [][]string{}
	}
	egyptMatchDetails := [][]string{} // Initialize a slice to store match details
	// Iterate through the records and filter matches related to Egypt
	for _, record := range records { 
		if record[1] == "Egypt" || record[2] == "Egypt" {
			egyptMatchDetails = append(egyptMatchDetails, record) // Append the record to the slice
		}
	}
	
	egyptMatchesCount := len(egyptMatchDetails) // Get the number of matches related to Egypt

	fmt.Printf("Found %d matches related to Egypt\n", egyptMatchesCount)

	return egyptMatchDetails // Return the match details slice

}

// Function to get the win percentage of a team
func getWinPercentage(team string, matches [][]string) float64 {
	// This function will calculate the win percentage of a team based on the matches data
	winCount := 0 // Initialize win count
	drawCount := 0 // Initialize draw count
	lossCount := 0 // Initialize loss count
	totalMatches := len(matches) // Get the total number of matches

	for _, match := range matches { // Iterate through the matches
		// fmt.Println("Match details: ", match[0], " ", match[1], " ", match[2], " ",match[3], " ",match[4]) // Print the match details
		if (match[1] == team && match[3] > match[4]) || (match[2] == team && match[3] < match[4]) { // Check if the team won
			winCount++ // Increment win count
		} else if (match[1] == team && match[3] == match[4]) || (match[2] == team && match[3] == match[4]) { // Check if the match was a draw
			drawCount++ // Increment draw count
		} else if (match[1] == team && match[3] < match[4]) || (match[2] == team && match[3] > match[4]) { // Check if the team lost
			lossCount++ // Increment loss count
		}

	}
	if totalMatches == 0 { // Check if there are no matches
		return 0.0 // Return 0.0 to avoid division by zero
	}
	fmt.Printf("Total matches won: %d\n", winCount) // Print total matches and wins
	fmt.Printf("Total matches drawed: %d\n", drawCount) // Print total matches and draws
	fmt.Printf("Total matches lost: %d\n", lossCount) // Print total matches and losses
	winPercentage := float64(winCount) / float64(totalMatches) * 100.0 // Calculate win percentage

	return winPercentage // Return the win percentage to 2 decimal places
}
