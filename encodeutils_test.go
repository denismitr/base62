package base62

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIntegerIDToString(t *testing.T) {
	Convey("Given an integer ID", t, func() {
		So(ToBase62(1), ShouldEqual, "1")
		So(ToBase62(2), ShouldEqual, "2")
		So(ToBase62(10), ShouldEqual, "a")
		So(ToBase62(209), ShouldEqual, "3n")
		So(ToBase62(209098), ShouldEqual, "Soy")
		So(ToBase62(20909809875), ShouldEqual, "mP5qJZ")
	})
}
