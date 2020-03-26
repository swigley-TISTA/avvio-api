package models

import (
	"encoding/json"
	"errors"
)

type Application struct {
	Name     string `json:"name"`
	ProjectName string `json:projectName`
	Description string `json:"description"`
}

var applicationIdLabel = "application"

func (application *Application) GetApplication(name string)  error {
	applicationId := getObjLabel(applicationIdLabel, name)
	objStr, err := currDb.GetValue(applicationId)
	json.Unmarshal([]byte(objStr), &application)
	return  err

}

func (application *Application) CreateApplication(name string, projectName string, description string) error {
	if name == "" {
		return errors.New("Name must be set")
	}
	applicationId := getObjLabel(applicationIdLabel, name)
	jsonEnc, err := json.Marshal(Application{name,projectName, description})
	if err == nil {
		currDb.CreateValue(applicationId, string(jsonEnc))
	}
	return err
}
