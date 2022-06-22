package main

import (
	"github.com/fghpdf/X-largest/tools"
	"github.com/fghpdf/X-largest/topx"
	"github.com/fghpdf/attic"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStdin(t *testing.T) {
	Convey("stdin input 6 lines and x is 3", t, func() {
		stdinInput := []string{"1426828011 9", "1426828028 350", "1426828037 25",
			"1426828056 231", "1426828058 109", "1426828066 111"}
		records := topx.HandleStdinTopX(stdinInput, 3)
		want := []string{"1426828028", "1426828066", "1426828056"}
		So(records, attic.ShouldSameWithoutOrder, want)
	})

	Convey("stdin input 6 lines and x is 7", t, func() {
		stdinInput := []string{"1426828011 9", "1426828028 350", "1426828037 25",
			"1426828056 231", "1426828058 109", "1426828066 111"}
		records := topx.HandleStdinTopX(stdinInput, 7)
		want := []string{"1426828011", "1426828028", "1426828066",
			"1426828058", "1426828037", "1426828056"}
		So(records, attic.ShouldSameWithoutOrder, want)
	})

	Convey("stdin input empty lines and x is 7", t, func() {
		var stdinInput []string
		records := topx.HandleStdinTopX(stdinInput, 7)
		var want []string
		So(records, attic.ShouldSameWithoutOrder, want)
	})
}

func TestHugeFile(t *testing.T) {
	Convey("get top x from huge file", t, func() {
		// get max number count from top_x file
		x := tools.CountMaxFromTopXFile(tools.TopXFileName)

		// get top x from top_x file
		want := topx.HandleSmallFileTopX(tools.TopXFileName, x)

		// get top x from huge file
		records := topx.HandleHugeFileTopX(tools.HugeFileName, x)
		So(records, attic.ShouldSameWithoutOrder, want)
	})
}
