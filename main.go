package main

import (
	"log"
	"net/http"
)

func main() {
	url := "https://api.github.com/repos/m-lukas/ibm_cloudfoundry/contents"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("API Request Creation failed!")
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("API call failed")
		return
	}

	defer resp.Body.Close()
	log.Println(resp.Body)

}
