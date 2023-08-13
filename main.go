package main

import (
	"flag"
	"fmt"
	"weathercli/api"
	"weathercli/outputs"
)

func main() {
	writeToCSV := flag.Bool("csv", true, "Enable CSV writing")
	writeToJSON := flag.Bool("json", false, "Enable JSON writing")
	flag.Parse()

	if flag.NArg() == 0 {
		// No additional arguments provided, print contents of datasets.json
		if err := outputs.PrintJSONFile("datasets.json"); err != nil {
			fmt.Println("Error printing JSON:", err)
		}
		return
	}

	// Scrape the data
	dataset := api.ApiCall()

	if *writeToJSON {
		if err := outputs.ConvertResultsToJSONFile(dataset, "datasets.json"); err != nil {
			fmt.Println("Error writing to JSON:", err)
		} else {
			fmt.Println("Dataset results converted and saved to datasets.json")
		}
	}

	if *writeToCSV {
		if err := outputs.ConvertResultsToCSV(dataset, "datasets.csv"); err != nil {
			fmt.Println("Error writing to CSV:", err)
		} else {
			fmt.Println("Dataset results converted and saved to datasets.csv")
		}
	}
}
