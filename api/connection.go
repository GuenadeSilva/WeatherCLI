package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Dataset struct {
	UID          string  `json:"uid"`
	MindDate     string  `json:"mindate"`
	MaxDate      string  `json:"maxdate"`
	Name         string  `json:"name"`
	Datacoverage float32 `json:"datacoverage"`
	Id           string  `json:"id"`
}

type Datasets struct {
	Results []Dataset `json:"results"`
}

// ApiCall is an exported function that makes an API call
func ApiCall() []Dataset {
	req, err := http.NewRequest("GET", "https://www.ncdc.noaa.gov/cdo-web/api/v2/datasets", nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file:", errEnv)
	}

	tokenID := os.Getenv("CDO_TOKEN")
	if tokenID == "" {
		log.Fatal("CDO_TOKEN not found in environment")
		return nil
	}

	req.Header.Set("token", tokenID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var datasets Datasets
	err = json.Unmarshal(body, &datasets)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return datasets.Results
}
