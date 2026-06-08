package storage

import (
	"encoding/json"
	"os"
)

type Row map[string]string

type Table struct {
	Name    string
	Columns []string
	Rows    []Row
}

type Database struct {
	Tables map[string]Table
}

func NewDatabase() *Database {
	return &Database{
		Tables: make(map[string]Table),
	}
}

const dbFile = "db.json"

func (db *Database) Save() error {
	data, err := json.MarshalIndent(db, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(dbFile, data, 0644)
}

func LoadDatabase() (*Database, error) {
	data, err := os.ReadFile(dbFile)
	if err != nil {
		return NewDatabase(), nil
	}

	var db Database

	err = json.Unmarshal(data, &db)
	if err != nil {
		return nil, err
	}

	if db.Tables == nil {
		db.Tables = make(map[string]Table)
	}

	return &db, nil
}
