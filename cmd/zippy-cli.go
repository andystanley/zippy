package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/zippy/internal/parser"
	"github.com/zippy/pkg/store"
)

func main() {
	path := flag.String("path", "./log", "path of the zippy store to open")
	flag.Parse()
	store.Open(*path)

	fmt.Println("Welcome to zippy!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">>> ")
		scanner.Scan()
		args := strings.Split(scanner.Text(), " ")
		parser.Parse(args)
	}
}
