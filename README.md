[![Build Status](https://travis-ci.org/laibulle/go-xee.svg?branch=master)](https://travis-ci.org/laibulle/go-xee)

# Xee Go SDK

This is an unofficial SDK for Xee

## Getting started


    package main
    
    import (
        "github.com/laibulle/go-xee"
    )
    
    func main() {
        sdk := xee.NewSDK("myclientid", "myclientsecret", "http://myapp.io/xee-calback")
    }
