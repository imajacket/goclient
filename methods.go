package goclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client[T]) Get() (T, error) {
	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		return c.model, err
	}

	req.Header = c.headers

	resp, err := c.client.Do(req)
	if err != nil {
		return c.model, err
	}

	if c.expectedCode != 0 {
		if resp.StatusCode != c.expectedCode {
			return c.model, errors.New("not expected response")
		}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.model, err
	}

	err = json.Unmarshal(body, &c.model)
	if err != nil {
		return c.model, err
	}

	return c.model, nil
}

func (c *Client[T]) Post(body any) (T, error) {
	var finalBody []byte

	switch body.(type) {
	case []byte:
		finalBody = body.([]byte)
	default:
		j, err := json.Marshal(body)
		if err != nil {
			finalBody = []byte(fmt.Sprintf("%v", body))
		}

		finalBody = j
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(finalBody))
	if err != nil {
		return c.model, err
	}

	req.Header = c.headers

	resp, err := c.client.Do(req)
	if err != nil {
		return c.model, err
	}

	if c.expectedCode != 0 {
		if resp.StatusCode != c.expectedCode {
			return c.model, errors.New("not expected response")
		}
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.model, err
	}

	err = json.Unmarshal(responseBody, &c.model)
	if err != nil {
		return c.model, err
	}

	return c.model, nil
}
