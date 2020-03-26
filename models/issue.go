package models

import (
	"encoding/json"
	"errors"
)

type Issue struct {
	Name     string `json:"name"`
	ApplicationName string `json:applicationName`
	Description string `json:"description"`
}

var issueIdLabel = "issue"

func (issue *Issue) GetIssue(name string)  error {
	issueId := getObjLabel(issueIdLabel, name)

	objStr, err := currDb.GetValue(issueId)

	json.Unmarshal([]byte(objStr), &issue)
	return  err

}

func (issue *Issue) CreateIssue(name string, applicationName string, description string) error {
	var err error
	if name == "" {
		err = errors.New("Name must be set")
		return err
	}
	issueId := getObjLabel(issueIdLabel, name)
	jsonEnc, err := json.Marshal(Issue{name,applicationName,description})
	if err == nil {
		currDb.CreateValue(issueId, string(jsonEnc))
	}
	return err
}

