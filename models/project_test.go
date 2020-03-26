package models
import (
	"github.com/tidwall/buntdb"
	"log"
	"testing"
)

func TestProject(t *testing.T) {
	var testProject Project
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)

	testProject.CreateProject("foo","bar")

	expected := Project{"foo", "bar", []string{}}
	err = testProject.GetProject("foo")
	if err == nil {
		if expected.Name == testProject.Name && expected.Description == testProject.Description {

		} else {
			t.Errorf("Response should be %s, got %s", expected, testProject)
		}

	} else {
		t.Error("Error getting task from DB.")
	}

}


func TestForEmptyProjectName(t *testing.T) {

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


func TestForInvalidProjectSearch(t *testing.T) {

	var testProject Project
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)
	err = testProject.GetProject("")

	if err == nil {

		t.Errorf("Should have errored because name is invalid.")
	} else {
		t.Log("test passed.")
	}

}
