package outputs

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"weathercli/api" // Adjust this import path according to your project
)

// ConvertResultsToJSONFile converts datasets.Results to a JSON file
func ConvertResultsToJSONFile(datasets []api.Dataset, filename string) error {
	if filename == "" {
		filename = "weather.json"
	}

	jsonData, err := json.MarshalIndent(datasets, "", "    ")
	if err != nil {
		return err
	}

	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = outputFile.Write(jsonData)
	return err
}

// PrintJSONFile prints the contents of a JSON file
func PrintJSONFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	prettyJSON, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(prettyJSON))
	return nil
}

func JSONReader() {
	// Open the JSON file
	file, err := os.Open("datasets.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the JSON data from the file
	dataBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Create a struct to hold the JSON data
	var jsonData []api.Dataset

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(dataBytes, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println(jsonData)

}
