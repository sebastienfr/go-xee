package xee_test

import (
	"testing"
    "github.com/laibulle/xee"
    "github.com/jarcoal/httpmock"

	. "github.com/smartystreets/goconvey/convey"
)

const (
    validToken   = "validtoken"
    invalidToken = "invalidtoken"
)

const (
    carResponseBody = `[{"id":1337,"uuid":"110e8400-e29b-11d4-a716-446655440000","name":"Mark-42","make":"Mark","model":"42","year":2014,"numberPlate":"M-42-TS","deviceId":42,"cardbId":210,"creationDate":"2014-09-23T12:49:48+00:00","lastUpdateDate":"2016-02-19T08:41:58+00:00"}]`
)

func TestCarSpec(t *testing.T) {
    httpmock.Activate()
    defer httpmock.DeactivateAndReset()

	Convey("Given a sdk", t, func() {
        sdk := xee.NewSDK("myclient", "mysecret", "http://localhost")

		Convey("When fetching cars with valid token", func() {
            httpmock.RegisterResponder("GET", "https://cloud.xee.com/v3/users/1/cars",
            httpmock.NewStringResponder(200, carResponseBody))

            cars, err := sdk.FindCars(1, validToken)
			Convey("No error", func() {
                So(len(cars), ShouldEqual, 1)
                So(err, ShouldBeNil)
			})
		})

        Convey("When fetching cars with invalid token", func() {
            httpmock.RegisterResponder("GET", "https://cloud.xee.com/v3/users/2/cars",
            httpmock.NewStringResponder(403, `[]`))

            cars, err := sdk.FindCars(2, validToken)
			Convey("No error", func() {
                So(len(cars), ShouldEqual, 0)
                So(err, ShouldEqual, xee.ErrForbidden)
			})
		})
	})
}