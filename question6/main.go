package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// custom error struct

type cError struct {
	statusCode int
	exitCode   int
	details    string
}

// implementing Error() method to satisfy 'error' interface
func (c cError) Error() string {
	return fmt.Sprintf("status %d: err %s\nexit status %d", c.statusCode, c.details, c.exitCode)
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type mockHttpClient struct{}

func (m *mockHttpClient) Do(*http.Request) (*http.Response, error) {
	// return random error
	rand.Seed(time.Now().UnixNano())

	errors := []cError{
		{
			statusCode: 503,
			exitCode:   1,
			details:    "unavailable",
		},
		{
			statusCode: 404,
			exitCode:   1,
			details:    "not found",
		},
		{
			statusCode: 500,
			exitCode:   1,
			details:    "internal server error",
		},
	}
	return nil, errors[rand.Intn(len(errors))]
}

// get sends a post request to the URL
func get(url string, client HTTPClient) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	// if the client is mock it will return an custom error
	// it the client is not mock it will return the original success response
	return client.Do(request)
}

func main() {
	var client HTTPClient
	if len(os.Args) < 2 || os.Args[1] != "--mock=true" {
		// use original http client
		client = &http.Client{}
	} else {
		// use mock client which will return random errors
		client = &mockHttpClient{}
	}
	res, err := get("https://api.agify.io/?name=meelad", client)
	if err != nil {
		fmt.Println("Some error occured")
		fmt.Println(err)
	} else {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println("Success response", string(body))
		fmt.Println("If you want to see the custom error implementation run this code with \"go run main.go --mock=true\"")
	}
}
