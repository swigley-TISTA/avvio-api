package models

import (
	"github.com/tidwall/buntdb"
	"log"
	"testing"
)

func TestIssue(t *testing.T) {
	var testIssue Issue
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)

	testIssue.CreateIssue("IssueName","AppName", "A really severe issue")

	expected := Issue{"IssueName", "AppName", "A really severe issue"}
	err = testIssue.GetIssue("IssueName")
	if err == nil {
		if expected.Name == testIssue.Name && expected.Description == testIssue.Description {

		} else {
			t.Errorf("Response should be %s, got %s", expected, testIssue)
		}

	} else {
		t.Error("Error getting Issue from DB.")
	}

}

func TestForEmptyIssueName(t *testing.T) {

	var testIssue Issue
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)
	err = testIssue.CreateIssue("","ApplicationName", "A really serious issue")

	if err == nil {

		t.Errorf("Should have errored because name is empty.")
	} else {
		t.Log("test passed.")
	}

}


func TestForInvalidIssueSearch(t *testing.T) {

	var testIssue Issue
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)
	err = testIssue.GetIssue("")

	if err != nil {

		t.Errorf("Should have errored because name is invalid.")
	} else {
		t.Log("test passed.")
	}

}
