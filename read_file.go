package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	records, err := readSingleFile("tmp/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// get top x frequent
	topX := topXFrequent(records, 3)
	for _, identifier := range topX {
		log.Println(identifier)
	}
}

// readSingleFile read single file line by line
func readSingleFile(fileName string) ([]*Record, error) {
	// open file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	// init scanner
	scanner := bufio.NewScanner(file)

	// scan file line by line
	scanner.Split(bufio.ScanLines)

	// scan into records
	records := make([]*Record, 0)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		identifier := text[0]
		count, _ := strconv.ParseInt(text[1], 10, 64)
		records = append(records, &Record{identifier, count})
	}

	// close file
	file.Close()

	return records, nil
}
