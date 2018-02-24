package base62

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIntegerIDToBase62String(t *testing.T) {
	Convey("Given an integer ID", t, func() {
		So(ToBase62(1), ShouldEqual, "1")
		So(ToBase62(2), ShouldEqual, "2")
		So(ToBase62(10), ShouldEqual, "a")
		So(ToBase62(209), ShouldEqual, "3n")
		So(ToBase62(209098), ShouldEqual, "Soy")
		So(ToBase62(20909809875), ShouldEqual, "mP5qJZ")
	})
}

func TestBase62StringToInteger(t *testing.T) {
	Convey("Given a base62 string", t, func() {
		So(ToBase10("mP5qJZ"), ShouldEqual, 20909809875)
		So(ToBase10("Soy"), ShouldEqual, 209098)
		So(ToBase10("3n"), ShouldEqual, 209)
		So(ToBase10("a"), ShouldEqual, 10)
		So(ToBase10("2"), ShouldEqual, 2)
		So(ToBase10("1"), ShouldEqual, 1)
	})
}
