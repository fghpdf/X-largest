package tools

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

func CountMaxFromTopXFile(fileName string) int64 {
	topXFile, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer topXFile.Close()

	// init scanner
	scanner := bufio.NewScanner(topXFile)

	// scan file line by line
	scanner.Split(bufio.ScanLines)

	res := int64(0)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		count, _ := strconv.ParseInt(text[1], 10, 64)
		if count == RandomMax-1 {
			res += 1
		}
	}

	return res
}
