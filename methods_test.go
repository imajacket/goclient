package goclient

import (
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	type TestStruct struct {
		UserId    int64  `json:"user_id,omitempty"`
		Id        int64  `json:"id,omitempty"`
		Title     string `json:"title,omitempty"`
		Completed bool   `json:"completed,omitempty"`
	}

	response, err := GoClient(TestStruct{}).Url("https://jsonplaceholder.typicode.com/todos/1").Get()
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Log(response)
}

func TestGetArray(t *testing.T) {
	type TestStruct struct {
		UserId    int64  `json:"user_id,omitempty"`
		Id        int64  `json:"id,omitempty"`
		Title     string `json:"title,omitempty"`
		Completed bool   `json:"completed,omitempty"`
	}

	response, err := GoClient([]TestStruct{}).Url("https://jsonplaceholder.typicode.com/posts").ExpectedResponseCode(http.StatusOK).Get()
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Log(response)
}

func TestPost(t *testing.T) {
	type TestStruct struct {
		UserId int64  `json:"user_id,omitempty"`
		Id     int64  `json:"id,omitempty"`
		Title  string `json:"title,omitempty"`
		Body   string `json:"body,omitempty"`
	}

	response, err := GoClient(TestStruct{}).Url("https://jsonplaceholder.typicode.com/posts").Post(TestStruct{
		Title:  "Testing",
		Body:   "Testing json",
		UserId: 1,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Log(response)
}

func TestReuse(t *testing.T) {
	type TestStruct struct {
		UserId int64  `json:"user_id,omitempty"`
		Id     int64  `json:"id,omitempty"`
		Title  string `json:"title,omitempty"`
		Body   string `json:"body,omitempty"`
	}

	client := GoClient(TestStruct{}).Url("https://jsonplaceholder.typicode.com/posts")

	response, err := client.Post(TestStruct{
		Title:  "Testing",
		Body:   "Testing json",
		UserId: 1,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	response2, err := client.Post(TestStruct{
		Title:  "Testing 2",
		Body:   "Testing json 2",
		UserId: 2,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Log(response, response2)
}
