package storage

type Table struct {
	Name string
}

type Database struct {
	Tables map[string]Table
}

func NewDatabase() *Database {
	return &Database{
		Tables: make(map[string]Table),
	}
}