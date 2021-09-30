package design_patterns

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetInstance(t *testing.T) {
	Convey("单例模式", t, func() {
		ins := GetInstance()
		ins.Add(3)
		So(ins.Get(), ShouldEqual, 3)
	})
}
