package topx

import (
	"strconv"
	"strings"
)

func readStdin(stdinRecords []string) []*Record {
	records := make([]*Record, 0)
	for _, stdinRecord := range stdinRecords {
		text := strings.Split(stdinRecord, " ")
		identifier := text[0]
		count, _ := strconv.ParseInt(text[1], 10, 64)
		records = append(records, &Record{identifier, count})
	}
	return records
}
