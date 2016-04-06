package xee_test

import (
	"testing"
    "github.com/laibulle/xee"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAppsSpec(t *testing.T) {
	Convey("Given a sdk", t, func() {
        sdk := xee.NewSDK("myclient", "mysecret", "http://localhost")

		Convey("When get Auth URI", func() {
            uri := sdk.GetAuthURI("azerty")
			Convey("The access should be granted", func() {
                So(uri, ShouldEqual, "https://cloud.xee.com/v1/auth/auth?client_id=myclient&redirect_uri=http%3A%2F%2Flocalhost&state=azerty")
			})
		})
	})
}
