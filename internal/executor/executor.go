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

		_, exists := db.Tables[cmd.TableName]

		if exists {
			return "", fmt.Errorf("table already exists: %s", cmd.TableName)
		}

		db.Tables[cmd.TableName] = storage.Table{Name: cmd.TableName, Columns: cmd.Columns}

		db.Save()
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

		for key := range cmd.Data {
			valid := false

			for _, col := range table.Columns{
				if key == col {
					valid = true
					break
				}
			}

			if !valid {
				return "", fmt.Errorf("invalid column '%s' for table %s", key, cmd.TableName)
			}
		}

		table.Rows = append(table.Rows, cmd.Data)

		db.Tables[cmd.TableName] = table

		db.Save()
		return "row inserted", nil

	case "SHOW_TABLES":
		var result strings.Builder

		for tableName, table := range db.Tables {
			result.WriteString(tableName)
			result.WriteString(" [")

			for i, col := range table.Columns {
				result.WriteString(col)

				if i < len(table.Columns)-1 {
					result.WriteString(" ")
				}
			}

			result.WriteString("]\n")
		}

		return result.String(), nil

	default:
		return "", fmt.Errorf("unknown action: %s", cmd.Action)
	}
}
