package models

import (
	"encoding/json"
	"errors"
)

type Project struct {
	Name     string `json:"name"`
	Description string `json:"description"`
	Teams []string `json:teams`
}

var projectIdLabel = "project"

var teams = []string{}
var applications = []string{}

func (project *Project) GetProject(name string)  error {
	if name == "" {
		return errors.New("Name must be set")
	}
	projectId := getObjLabel(projectIdLabel, name)

	objStr, err := currDb.GetValue(projectId)

	json.Unmarshal([]byte(objStr), &project)
	return err
}

func (project *Project) CreateProject(name string, description string) error {
	if name == "" {
		return errors.New("Name must be set")
	}
	projectId := getObjLabel(projectIdLabel, name)
	jsonEnc, err := json.Marshal(Project{name,description,[]string{}})
	if err == nil {
		currDb.CreateValue(projectId, string(jsonEnc))
	}
	return nil
}

func (project *Project) AddTeam(team string) error {
	existing_team := Team{}
	existing_team.GetTeam(team)
	if existing_team.Name != "" {
		teams = append(teams, team)
	}
	return nil
}

func (project *Project) AddApplication(application string) error {
	existingApplication := Application{}
	existingApplication.GetApplication(application)
	if existingApplication.Name != "" {
		teams = append(teams, application)
	}
	return nil
}

func (project *Project) InitApplications() {

}
