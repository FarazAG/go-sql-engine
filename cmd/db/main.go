package main

import (
	"sql-db/internal/repl"
	"sql-db/internal/storage"
)

func main() {
	db, err := storage.LoadDatabase()
	if err != nil {
		panic(err)
	}

	repl.Start(db)
}
