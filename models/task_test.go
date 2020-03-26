package models

import (
"github.com/tidwall/buntdb"
"log"
"testing"
)

func TestTask(t *testing.T) {
	var testTask Task
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)

	testTask.CreateTask("foo","bar")

	expected := Task{"foo", "bar"}
	err = testTask.GetTask("foo")
	if err == nil {
		if expected.Name == testTask.Name && expected.Description == testTask.Description {

		} else {
			t.Errorf("Response should be %s, got %s", expected, testTask)
		}

	} else {
		t.Error("Error getting task from DB.")
	}

}


func TestForEmptyTaskName(t *testing.T) {

	var testProject Project
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)
	err = testProject.CreateProject("","ApplicationName")

	if err == nil {

		t.Errorf("Should have errored because name is empty.")
	} else {
		t.Log("test passed.")
	}

}


func TestForInvalidTaskSearch(t *testing.T) {

	var testTask Task
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)
	err = testTask.GetTask("")

	if err == nil {

		t.Errorf("Should have errored because name is invalid.")
	} else {
		t.Log("test passed.")
	}

}
