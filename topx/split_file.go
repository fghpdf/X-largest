package topx

import (
	"bufio"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/sirupsen/logrus"
	"hash/fnv"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

const MaxMemorySize = 1024 * 1024 * 1024 // 1G
const TmpFolderName = "tmp"
const AllTmpFiles = "tmp/tmp_*.txt"

// splitFile split a huge file into several small files line by line
func splitFile(fileName string) error {
	// clear tmp flies
	err := clearTmpFiles()
	if err != nil {
		return err
	}

	start := time.Now()

	// open file
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	fileStat, _ := file.Stat()
	// calc process num by file size
	processNum := calcProcessNum(fileStat.Size())
	logrus.Infoln("The large file will be split into", processNum,
		"small files, it will take a while...")

	// ensure tmp folder exists
	err = os.MkdirAll(TmpFolderName, os.ModePerm)
	if err != nil {
		return err
	}
	// create many small files
	tmpFlies := make([]*os.File, 0)
	for i := 0; i < processNum; i++ {
		// open file
		tmpFile, err := os.OpenFile(
			fmt.Sprintf("tmp/tmp_%d.txt", i), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		tmpFlies = append(tmpFlies, tmpFile)
	}

	// init scanner
	scanner := bufio.NewScanner(file)
	// scan file line by line
	scanner.Split(bufio.ScanLines)

	// create and start new bar
	bar := pb.StartNew(int(fileStat.Size() / int64(processNum)))
	wg := &sync.WaitGroup{}
	for scanner.Scan() {
		// hash element to a small file
		index := hash(scanner.Text()) % uint32(processNum)
		if index == 0 {
			bar.Add(len(scanner.Text()) + 1)
		}
		tmpFile := tmpFlies[index]
		wg.Add(1)
		go writeTmpFile(wg, scanner.Text(), tmpFile, index)
	}
	wg.Wait()

	// close files
	for _, tmpFile := range tmpFlies {
		tmpFile.Close()
	}
	bar.Finish()

	// calc time
	elapsed := time.Since(start)
	logrus.Infoln("split file elapsed:", elapsed)
	return nil
}

func writeTmpFile(wg *sync.WaitGroup, content string, file *os.File, index uint32) {
	defer wg.Done()
	if file == nil {
		logrus.Warnln("file is nil", index)
		return
	}
	_, err := file.WriteString(content + "\n")
	if err != nil {
		panic(err)
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func clearTmpFiles() error {
	files, err := filepath.Glob(AllTmpFiles)
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

func calcProcessNum(fileSize int64) int {
	coreCount := runtime.NumCPU()
	if coreCount == 0 {
		logrus.Panicln("core count is 0, please check your system")
	}

	tmpFileSize := MaxMemorySize / (coreCount * 16)
	processNum := fileSize / int64(tmpFileSize)
	return int(processNum)
}
