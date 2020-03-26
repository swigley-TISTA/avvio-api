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

func TestGetTeamsHandler(t *testing.T) {

	teams = []models.Team{
		{"cook", "Make something to eat"},
	}

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(GetTeamHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Team{"cook", "Make something to eat"}
	b := []models.Team{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestGetTeamsByIdHandler(t *testing.T) {

	teams = []models.Team{
		{"get groceries", "Buy some groceries"},
		{"cook", "Make something to eat"},
	}

	req, err := http.NewRequest("GET", "/team/cook", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := mux.NewRouter()
	hf.HandleFunc("/team/{id}", GetTeamHandler)

	// hf := http.HandlerFunc(GetTeamHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Team{"cook", "Make something to eat"}
	b := []models.Team{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestCreateTeamsHandler(t *testing.T) {

	teams = []models.Team{
		{"get groceries", "Buy some groceries"},
	}

	form := newCreateTeamForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(CreateTeamHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := models.Team{"get groceries", "Buy some groceries"}

	if err != nil {
		t.Fatal(err)
	}

	actual := teams[1]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func newCreateTeamForm() *url.Values {
	form := url.Values{}
	form.Set("name", "get groceries")
	form.Set("description", "Buy some groceries")
	return &form
}
