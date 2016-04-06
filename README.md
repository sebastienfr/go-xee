# Xee Go SDK

This is an unofficial SDK for Xee

## Getting started


    package main
    
    import (
        "github.com/laibulle/xee"
    )
    
    func main() {
        sdk := xee.NewSDK("myclientid", "myclientsecret", "http://myapp.io/xee-calback")
    }
