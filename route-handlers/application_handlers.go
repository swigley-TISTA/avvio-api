package route_handlers

import (
	"avvio-api/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var applications []models.Application


func GetApplicationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	applicationListBytes, err := json.Marshal(applications)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if id == "" {

		w.Write(applicationListBytes)
	} else {
		outApplication := filterApplicationsByNamePrefix(applications, id)
		applicationBytes, err := json.Marshal(outApplication)
		if err == nil {
			w.Write(applicationBytes)
		}
	}
}

func filterApplicationsByNamePrefix(applications []models.Application, namePrefix string) []models.Application {

	outApplication := []models.Application{}
	for _, t := range applications {
		if(strings.HasPrefix(t.Name, namePrefix)) {
			outApplication = append(outApplication, t)
		}
	}
	return outApplication

}

func CreateApplicationHandler(w http.ResponseWriter, r *http.Request) {

	application := models.Application{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the application from the form info
	application.Name = r.Form.Get("name")
	application.Description = r.Form.Get("description")

	// Append our existing list of birds with a new entry
	applications = append(applications, application)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)

}

