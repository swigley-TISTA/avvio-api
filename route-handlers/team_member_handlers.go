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
