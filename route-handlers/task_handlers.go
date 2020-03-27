package route_handlers

import (
	"avvio-api/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var tasks []models.Task


// GetTaskHandler godoc
// @Summary Create tasks
// @Description get tasks
// @Accept  json
// @Produce  json
// @Param name path string false "search by name"
// @Success 200 {array} models.Task
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /task [get]
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
		 if strings.HasPrefix(t.Name, namePrefix) {
			 outTask = append(outTask, t)
		 }
	 }
	return outTask

}

// CreateTaskHandler godoc
// @Summary Create tasks
// @Description create tasks
// @Accept  json
// @Produce  json
// @Param task body models.Task false "task"
// @Success 200
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /task [post]
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
