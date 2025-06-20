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
	OrderServiceURL = "http://order-service:8084"
)

type RequestParams struct {
	Method      string
	URL         string
	Path        string
	QueryParams url.Values
}

func SendRequestToService(method string, address string, path string, queryParams url.Values) (*http.Response, error) {
	request := createRequest(RequestParams{
		Method:      method,
		URL:         address,
		Path:        path,
		QueryParams: queryParams,
	})
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Couldn't get response from service.")
		return nil, err
	}
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
