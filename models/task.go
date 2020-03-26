package models

import (
	"encoding/json"
	"errors"
)

type Task struct {
	Name     string `json:"name"`
	Description string `json:"description"`
}

var taskIdLabel = "task"

func (task *Task) GetTask(name string)  error {
	taskId := getObjLabel(taskIdLabel, name)

	objStr, err := currDb.GetValue(taskId)

	json.Unmarshal([]byte(objStr), &task)
	return  err

}

func (task *Task) CreateTask(name string, description string) error {
	if name == "" {
		return errors.New("Name must be set")
	}
	taskId := getObjLabel(taskIdLabel, name)
	jsonEnc, err := json.Marshal(Task{name,description})
	if err == nil {
		currDb.CreateValue(taskId, string(jsonEnc))
	}
	return err
}
