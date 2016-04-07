package xee_test

import (
	"testing"
    "github.com/laibulle/go-xee"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/jarcoal/httpmock"
)

const (
    tokenResponse = `{"access_token":"ezddez","refresh_token":"frefrefr","expires_in":3600}`
)
func TestAppsSpec(t *testing.T) {
    httpmock.Activate()
    defer httpmock.DeactivateAndReset()

    sdk := xee.NewSDK("myclient", "mysecret", "http://localhost")

	Convey("Given a sdk", t, func() {
		Convey("When get Auth URI", func() {
            uri := sdk.GetAuthURI("azerty")
			Convey("The redirect URI should be well formatted", func() {
                So(uri, ShouldEqual, "https://cloud.xee.com/v1/auth/auth?client_id=myclient&redirect_uri=http%3A%2F%2Flocalhost&state=azerty")
			})
		})

	})

    Convey("Given a down Xee server", t, func() {
         Convey("When asking an access token", func() {
            _, err := sdk.GetTokenFromCode("azerty")
			Convey("AccessToken should be valid", func() {
                So(err, ShouldNotBeNil)
			})
		})
    })

    Convey("Given a valid code", t, func(){
        httpmock.RegisterResponder("POST", "https://cloud.xee.com/v1/auth/access_token.json",
            httpmock.NewStringResponder(201, tokenResponse))

        Convey("When asking an access token", func() {
            token, err := sdk.GetTokenFromCode("azerty")
			Convey("AccessToken should be valid", func() {
                So(err, ShouldBeNil)
                So(token.AccessToken, ShouldEqual, "ezddez")
			})
		})
    })

    Convey("Given a valid refresh_token", t, func(){
        httpmock.RegisterResponder("POST", "https://cloud.xee.com/v1/auth/access_token.json",
            httpmock.NewStringResponder(201, tokenResponse))

        Convey("When asking an access token", func() {
            token, err := sdk.GetTokenFromRefreshToken("azerty")
			Convey("AccessToken should be valid", func() {
                So(err, ShouldBeNil)
                So(token.AccessToken, ShouldEqual, "ezddez")
			})
		})
    })
}
