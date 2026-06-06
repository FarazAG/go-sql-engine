package main

import (
	"sql-db/internal/repl"
	"sql-db/internal/storage"
)

func main() {
	db := storage.NewDatabase()
	repl.Start(db)
}
