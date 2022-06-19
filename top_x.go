package main

import (
	"container/heap"
)

// Record define
type Record struct {
	Identifier string
	count      int64
}

func topXFrequent(records []*Record, x int64) []string {
	recordsLen := int64(len(records))
	if recordsLen <= x {
		// all records are in the top x
		return getIdentifiers(records)
	}

	q := PriorityQueue{}
	for _, record := range records {
		heap.Push(&q, record)
	}
	var result []string
	for int64(len(result)) < x {
		item := heap.Pop(&q).(*Record)
		result = append(result, item.Identifier)
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
