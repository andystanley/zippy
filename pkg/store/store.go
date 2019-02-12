package store

import (
	"bufio"
	"os"
	"strings"
)

var logPath string
var store = make(map[string]string)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Open opens a zippy store at the log path (ex. /Users/Andrew/Desktop/log). If the log
// does not exist it will be created
func Open(path string) {
	logPath = path
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_RDONLY, 0644)
	check(err)
	defer logFile.Close()

	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		switch args[0] {
		case "set":
			key := args[1]
			value := args[2]
			store[key] = value
		case "delete":
			key := args[1]
			delete(store, key)
		}
	}
}

// Clear empties the in memory zippy store and the log file
func Clear() {
	store = make(map[string]string)
	err := os.Remove(logPath)
	check(err)
}

// Delete removes a key
func Delete(key string) {
	delete(store, key)
	updateLog("delete " + key)
}

// Get retrieves a value based on a key
func Get(key string) string {
	return store[key]
}

// Keys returns a list of all the keys
func Keys() (keys []string) {
	for key := range store {
		keys = append(keys, key)
	}
	return keys
}

// Set a value
func Set(key string, value string) {
	store[key] = value
	updateLog("set " + key + " " + value)
}

// updateLog adds a new line to the append only log
func updateLog(value string) {
	logFile, err := os.OpenFile("../log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer logFile.Close()

	bytes := []byte(value + "\n")
	logFile.Write(bytes)
}
