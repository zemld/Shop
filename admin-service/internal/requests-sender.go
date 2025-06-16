package internal

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const (
	GET  = "GET"
	POST = "POST"
)

func SendRequestToItemService(method string, path string, queryParams url.Values) (*http.Response, error) {
	request, _ := http.NewRequest(method, fmt.Sprintf("http://item-service:8083%s", path), nil)
	if len(queryParams) > 0 {
		q := request.URL.Query()
		for key, values := range queryParams {
			for _, value := range values {
				q.Add(key, value)
			}
		}
		request.URL.RawQuery = q.Encode()
		log.Printf("Request URL with query parameters: %s\n", request.URL.String())
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Couldn't get response from user service.")
		return nil, err
	}
	log.Printf("Got response from user service: %s\n", response.Body)
	return response, nil
}
