package models
import (
	"github.com/tidwall/buntdb"
	"log"
	"testing"
)

func TestPipeline(t *testing.T) {
	var testPipeline Pipeline
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)

	testPipeline.CreatePipeline("PipelineName","ApplicationName", "A really reliable pipeline")

	expected := Pipeline{"PipelineName", "ApplicationName", "A really reliable pipeline"}
	err = testPipeline.GetPipeline("PipelineName")
	if err == nil {
		if expected.Name == testPipeline.Name && expected.Description == testPipeline.Description {

		} else {
			t.Errorf("Response should be %s, got %s", expected, testPipeline)
		}

	} else {
		t.Error("Error getting task from DB.")
	}

}

func TestForEmptyPipelineName(t *testing.T) {

	var testPipeline Pipeline
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)
	err = testPipeline.CreatePipeline("","ApplicationName", "A really reliable pipeline")

	if err == nil {

		t.Errorf("Should have errored because name is empty.")
	} else {
		t.Log("test passed.")
	}

}


func TestForInvalidPipelineSearch(t *testing.T) {

	var testPipeline Pipeline
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitStore(db)
	err = testPipeline.GetPipeline("")

	if err == nil {

		t.Errorf("Should have errored because name is invalid.")
	} else {
		t.Log("test passed.")
	}

}
