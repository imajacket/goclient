package goclient

import "net/http"

type Client[T any] struct {
	client       *http.Client
	headers      http.Header
	url          string
	expectedCode int
	model        T
	rawResponse  []byte
}

// GoClient - initialize Client
func GoClient[T any](responseModel T) *Client[T] {
	client := Client[T]{
		client: &http.Client{},
		model:  responseModel,
	}

	return &client
}

// Headers - set headers for Client
func (c *Client[T]) Headers(headers http.Header) *Client[T] {
	c.headers = headers
	return c
}

// Url - set url for Client
func (c *Client[T]) Url(url string) *Client[T] {
	c.url = url
	return c
}

// ExpectedResponseCode - (optional) add an expected HTTP status code to Client
func (c *Client[T]) ExpectedResponseCode(code int) *Client[T] {
	c.expectedCode = code
	return c
}

// GetRawResponse - Get the raw response from the last API call made with the Client
func (c *Client[T]) GetRawResponse() []byte {
	return c.rawResponse
}
