package models

import (
	"encoding/json"
	"errors"
)

type TeamMember struct {
	Name     string `json:"name"`
	Description string `json:"description"`
}

var teamMemberIdLabel = "team_member"

func (teamMember *TeamMember) GetTeamMember(name string)  error {
	if name == "" {
		return errors.New("Name must be set")
	}
	teamId := getObjLabel(teamIdLabel, name)

	objStr, err := currDb.GetValue(teamId)

	json.Unmarshal([]byte(objStr), &teamMember)
	return err
}

func (teamMember *TeamMember) CreateTeamMember(name string, description string) error {
	if name == "" {
		return errors.New("Name must be set")
	}
	teamId := getObjLabel(teamIdLabel, name)
	jsonEnc, err := json.Marshal(Team{name,description})
	if err == nil {
		currDb.CreateValue(teamId, string(jsonEnc))
	}
	return err
}

