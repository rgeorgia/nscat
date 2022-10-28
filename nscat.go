package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func hasKeyWord(word string) bool {
	keyWords := []string{"pw", "password", "dbpass", "_key", "_key_", "credential", "pat"}
	for _, v := range keyWords {
		if strings.Contains(strings.ToLower(word), v) {
			return true
		}
	}
	return false
}

func main() {
	var filename string
	flag.StringVar(&filename, "f", ".env", "File Name")
	flag.Parse()

	dat, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineSlice := strings.Split(line, "=")
		if hasKeyWord(lineSlice[0]) {
			fmt.Printf("%s = %s\n", lineSlice[0], strings.Repeat("X", len(lineSlice[1])))
		} else {
			fmt.Println(line)
		}
	}
}
