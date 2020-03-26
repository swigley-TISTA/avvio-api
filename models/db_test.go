package models

import (
	"fmt"
	"github.com/tidwall/buntdb"
	"log"
	"testing"
)

func TestStore(t *testing.T) {
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)
	var s = &dbStore{db: db}
	s.CreateValue("foo", "bar")

	returnVal, err := s.GetValue("foo")
	if err != nil {
		fmt.Print(err)
		t.Error("Error getting value from DB.")
	}

	expected := "bar"
	if returnVal != expected {
		t.Errorf("Response should be %s, got %s", expected, returnVal)
	}

}
