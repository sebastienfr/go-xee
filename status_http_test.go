package xee_test

import (
	"github.com/jarcoal/httpmock"
	"github.com/laibulle/go-xee"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	statusResponseBody = `{"accelerometer":{"x":-768,"y":240,"z":4032,"date":"2016-03-01T02:24:20.000000+00:00"},"location":{"latitude":50.67815,"longitude":3.208155,"altitude":31.8,"satellites":4,"heading":167,"date":"2016-03-01T02:24:20.000000+00:00"},"signals":[{"name":"LockSts","value":0,"source":"CAN","date":"2016-03-01T02:24:24.000000+00:00"},{"name":"Odometer","value":34512.1,"source":"CAN","date":"2016-03-01T02:24:27.116000+00:00"}]}`
)

func TestStatusSpec(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	sdk := xee.NewSDK("myidentifier", "mysecret", "http://localhost/xee-callback")

	Convey("Given a up Xee server", t, func() {
		Convey("When fetching cars with valid token", func() {
			httpmock.RegisterResponder("GET", "https://cloud.xee.com/v3/cars/1/status",
				httpmock.NewStringResponder(200, statusResponseBody))

			status, err := sdk.FindCarStatus(1, validToken)

			Convey("No error", func() {
				So(err, ShouldBeNil)
				So(len(status.Signals), ShouldEqual, 2)
			})
		})
	})
}
