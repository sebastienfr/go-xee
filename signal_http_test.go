package xee_test

import (
	"testing"
    "github.com/laibulle/go-xee"
    "github.com/jarcoal/httpmock"
    "fmt"
    "time"

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
		Convey("When fetching car signals", func() {
            httpmock.RegisterResponder("GET", "https://cloud.xee.com/v3/cars/1/signals",
            httpmock.NewStringResponder(200, fmt.Sprintf("[%s]", signalsResponseBody)))

            now := time.Now()
            limit := 1
            names := []string{"VehicleSpeed"}

            signals, err := sdk.FindSignals(1, validToken, &names, &limit, &now, &now)

			Convey("No error", func() {
                So(err, ShouldBeNil)
                So(len(signals), ShouldEqual, 1)
			})
		})

        Convey("When fetching trip signals", func() {
            httpmock.RegisterResponder("GET", "https://cloud.xee.com/v3/cars/1/trips/1/signals",
            httpmock.NewStringResponder(200, fmt.Sprintf("[%s]", signalsResponseBody)))

            names := []string{"VehicleSpeed"}

            signals, err := sdk.FindSignalsByTrip(1, "1", validToken, &names)

			Convey("No error", func() {
                So(err, ShouldBeNil)
                So(len(signals), ShouldEqual, 1)
			})
		})
	})
}
