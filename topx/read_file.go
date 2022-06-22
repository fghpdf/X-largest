package topx

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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
