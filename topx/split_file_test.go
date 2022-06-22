package topx

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSplitFile(t *testing.T) {
	Convey("split not exist file will throw error", t, func() {
		So(splitFile("not_exist_file"), ShouldNotBeEmpty)
	})
}
