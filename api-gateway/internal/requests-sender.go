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

const (
	UserServiceURL  = "http://user-service:8081"
	AdminServiceURL = "http://admin-service:8082"
)

type RequestParams struct {
	Method      string
	URL         string
	Path        string
	QueryParams url.Values
}

func SendRequestToUserService(method string, path string, queryParams url.Values) (*http.Response, error) {
	request := createRequest(RequestParams{
		Method:      method,
		URL:         UserServiceURL,
		Path:        path,
		QueryParams: queryParams,
	})
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Couldn't get response from user service.")
		return nil, err
	}
	log.Printf("Got response from user service: %s\n", response.Body)
	return response, nil
}

func SendRequestToAdminService(method string, path string, queryParams url.Values) (*http.Response, error) {
	request := createRequest(RequestParams{
		Method:      method,
		URL:         AdminServiceURL,
		Path:        path,
		QueryParams: queryParams})
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Couldn't get response from admin service.")
		return nil, err
	}
	log.Printf("Got response from admin service: %s\n", response.Body)
	return response, nil
}

func createRequest(r RequestParams) *http.Request {
	request, _ := http.NewRequest(r.Method, fmt.Sprintf("%s%s", r.URL, r.Path), nil)
	if len(r.QueryParams) > 0 {
		q := request.URL.Query()
		for key, values := range r.QueryParams {
			for _, value := range values {
				q.Add(key, value)
			}
		}
		request.URL.RawQuery = q.Encode()
		log.Printf("Request URL with query parameters: %s\n", request.URL.String())
	}
	return request
}
