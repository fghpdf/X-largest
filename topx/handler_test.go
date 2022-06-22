package topx

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHandleHugeFileTopX(t *testing.T) {
	Convey("handle not exist file will panic", t, func() {
		So(func() {
			HandleHugeFileTopX("not_exist_file", 1)
		}, ShouldPanic)
	})
}

func TestHandleSmallFileTopX(t *testing.T) {
	Convey("handle not exist file will panic", t, func() {
		So(func() {
			HandleSmallFileTopX("not_exist_file", 1)
		}, ShouldPanic)
	})
}
