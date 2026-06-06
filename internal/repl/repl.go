package repl

import (
	"bufio"
	"fmt"
	"os"
	"sql-db/internal/executor"
	"sql-db/internal/parser"
	"sql-db/internal/storage"
)


func Start(db *storage.Database) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Mini SQL DB started. Type commands:")

	for {
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')

		cmd, err := parser.Parse(input)

		if err != nil{
			fmt.Println("Error: ", err)
			continue
		}

		result, err := executor.Execute(cmd, db)

		if err != nil{
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println(result)
	}
}