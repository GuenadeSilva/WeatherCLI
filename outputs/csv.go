package outputs

import (
	"encoding/csv"
	"fmt"
	"os"
	"weathercli/api"
)

func ConvertResultsToCSV(datasets []api.Dataset, destination string) error {
	if destination == "" {
		destination = "weather.csv"
	}

	// Create a new file to store CSV data
	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Write the header of the CSV file and the successive rows by iterating through the JSON struct array
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"UID", "MindDate", "MaxDate", "Name", "Datacoverage", "ID"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range datasets {
		csvRow := []string{r.UID, r.MindDate, r.MaxDate, r.Name, fmt.Sprint(r.Datacoverage), r.Id}
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}
