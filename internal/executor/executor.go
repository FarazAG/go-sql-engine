package executor

import (
	"fmt"
	"sql-db/internal/parser"
	"sql-db/internal/storage"
	"strings"
)

func Execute(cmd parser.Command, db *storage.Database) (string, error) {
	switch cmd.Action {
	case "CREATE":
		db.Tables[cmd.TableName] = storage.Table{Name: cmd.TableName}
		return "table created: " + cmd.TableName, nil

	case "SELECT":
		table, exists := db.Tables[cmd.TableName]
		if !exists {
			return "", fmt.Errorf("table not found: %s", cmd.TableName)
		}
		if len(table.Rows) == 0 {
			return "empty table", nil
		}

		var result strings.Builder

		for _, row := range table.Rows {

			if cmd.WhereField != "" {
				value, exists := row[cmd.WhereField]

				if !exists || value != cmd.WhereValue {
					continue
				}
			}

			for k, v := range row {
				result.WriteString(k)
				result.WriteString("=")
				result.WriteString(v)
				result.WriteString(" ")
			}
			result.WriteString("\n")
		}

		return result.String(), nil

	case "INSERT":
		table, exists := db.Tables[cmd.TableName]

		if !exists {
			return "", fmt.Errorf("table not found: %s", cmd.TableName)
		}
		table.Rows = append(table.Rows, cmd.Data)

		db.Tables[cmd.TableName] = table

		return "row inserted", nil

	default:
		return "", fmt.Errorf("unknown action: %s", cmd.Action)
	}
}
