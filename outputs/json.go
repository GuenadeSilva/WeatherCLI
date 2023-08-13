package outputs

import (
	"encoding/json"
	"fmt"
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
