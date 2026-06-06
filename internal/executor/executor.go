package executor

import (
	"fmt"
	"sql-db/internal/parser"
	"sql-db/internal/storage"
)

func Execute(cmd parser.Command, db *storage.Database) (string, error){
	switch cmd.Action{
	case "CREATE":
		db.Tables[cmd.Name] = storage.Table{Name: cmd.Name}
		return "table created: " + cmd.Name, nil
	
	case "SELECT":
		if _, ok := db.Tables[cmd.Name]; !ok{
			return "", fmt.Errorf("table not found: %s", cmd.Name)
		}
		return "table exists: " + cmd.Name, nil

	default:
		return "", fmt.Errorf("unknown action: %s", cmd.Action) 
	}
}