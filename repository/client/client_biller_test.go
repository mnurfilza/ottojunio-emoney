package client

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBillersService(t *testing.T) {
	Convey("testing biller service", t, func() {
		billerSvc := NewHttpClient()
		Convey("testing for biiler service list", func() {
			res, err := billerSvc.GetListBiller()
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(len(res.Data), ShouldBeGreaterThan, 0)
			So(res.Code, ShouldEqual, 200)
		})

		Convey("testing biller detail service", func() {
			res, err := billerSvc.GetDetailBiller("1")
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.Data.Product, ShouldEqual, "PLN Token")
		})
	})
}
