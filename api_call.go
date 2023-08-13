package api_call

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	req, err := http.NewRequest("GET", "https://www.ncdc.noaa.gov/cdo-web/api/v2/datasets", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	tokenID := os.Getenv("CDO_TOKEN")
	// Replace "Your-Token" with your actual token.
	req.Header.Set("token", tokenID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
