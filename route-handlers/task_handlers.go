package route_handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"avvio-api/models"
)

var tasks []models.Task


func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	taskListBytes, err := json.Marshal(tasks)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if id == "" {

		w.Write(taskListBytes)
	} else {
		outTask := filterTasksByNamePrefix(tasks, id)
		taskBytes, err := json.Marshal(outTask)
		if err == nil {
			w.Write(taskBytes)
		}
	}
}

 func filterTasksByNamePrefix(tasks []models.Task, namePrefix string) []models.Task {

	 outTask := []models.Task{}
	 for _, t := range tasks {
		 if(strings.HasPrefix(t.Name, namePrefix)) {
			 outTask = append(outTask, t)
		 }
	 }
	return outTask

}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {

	task := models.Task{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the task from the form info
	task.Name = r.Form.Get("name")
	task.Description = r.Form.Get("description")

	// Append our existing list of birds with a new entry
	tasks = append(tasks, task)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)

}
