package xee_test

import (
	"testing"
    "github.com/laibulle/go-xee"
    "github.com/jarcoal/httpmock"
    "fmt"

	. "github.com/smartystreets/goconvey/convey"
)

const (
    signalsResponseBody = `{"name":"LockSts","value":0,"source":"CAN","date":"2016-03-01T02:24:24.000000+00:00"}`
)

func TestSignalsSpec(t *testing.T) {
    httpmock.Activate()
    defer httpmock.DeactivateAndReset()

    sdk := xee.NewSDK("myidentifier", "mysecret", "http://localhost/xee-callback")

	Convey("Given a up Xee server", t, func() {
		Convey("When fetching cars with valid token", func() {
            httpmock.RegisterResponder("GET", "https://cloud.xee.com/v3/cars/1/signals",
            httpmock.NewStringResponder(200, fmt.Sprintf("[%s]", signalsResponseBody)))

            signals, err := sdk.FindSignals(1, validToken, nil, nil, nil, nil)

			Convey("No error", func() {
                So(err, ShouldBeNil)
                So(len(signals), ShouldEqual, 1)
			})
		})
	})
}
