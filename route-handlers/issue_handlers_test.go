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

func TestGetIssuesHandler(t *testing.T) {

	issues = []models.Issue{
		{"cook", "Make something to eat", ""},
	}

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(GetIssueHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Issue{"cook", "Make something to eat", ""}
	b := []models.Issue{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual.Name != expected.Name || actual.Description != actual.Description {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestGetIssuesByIdHandler(t *testing.T) {

	issues = []models.Issue{
		{"get groceries", "Buy some groceries", ""},
		{"cook", "Make something to eat", ""},
	}

	req, err := http.NewRequest("GET", "/issue/cook", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := mux.NewRouter()
	hf.HandleFunc("/issue/{id}", GetIssueHandler)

	// hf := http.HandlerFunc(GetIssueHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Issue{"cook", "Make something to eat",""}
	b := []models.Issue{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual.Name != expected.Name || actual.Description != actual.Description {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestCreateIssuesHandler(t *testing.T) {

	issues = []models.Issue{
		{"get groceries", "Buy some groceries", ""},
	}

	form := newCreateIssueForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(CreateIssueHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Issue{"get groceries", "Buy some groceries", ""}

	if err != nil {
		t.Fatal(err)
	}

	actual := issues[1]

	if actual.Name != expected.Name || actual.Description != actual.Description {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func newCreateIssueForm() *url.Values {
	form := url.Values{}
	form.Set("name", "get groceries")
	form.Set("description", "Buy some groceries")
	form.Set("teams", "[]")
	return &form
}
