package parser

import (
	"fmt"
	"strings"
)

type Command struct {
	Action     string
	Target     string
	TableName  string
	Data       map[string]string
	WhereField string
	WhereValue string
}

func Parse(input string) (Command, error) {

	input = strings.TrimSpace(input)

	if input == "" {
		return Command{}, fmt.Errorf("empty input")
	}

	parts := strings.Split(input, " ")

	switch parts[0] {
	case "CREATE":
		if len(parts) != 3 {
			return Command{}, fmt.Errorf("invalid CREATE syntax")
		}
		return Command{
			Action:    "CREATE",
			Target:    parts[1],
			TableName: parts[2],
		}, nil

	case "SELECT":
		if len(parts) == 2 {
			return Command{
				Action:    "SELECT",
				TableName: parts[1],
			}, nil
		}

		if len(parts) == 4 {
			if parts[2] == "WHERE" {
				kv := strings.Split(parts[3], "=")
				if len(kv) != 2 {
					return Command{}, fmt.Errorf("invalid WHERE clause syntax: expected key=value")
				}

				return Command{
					Action:     "SELECT",
					TableName:  parts[1],
					WhereField: kv[0],
					WhereValue: kv[1],
				}, nil
			}
		}

		return Command{}, fmt.Errorf("invalid SELECT syntax. Use 'SELECT table' or 'SELECT table WHERE k=v'")

	case "INSERT":
		if len(parts) < 3 {
			return Command{}, fmt.Errorf("invalid INSERT syntax")
		}

		data := make(map[string]string)

		for _, pair := range parts[2:] {
			kv := strings.Split(pair, "=")

			if len(kv) != 2 {
				return Command{}, fmt.Errorf("invalid key=value pair: %s", pair)
			}

			data[kv[0]] = kv[1]
		}

		return Command{
			Action:    "INSERT",
			TableName: parts[1],
			Data:      data,
		}, nil

	default:
		return Command{}, fmt.Errorf("uknown command: %s", parts[0])

	}

}
