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

func TestGetPipelinesHandler(t *testing.T) {

	pipelines = []models.Pipeline{
		{"cook", "Make something to eat", ""},
	}

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(GetPipelineHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Pipeline{"cook", "Make something to eat", ""}
	b := []models.Pipeline{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual.Name != expected.Name || actual.Description != actual.Description {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestGetPipelinesByIdHandler(t *testing.T) {

	pipelines = []models.Pipeline{
		{"get groceries", "Buy some groceries", ""},
		{"cook", "Make something to eat", ""},
	}

	req, err := http.NewRequest("GET", "/pipeline/cook", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := mux.NewRouter()
	hf.HandleFunc("/pipeline/{id}", GetPipelineHandler)

	// hf := http.HandlerFunc(GetPipelineHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Pipeline{"cook", "Make something to eat",""}
	b := []models.Pipeline{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual.Name != expected.Name || actual.Description != actual.Description {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestCreatePipelinesHandler(t *testing.T) {

	pipelines = []models.Pipeline{
		{"get groceries", "Buy some groceries", ""},
	}

	form := newCreatePipelineForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(CreatePipelineHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Pipeline{"get groceries", "Buy some groceries", ""}

	if err != nil {
		t.Fatal(err)
	}

	actual := pipelines[1]

	if actual.Name != expected.Name || actual.Description != actual.Description {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func newCreatePipelineForm() *url.Values {
	form := url.Values{}
	form.Set("name", "get groceries")
	form.Set("description", "Buy some groceries")
	form.Set("teams", "[]")
	return &form
}
