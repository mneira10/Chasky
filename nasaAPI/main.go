package nasaapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const mainURL = "https://api.nasa.gov"
const apodEndpoint = "/planetary/apod?api_key=DEMO_KEY"

type APODResponse struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
	Explanation    string `json:"explanation"`
}

// GetAPODURL retrieves NASA's astronomy picture of the day's URL
func GetAPODURL() string {

	completeURL := mainURL + apodEndpoint

	req, err := http.NewRequest("GET", completeURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	apodResponse := APODResponse{}

	jsonErr := json.Unmarshal(body, &apodResponse)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return apodResponse.URL

}
