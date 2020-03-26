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

func TestGetProjectsHandler(t *testing.T) {

	projects = []models.Project{
		{"cook", "Make something to eat", []string{}},
	}

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(GetProjectHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Project{"cook", "Make something to eat", []string{}}
	b := []models.Project{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual.Name != expected.Name || actual.Description != actual.Description {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestGetProjectsByIdHandler(t *testing.T) {

	projects = []models.Project{
		{"get groceries", "Buy some groceries", []string{}},
		{"cook", "Make something to eat", []string{}},
	}

	req, err := http.NewRequest("GET", "/project/cook", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := mux.NewRouter()
	hf.HandleFunc("/project/{id}", GetProjectHandler)

	// hf := http.HandlerFunc(GetProjectHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Project{"cook", "Make something to eat", []string{}}
	b := []models.Project{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual.Name != expected.Name || actual.Description != actual.Description {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestCreateProjectsHandler(t *testing.T) {

	projects = []models.Project{
		{"get groceries", "Buy some groceries", []string{}},
	}

	form := newCreateProjectForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(CreateProjectHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Project{"get groceries", "Buy some groceries", []string{}}

	if err != nil {
		t.Fatal(err)
	}

	actual := projects[1]

	if actual.Name != expected.Name || actual.Description != actual.Description {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func newCreateProjectForm() *url.Values {
	form := url.Values{}
	form.Set("name", "get groceries")
	form.Set("description", "Buy some groceries")
	form.Set("teams", "[]")
	return &form
}
