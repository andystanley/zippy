package parser

import (
	"bufio"
	"fmt"
	"os"

	"github.com/zippy/pkg/store"
)

// Parse parses REPL commands
func Parse(args []string) {
	switch args[0] {
	case "exit":
		os.Exit(0)
	case "help":
		fmt.Println("Available Commands:")
		fmt.Println("	set key value")
		fmt.Println("	get key")
		fmt.Println("	delete key")
		fmt.Println("	keys")
	case "keys":
		keys := store.Keys()
		for _, key := range keys {
			fmt.Println(key)
		}
	case "delete":
		if len(args) == 2 {
			store.Delete(args[1])
		} else {
			fmt.Println("Usage: delete key")
		}
	case "get":
		if len(args) == 2 {
			fmt.Println(store.Get(args[1]))
		} else {
			fmt.Println("Usage: get key")
		}
	case "set":
		if len(args) == 3 {
			store.Set(args[1], args[2])
		} else {
			fmt.Println("Usage: set key value")
		}
	case "clear":
		if len(args) == 1 {
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("Are you sure you wish to delete the zippy store? yes to confirm")
			scanner.Scan()
			if scanner.Text() == "yes" {
				store.Clear()
			}
		}
	default:
		fmt.Println("Invalid Command")
	}
}
