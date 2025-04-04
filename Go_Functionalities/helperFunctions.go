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

	fmt.Println("Enter the name of the team:")
	var team string
	fmt.Scanln(&team) // Read the team name from user input

	countryMatches := getCountryMatches(team)

	fmt.Printf("Win percentage of %s: %.2f%%", team,getWinPercentage(team, countryMatches)) // Print the win percentage of Egypt
}

// This function will read the csv file and filter the matches related to Egypt 
func getCountryMatches(team string) [][]string {
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
	teamMatchDetails := [][]string{} // Initialize a slice to store match details
	// Iterate through the records and filter matches related to team
	for _, record := range records { 
		if record[1] == team || record[2] == team {
			teamMatchDetails = append(teamMatchDetails, record) // Append the record to the slice
		}
	}
	
	teamMatchesCount := len(teamMatchDetails) // Get the number of matches related to team

	fmt.Printf("Found %d matches related to %s\n", teamMatchesCount, team)

	return teamMatchDetails // Return the match details slice

}

// Function to get the win percentage of a team
func getWinPercentage(team string, matches [][]string) float64 {
	// This function will calculate the win percentage of a team based on the matches data
	winCount,drawCount,lossCount := 0,0,0 
	totalMatches := len(matches) // Get the total number of matches

	for _, match := range matches { // Iterate through the matches
		// fmt.Println("Match details: ", match[0], " ", match[1], " ", match[2], " ",match[3], " ",match[4]) // Print the match details
		if (match[1] == team && match[3] > match[4]) || (match[2] == team && match[3] < match[4]) { // Check if the team won
			winCount++ 
		} else if (match[1] == team && match[3] == match[4]) || (match[2] == team && match[3] == match[4]) { // Check if the match was a draw
			drawCount++ 
		} else if (match[1] == team && match[3] < match[4]) || (match[2] == team && match[3] > match[4]) { // Check if the team lost
			lossCount++ 
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
