package storage

type Row map[string]string

type Table struct {
	Name string
	Rows []Row
}

type Database struct {
	Tables map[string]Table
}

func NewDatabase() *Database {
	return &Database{
		Tables: make(map[string]Table),
	}
}
