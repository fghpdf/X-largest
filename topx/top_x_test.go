package topx

import (
	"github.com/fghpdf/attic"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetTopXFrequent(t *testing.T) {
	Convey("get top 3 frequent from 6 records", t, func() {
		records := []*Record{{"1426828011", 9},
			{"1426828028", 350}, {"1426828037", 25},
			{"1426828056", 231}, {"1426828058", 109},
			{"1426828066", 111}}
		got := getTopXFrequent(records, 3)
		want := []string{"1426828028", "1426828066", "1426828056"}
		So(got, attic.ShouldSameWithoutOrder, want)
	})

	Convey("get top 5 frequent from 2 records with X bigger than records", t, func() {
		records := []*Record{{"1426828011", 9},
			{"1426828066", 111}}
		got := getTopXFrequent(records, 5)
		want := []string{"1426828011", "1426828066"}
		So(got, attic.ShouldSameWithoutOrder, want)
	})

	Convey("get top 5 frequent from empty records", t, func() {
		var records []*Record
		got := getTopXFrequent(records, 5)
		var want []string
		So(got, attic.ShouldSameWithoutOrder, want)
	})
}
