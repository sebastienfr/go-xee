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
    tripResponseBody = `{"uuid":"56b43a4f051f29071f14218d","beginLocation":{"latitude":50.6817,"longitude":3.08202,"altitude":2,"heading":0,"hdp":4,"satCount":1,"date":"2016-01-29T18:36:17Z"},"stopLocation":{"latitude":50.6817,"longitude":3.08202,"altitude":2,"heading":0,"hdp":4,"satCount":1,"date":"2016-01-29T18:36:17Z"},"beginDate":"2016-01-29T18:39:17Z","stopDate":"2016-01-29T19:15:15Z","creationDate":"2016-01-29T18:39:17Z","lastUpdateDate":"2016-01-29T19:15:15Z"}`
)

func TestTripsSpec(t *testing.T) {
    httpmock.Activate()
    defer httpmock.DeactivateAndReset()

    sdk := xee.NewSDK("myidentifier", "mysecret", "http://localhost/xee-callback")

	Convey("Given a up Xee server", t, func() {
		Convey("When fetching cars with valid token", func() {
            httpmock.RegisterResponder("GET", "https://cloud.xee.com/v3/cars/1/trips",
            httpmock.NewStringResponder(200, fmt.Sprintf("[%s]", tripResponseBody)))

            now := time.Now()
            trips, err := sdk.FindTrips(1, validToken, &now, &now)

			Convey("No error", func() {
                So(err, ShouldBeNil)
                So(len(trips), ShouldEqual, 1)
			})
		})
	})
}
