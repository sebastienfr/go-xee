package xee_test

import (
	"testing"
    "github.com/laibulle/go-xee"
    "github.com/jarcoal/httpmock"
    "fmt"

	. "github.com/smartystreets/goconvey/convey"
)

const (
    locationResponseBody = `{"latitude":50.67815,"longitude":3.208155,"altitude":31.8,"satellites":4,"heading":167,"date":"2016-03-01T02:24:20.000000+00:00"}`
)

func TestLocationSpec(t *testing.T) {
    httpmock.Activate()
    defer httpmock.DeactivateAndReset()

    sdk := xee.NewSDK("myidentifier", "mysecret", "http://localhost/xee-callback")

	Convey("Given a up Xee server", t, func() {
		Convey("When fetching cars with valid token", func() {
            httpmock.RegisterResponder("GET", "https://cloud.xee.com/v3/cars/1/locations",
            httpmock.NewStringResponder(200, fmt.Sprintf("[%s]", locationResponseBody)))

            locations, err := sdk.FindLocations(1, validToken, nil, nil, nil)

			Convey("No error", func() {
                So(err, ShouldBeNil)
                So(len(locations), ShouldEqual, 1)
			})
		})
	})
}
