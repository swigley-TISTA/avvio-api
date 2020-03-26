package models
import (
	"github.com/tidwall/buntdb"
	"log"
	"testing"
)

func TestTeam(t *testing.T) {
	var testTeam Team
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)

	testTeam.CreateTeam("The A Team","A really fun team")

	expected := Team{"The A Team", "A really fun team"}
	err = testTeam.GetTeam("The A Team")
	if err == nil {
		if expected.Name == testTeam.Name && expected.Description == testTeam.Description {

		} else {
			t.Errorf("Response should be %s, got %s", expected, testTeam)
		}

	} else {
		t.Error("Error getting task from DB.")
	}

}



func TestForEmptyTeamName(t *testing.T) {

	var testTeam Team
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)
	err = testTeam.CreateTeam("","A really fun team")

	if err == nil {

		t.Errorf("Should have errored because name is empty.")
	} else {
		t.Log("test passed.")
	}

}


func TestForInvalidTeamSearch(t *testing.T) {

	var testTeam Team
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)
	err = testTeam.GetTeam("")

	if err == nil {

		t.Errorf("Should have errored because name is invalid.")
	} else {
		t.Log("test passed.")
	}

}
