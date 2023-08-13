package outputs

import (
	"encoding/json"
	"os"
	"weathercli/api" // Adjust this import path according to your project
)

// ConvertResultsToJSONFile converts datasets.Results to a JSON file
func ConvertResultsToJSONFile(datasets []api.Dataset, filename string) error {
	if filename == "" {
		filename = "weather.csv"
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
