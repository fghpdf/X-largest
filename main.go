package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/fghpdf/X-largest/topx"
	"os"
	"strings"
)

var fileName string
var x int64

func init() {
	flag.StringVar(&fileName, "file", "", "file name")
	flag.Int64Var(&x, "x", 0, "x")
}

func main() {
	flag.Parse()

	if x <= 0 {
		fmt.Println("x must be greater than 0")
		return
	}

	if strings.TrimSpace(fileName) == "" {
		scanner := bufio.NewScanner(os.Stdin)
		stdinInput := make([]string, 0)
		var text string
		for text != "done" {
			fmt.Print("Enter text following <unique> <count>, input done can finish: ")
			scanner.Scan()
			text = scanner.Text()
			if text != "done" {
				stdinInput = append(stdinInput, text)
			}
		}
		records := topx.HandleStdinTopX(stdinInput, x)
		for _, record := range records {
			fmt.Println(record)
		}
	} else {
		records := topx.HandleHugeFileTopX(fileName, x)
		for _, record := range records {
			fmt.Println(record)
		}
	}
}
