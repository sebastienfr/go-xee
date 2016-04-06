package xee_test

import (
    "github.com/laibulle/go-xee"
)
func ExampleNewSDK_newSDK() {
    xee.NewSDK("myidentifier", "mysecret", "http://127.0.0.1:3000/xee-callback")
}
