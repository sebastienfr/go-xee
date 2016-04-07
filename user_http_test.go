package xee_test

import (
	"github.com/jarcoal/httpmock"
	"github.com/laibulle/go-xee"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	userResponseBody = `{"id":42,"uuid":"110e8400-e29b-11d4-a716-446655440000","lastName":"Doe","firstName":"John","nickname":"Johny","gender":"MALE","birthDate":"2016-01-11T00:00:00+00:00","licenceDeliveryDate":"2014-08-13T00:00:00+00:00","role":"dev","isLocationEnabled":true,"creationDate":"2014-08-13T15:20:58+00:00","lastUpdateDate":"2016-02-12T09:07:47+00:00"}`
)

func TestUserSpec(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	sdk := xee.NewSDK("myidentifier", "mysecret", "http://localhost/xee-callback")

	Convey("Given a up Xee server", t, func() {
		Convey("When fetching cars with valid token", func() {
			httpmock.RegisterResponder("GET", "https://cloud.xee.com/v3/users/me",
				httpmock.NewStringResponder(200, userResponseBody))

			user, err := sdk.GetMe(validToken)

			Convey("No error", func() {
				So(err, ShouldBeNil)
				So(user.FirstName, ShouldEqual, "John")
			})
		})
	})
}
