package topx

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func HandleHugeFileTopX(fileName string, x int64) []string {
	start := time.Now()

	// check file exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		logrus.Panicf("File %s does not exist\n", fileName)
	}

	// split file
	err := splitFile(fileName)
	if err != nil {
		logrus.Panicf("Split file %s failed: %s\n", fileName, err)
	}

	splitTime := time.Since(start)
	logrus.Infof("Split file finished, it takes %s\n", splitTime)

	// get top x frequent from split files
	records, err := getTopXFromSplitFiles(x)
	if err != nil {
		logrus.Panicf("Get top x frequent from split files failed: %s\n", err)
	}

	topXTime := time.Since(start)
	logrus.Infof("Get top x frequent finished, it takes %s\n", topXTime-splitTime)

	return records
}

func HandleStdinTopX(stdinInput []string, x int64) []string {
	if len(stdinInput) == 0 {
		return nil
	}

	// read stdin
	stdinRecords := readStdin(stdinInput)

	// top x frequent
	records := topXFrequent(stdinRecords, x)

	return getIdentifiers(records)
}

func HandleSmallFileTopX(fileName string, x int64) []string {
	// check file exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		logrus.Panicf("File %s does not exist\n", fileName)
	}

	// get top x frequent from small file
	tmpRecords, err := readSingleFile(fileName)
	if err != nil {
		logrus.Panicf("Get top x frequent from split files failed: %s\n", err)
	}

	// top x frequent
	records := topXFrequent(tmpRecords, x)

	return getIdentifiers(records)
}

func getTopXFromSplitFiles(x int64) ([]string, error) {
	// read all split files
	files, err := filepath.Glob("tmp/tmp_*.txt")
	if err != nil {
		return nil, err
	}

	recordsQueue := make(chan []*Record, 1)

	// handle each file
	wg := sync.WaitGroup{}
	wg.Add(len(files))
	records := make([]*Record, 0)
	for _, fileName := range files {
		go func(fileName string) {
			tmpRecords, err := readSingleFile(fileName)
			if err != nil {
				panic(err)
			}
			tmpRecords = topXFrequent(tmpRecords, x)
			recordsQueue <- tmpRecords
		}(fileName)
	}

	go func() {
		for tmpRecords := range recordsQueue {
			records = append(records, tmpRecords...)
			wg.Done()
		}
	}()
	wg.Wait()

	// top x frequent
	return getTopXFrequent(records, x), nil
}
