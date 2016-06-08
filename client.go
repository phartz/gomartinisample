package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getResponseFromURL(url string) ([]byte, error) {
	fmt.Printf("Start getting bytes from [%s].\n", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error code received [%d/%s].", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyData, nil
}

func unmarshalVersionFromBytes(bytes []byte) (*Version, error) {
	var version Version
	err := json.Unmarshal(bytes, &version)
	if err != nil {
		return nil, err
	}

	return &version, nil
}

func main() {
	urls := [2]string{
		"http://localhost:3000/info",
		"http://localhost:3000/renderer",
	}

	for _, url := range urls {
		bytesReceived, err := getResponseFromURL(url)
		if err != nil {
			fmt.Printf("Cann't receive bytes from url [%s].\nError: %s", url, err)
		}

		version, err := unmarshalVersionFromBytes(bytesReceived)
		if err != nil {
			fmt.Printf("Cann't unmarshal bytes from url [%s].\nError: %s", url, err)
		}

		fmt.Printf("Version is [%d].\n", version.ID)
	}
}
