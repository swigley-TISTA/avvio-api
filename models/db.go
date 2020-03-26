package models

import (
	"fmt"
    "github.com/tidwall/buntdb"
)

type Store interface {
	CreateValue(key string, value string) error
	GetValue(key string) (string, error)
}

var store Store
var currDb dbStore

type dbStore struct {
	db *buntdb.DB
}

func InitStore(s *buntdb.DB) {
	currDb.db = s
	currDb.db.CreateIndex("idx_name","*", buntdb.IndexJSON("name"))
}

func getObjLabel(objIdLabel string,id string) string {
	return objIdLabel + ":" + id
}

func (currDb *dbStore) CreateValue(key string, value string) (string, error) {
	err := currDb.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})
	return "", err
}

func (currDb *dbStore) GetValue(key string) (string,error)  {
	outvar := ""
	err := currDb.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil{
			return  err
		}
		outvar = val
		fmt.Printf("value is %s\n", val)
		return nil
	})
	return outvar, err
}

func (currDB *dbStore) GetAllValues(objType string) ([]string,error) {
	outvar := []string{}
	err := currDb.db.View(func(tx *buntdb.Tx) error {
		tx.AscendKeys(objType + ":*", func(k, v string) bool {
			if true == true {
				outvar = append(outvar, k)
			}
			return  true // continue
		})
		return nil

		})
	return outvar, err
}
