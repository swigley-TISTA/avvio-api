package route_handlers

import (
	"avvio-api/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var projects []models.Project


func GetProjectHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	projectListBytes, err := json.Marshal(projects)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if id == "" {

		w.Write(projectListBytes)
	} else {
		outProject := filterProjectsByNamePrefix(projects, id)
		projectBytes, err := json.Marshal(outProject)
		if err == nil {
			w.Write(projectBytes)
		}
	}
}

func filterProjectsByNamePrefix(projects []models.Project, namePrefix string) []models.Project {

	outProject := []models.Project{}
	for _, t := range projects {
		if(strings.HasPrefix(t.Name, namePrefix)) {
			outProject = append(outProject, t)
		}
	}
	return outProject

}

func CreateProjectHandler(w http.ResponseWriter, r *http.Request) {

	project := models.Project{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the project from the form info
	project.Name = r.Form.Get("name")
	project.Description = r.Form.Get("description")

	// Append our existing list of birds with a new entry
	projects = append(projects, project)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)

}

