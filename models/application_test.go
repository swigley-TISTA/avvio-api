package models
import (
	"github.com/tidwall/buntdb"
	"log"
	"testing"
)

func TestApplication(t *testing.T) {
	var testApplication Application
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)

	err = testApplication.CreateApplication("AppName","ProjectName", "A really useful project")

	expected := Application{"AppName", "ProjectName", "A really useful project"}
	err = testApplication.GetApplication("AppName")
	if err == nil {
		if expected.Name == testApplication.Name && expected.Description == testApplication.Description {
			t.Log("test passed.")
		} else {
			t.Errorf("Response should be %s, got %s", expected, testApplication)
		}

	} else {
		t.Error("Error getting task from DB.")
	}
}

func TestForEmptyName(t *testing.T) {

	var testApplication Application
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)
	err = testApplication.CreateApplication("","ProjectName", "A really useful project")

		if err == nil {

			t.Errorf("Should have errored because name is empty.")
		} else {
		t.Log("test passed.")
	}

}


func TestForInvalidAppSearch(t *testing.T) {

	var testApplication Application
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)
	err = testApplication.GetApplication("")

	if err == nil {

		t.Errorf("Should have errored because name is invalid.")
	} else {
		t.Log("test passed.")
	}

}
