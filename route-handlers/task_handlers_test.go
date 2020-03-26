package route_handlers

import (
	"avvio-api/models"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetTasksHandler(t *testing.T) {

	tasks = []models.Task{
		{"cook", "Make something to eat"},
	}

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(GetTaskHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Task{"cook", "Make something to eat"}
	b := []models.Task{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestGetTasksByIdHandler(t *testing.T) {

	tasks = []models.Task{
		{"get groceries", "Buy some groceries"},
		{"cook", "Make something to eat"},
	}

	req, err := http.NewRequest("GET", "/task/cook", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := mux.NewRouter()
	hf.HandleFunc("/task/{id}", GetTaskHandler)

	// hf := http.HandlerFunc(GetTaskHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Task{"cook", "Make something to eat"}
	b := []models.Task{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestCreateTasksHandler(t *testing.T) {

	tasks = []models.Task{
		{"get groceries", "Buy some groceries"},
	}

	form := newCreateTaskForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(CreateTaskHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Task{"get groceries", "Buy some groceries"}

	if err != nil {
		t.Fatal(err)
	}

	actual := tasks[1]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func newCreateTaskForm() *url.Values {
	form := url.Values{}
	form.Set("name", "get groceries")
	form.Set("description", "Buy some groceries")
	return &form
}
