package models

import (
	"encoding/json"
	"errors"
)

type Pipeline struct {
	Name     string `json:"name"`
	ApplicationName string `json:applicationName`
	Description string `json:"description"`
}

var pipelineIdLabel = "pipeline"

func (pipeline *Pipeline) GetPipeline(name string)  error {
	pipelineId := getObjLabel(pipelineIdLabel, name)

	objStr, err := currDb.GetValue(pipelineId)


	json.Unmarshal([]byte(objStr), &pipeline)
	return  err

}

func (pipeline *Pipeline) CreatePipeline(name string, applicationName string, description string) error {
	if name == "" {
		return errors.New("Name must be set")
	}
	pipelineId := getObjLabel(pipelineIdLabel, name)
	jsonEnc, err := json.Marshal(Pipeline{name,applicationName,description})
	if err == nil {
		currDb.CreateValue(pipelineId, string(jsonEnc))
	}
	return nil
}

