package parser

import (
	"fmt"
	"strings"
)

type Command struct {
	Action string
	Target string
	Name   string
}

func Parse(input string) (Command, error) {

	input = strings.TrimSpace(input)

	if input == ""{
		return Command{}, fmt.Errorf("empty input")
	}

	parts := strings.Split(input, " ")

	switch parts[0]{
	case "CREATE":	
		if len(parts) != 3{
			return Command{}, fmt.Errorf("invalid CREATE syntax")
		}
		return Command{
			Action: "CREATE",
			Target: parts[1],
			Name: parts[2],
		}, nil

	case "SELECT":
        if len(parts) != 2 {
            return Command{}, fmt.Errorf("invalid SELECT syntax")
        }
        return Command{
            Action: "SELECT",
            Name:   parts[1],
        }, nil

	default:
		return Command{}, fmt.Errorf("uknown command: %s", parts[0])			

	}

}