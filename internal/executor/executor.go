package executor

import (
	"fmt"
	"sql-db/internal/parser"
	"sql-db/internal/storage"
)

func Execute(cmd parser.Command, db *storage.Database) (string, error) {
	switch cmd.Action {
	case "CREATE":
		db.Tables[cmd.TableName] = storage.Table{Name: cmd.TableName}
		return "table created: " + cmd.TableName, nil

	case "SELECT":
		if _, ok := db.Tables[cmd.TableName]; !ok {
			return "", fmt.Errorf("table not found: %s", cmd.TableName)
		}
		return "table exists: " + cmd.TableName, nil

	default:
		return "", fmt.Errorf("unknown action: %s", cmd.Action)
	}
}
