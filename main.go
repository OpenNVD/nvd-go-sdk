package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	data, err := GetCVEs(20, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Print the decoded data
	fmt.Println("Results Per Page:", data.ResultsPerPage)
	fmt.Println("Total Results:", data.TotalResults)
}

func GetCVEs(resultsPerPage int, startIndex int) (*VulnerabilitiesResponse, error) {
	url := fmt.Sprintf("https://services.nvd.nist.gov/rest/json/cves/2.0/?resultsPerPage=%d&startIndex=%d", resultsPerPage, startIndex)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return nil, err
	}
	defer response.Body.Close()

	var data VulnerabilitiesResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	return &data, nil
}
