package topx

import (
	"container/heap"
)

// Record define
type Record struct {
	Identifier string
	count      int64
}

func getTopXFrequent(records []*Record, x int64) []string {
	return getIdentifiers(topXFrequent(records, x))
}

func topXFrequent(records []*Record, x int64) []*Record {
	recordsLen := int64(len(records))
	if recordsLen <= x {
		// all records are in the top x
		return records
	}

	q := PriorityQueue{}
	for _, record := range records {
		heap.Push(&q, record)
	}
	var result []*Record
	for int64(len(result)) < x {
		item := heap.Pop(&q).(*Record)
		result = append(result, item)
	}
	return result
}

func getIdentifiers(records []*Record) []string {
	var result []string
	for _, record := range records {
		result = append(result, record.Identifier)
	}
	return result
}
