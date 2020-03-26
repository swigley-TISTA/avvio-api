package models
import (
	"github.com/tidwall/buntdb"
	"log"
	"testing"
)

func TestTeamMember(t *testing.T) {
	var testTeamMember TeamMember
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)

	testTeamMember.CreateTeamMember("Tom","A really great guy")

	expected := TeamMember{"Tom", "A really great guy"}
	err = testTeamMember.GetTeamMember("Tom")
	if err == nil {
		if expected.Name == testTeamMember.Name && expected.Description == testTeamMember.Description {

		} else {
			t.Errorf("Response should be %s, got %s", expected, testTeamMember)
		}

	} else {
		t.Error("Error getting task from DB.")
	}

}


func TestForEmptyTeamMemberName(t *testing.T) {

	var testTeamMember TeamMember
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)
	err = testTeamMember.CreateTeamMember("","A really great guy")

	if err == nil {

		t.Errorf("Should have errored because name is empty.")
	} else {
		t.Log("test passed.")
	}

}


func TestForInvalidTeamMemberSearch(t *testing.T) {

	var testTeamMember TeamMember
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)
	err = testTeamMember.GetTeamMember("")

	if err == nil {

		t.Errorf("Should have errored because name is invalid.")
	} else {
		t.Log("test passed.")
	}

}
