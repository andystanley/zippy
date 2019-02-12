package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zippy/internal/parser"
	"github.com/zippy/pkg/store"
)

func main() {
	store.Open("../log")
	fmt.Println("Welcome to zippy!")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">>> ")
		scanner.Scan()
		args := strings.Split(scanner.Text(), " ")
		parser.Parse(args)
	}
}
