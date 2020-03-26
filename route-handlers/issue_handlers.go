package route_handlers

import (
	"avvio-api/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var issues []models.Issue


func GetIssueHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	issueListBytes, err := json.Marshal(issues)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if id == "" {

		w.Write(issueListBytes)
	} else {
		outIssue := filterIssuesByNamePrefix(issues, id)
		issueBytes, err := json.Marshal(outIssue)
		if err == nil {
			w.Write(issueBytes)
		}
	}
}

func filterIssuesByNamePrefix(issues []models.Issue, namePrefix string) []models.Issue {

	outIssue := []models.Issue{}
	for _, t := range issues {
		if(strings.HasPrefix(t.Name, namePrefix)) {
			outIssue = append(outIssue, t)
		}
	}
	return outIssue

}

func CreateIssueHandler(w http.ResponseWriter, r *http.Request) {

	issue := models.Issue{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the issue from the form info
	issue.Name = r.Form.Get("name")
	issue.Description = r.Form.Get("description")

	// Append our existing list of birds with a new entry
	issues = append(issues, issue)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)

}

