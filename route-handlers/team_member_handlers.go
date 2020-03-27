package route_handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"avvio-api/models"
)

var teamMembers []models.TeamMember

// GetTeamMemberHandler godoc
// @Summary Create teamMembers
// @Description get teamMembers
// @Accept  json
// @Produce  json
// @Param name path string false "search by name"
// @Success 200 {array} models.TeamMember
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /teamMember [get]
func GetTeamMemberHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	teamMemberListBytes, err := json.Marshal(teamMembers)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if id == "" {

		w.Write(teamMemberListBytes)
	} else {
		outTeamMember := filterTeamMembersByNamePrefix(teamMembers, id)
		teamMemberBytes, err := json.Marshal(outTeamMember)
		if err == nil {
			w.Write(teamMemberBytes)
		}
	}
}

func filterTeamMembersByNamePrefix(teamMembers []models.TeamMember, namePrefix string) []models.TeamMember {

	outTeamMember := []models.TeamMember{}
	for _, t := range teamMembers {
		if(strings.HasPrefix(t.Name, namePrefix)) {
			outTeamMember = append(outTeamMember, t)
		}
	}
	return outTeamMember

}

// CreateTeamMemberHandler godoc
// @Summary Create teamMember
// @Description create teamMember
// @Accept  json
// @Produce  json
// @Param teamMember body models.TeamMember false "teamMember"
// @Success 200
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /teamMember [post]
func CreateTeamMemberHandler(w http.ResponseWriter, r *http.Request) {

	teamMember := models.TeamMember{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the teamMember from the form info
	teamMember.Name = r.Form.Get("name")
	teamMember.Description = r.Form.Get("description")

	// Append our existing list of birds with a new entry
	teamMembers = append(teamMembers, teamMember)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)

}
