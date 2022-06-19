package main

import (
	"testing"
)

func Test_topXFrequent(t *testing.T) {
	type args struct {
		records []*Record
		x       int64
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"topX", args{[]*Record{{"1426828011", 9},
			{"1426828028", 350}, {"1426828037", 25},
			{"1426828056", 231}, {"1426828058", 109},
			{"1426828066", 111}}, 3}, []string{"1426828028", "1426828066", "1426828056"}},
		{"topX with X bigger than records", args{[]*Record{{"1426828011", 9},
			{"1426828066", 111}}, 5}, []string{"1426828011", "1426828066"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topXFrequent(tt.args.records, tt.args.x); !isSameStringSlice(got, tt.want) {
				t.Errorf("topXFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}

// isSameStringSlice compare two string slice without order
func isSameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}
