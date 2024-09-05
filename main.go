package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func csvToJson() {
	// Assume the user is always passing in a CSV file as the first argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: csvToJson input.csv")
		return
	}
	csvFile := os.Args[1]

	// Open CSV file
	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	// Read CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV data:", err)
		return
	}

	// Convert CSV to JSON
	if len(records) == 0 {
		fmt.Println("No records found in CSV file")
		return
	}

	headers := records[0]
	var result []map[string]string

	for _, record := range records[1:] {
		row := make(map[string]string)
		for i, header := range headers {
			row[header] = record[i]
		}
		result = append(result, row)
	}

	// Write JSON data to stdout
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON data:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func main() {
	csvToJson()
}
