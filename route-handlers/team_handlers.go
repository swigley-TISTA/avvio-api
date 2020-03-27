package route_handlers

import (
	"avvio-api/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var teams []models.Team


// GetTeamHandler godoc
// @Summary Create teams
// @Description get teams
// @Accept  json
// @Produce  json
// @Param name path string false "search by name"
// @Success 200 {array} models.Team
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /team [get]
func GetTeamHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	teamListBytes, err := json.Marshal(teams)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if id == "" {

		w.Write(teamListBytes)
	} else {
		outTeam := filterTeamsByNamePrefix(teams, id)
		teamBytes, err := json.Marshal(outTeam)
		if err == nil {
			w.Write(teamBytes)
		}
	}
}

func filterTeamsByNamePrefix(teams []models.Team, namePrefix string) []models.Team {

	outTeam := []models.Team{}
	for _, t := range teams {
		if(strings.HasPrefix(t.Name, namePrefix)) {
			outTeam = append(outTeam, t)
		}
	}
	return outTeam

}

// CreateTeamHandler godoc
// @Summary Create teams
// @Description create teams
// @Accept  json
// @Produce  json
// @Param team body models.Team false "team"
// @Success 200
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /team [post]
func CreateTeamHandler(w http.ResponseWriter, r *http.Request) {

	team := models.Team{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the team from the form info
	team.Name = r.Form.Get("name")
	team.Description = r.Form.Get("description")

	// Append our existing list of birds with a new entry
	teams = append(teams, team)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)

}
