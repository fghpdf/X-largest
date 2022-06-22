package tools

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/cheggaaa/pb/v3"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// snowflake id length is 19
// count length is 5 (max)
// so we can assume one record length is 19+1+5 = 24 (including space)
const recordLength = 25

// file size 1GB
const fileSize = 1024 * 1024 * 1024

const RandomMax = 100000
const TopXTThreshold = 99990

const HugeFileName = "tmp/gen_records.txt"
const TopXFileName = "tmp/gen_top_x.txt"
const AllGenFiles = "tmp/gen_*.txt"

// GenerateHugeFile generate a huge file with random count in range [0, 100000)
// file is 1GB named as "tmp/records.txt;
// And will generate a file named "tmp/top_x.txt",
// it will save the records which is bigger than 99990,
// this file can be used to check result of topX
func GenerateHugeFile() {
	logrus.Infoln("Generating huge file, it will take a while...")
	// clear all gen files
	_ = clearGenFiles()

	recordCount := fileSize / recordLength

	// open big file
	file, err := os.OpenFile(HugeFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer file.Close()

	// open top-x file for test
	topXFile, err := os.OpenFile(TopXFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer topXFile.Close()

	// create and start new bar
	bar := pb.StartNew(recordCount)
	// init snowflake
	node, _ := snowflake.NewNode(1)
	for i := 0; i < recordCount; i++ {
		bar.Increment()
		// generate a snowflake ID.
		id := node.Generate()
		// generate random count
		count := randomNumberInRange(1, RandomMax)
		if _, err := file.WriteString(
			fmt.Sprintf("%s %d\n", id.String(), count)); err != nil {
			logrus.Fatalln(err)
		}
		if count >= TopXTThreshold {
			if _, err := topXFile.WriteString(
				fmt.Sprintf("%s %d\n", id.String(), count)); err != nil {
				logrus.Fatalln(err)
			}
		}
	}

	// finish bar
	bar.Finish()
}

func randomNumberInRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func clearGenFiles() error {
	files, err := filepath.Glob(AllGenFiles)
	if err != nil {
		return err
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return err
		}
	}

	return nil
}
