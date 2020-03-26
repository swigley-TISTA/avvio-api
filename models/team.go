package models

import (
	"encoding/json"
	"errors"
)

type Team struct {
	Name     string `json:"name"`
	Description string `json:"description"`
}

var teamIdLabel = "team"

func (team *Team) GetTeam(name string) error {
	if name == "" {
		return errors.New("Name must be set")
	}
	teamId := getObjLabel(teamIdLabel, name)

	objStr, err := currDb.GetValue(teamId)

	json.Unmarshal([]byte(objStr), &team)
	return err

}

func (team *Team) CreateTeam(name string, description string) error {
	if name == "" {
		return errors.New("Name must be set")
	}
	teamId := getObjLabel(teamIdLabel, name)
	jsonEnc, err := json.Marshal(Team{name,description})
	if err == nil {
		currDb.CreateValue(teamId, string(jsonEnc))
	}
	return nil
}

