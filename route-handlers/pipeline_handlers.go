package route_handlers

import (
	"avvio-api/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var pipelines []models.Pipeline

// GetPipelineHandler godoc
// @Summary Create Pipeline
// @Description get Pipeline
// @Accept  json
// @Produce  json
// @Param name path string false "search by name"
// @Success 200 {array} models.Pipeline
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /pipeline [get]
func GetPipelineHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	pipelineListBytes, err := json.Marshal(pipelines)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if id == "" {

		w.Write(pipelineListBytes)
	} else {
		outPipeline := filterPipelinesByNamePrefix(pipelines, id)
		pipelineBytes, err := json.Marshal(outPipeline)
		if err == nil {
			w.Write(pipelineBytes)
		}
	}
}

func filterPipelinesByNamePrefix(pipelines []models.Pipeline, namePrefix string) []models.Pipeline {

	outPipeline := []models.Pipeline{}
	for _, t := range pipelines {
		if(strings.HasPrefix(t.Name, namePrefix)) {
			outPipeline = append(outPipeline, t)
		}
	}
	return outPipeline

}

// CreatePipelineHandler godoc
// @Summary Create pipelines
// @Description create pipelines
// @Accept  json
// @Produce  json
// @Param task body models.Pipeline false "pipeline"
// @Success 200
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /pipeline [post]
func CreatePipelineHandler(w http.ResponseWriter, r *http.Request) {

	pipeline := models.Pipeline{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the pipeline from the form info
	pipeline.Name = r.Form.Get("name")
	pipeline.Description = r.Form.Get("description")

	// Append our existing list of birds with a new entry
	pipelines = append(pipelines, pipeline)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)

}

