[![Build Status](https://travis-ci.org/laibulle/go-xee.svg?branch=master)](https://travis-ci.org/laibulle/go-xee)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/laibulle/go-xee/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/laibulle/go-xee)
[![Coverage Status](http://codecov.io/github/laibulle/go-xee/coverage.svg?branch=master)](http://codecov.io/github/laibulle/go-xee?branch=master)

# Xee Go SDK

This is an unofficial SDK for Xee. 

## Getting started


    package main

    import (
        "github.com/laibulle/go-xee"
        "net/http"
        "io"
        "fmt"
    )

    var (
        sdk *xee.SDK
    )

    // Handle home and redirect Xee Login
    func xeeCallback(w http.ResponseWriter, r *http.Request) {
        code := r.URL.Query().Get("code")

        // Get a token from auth code
        token, err := sdk.GetTokenFromCode(code)
        checkError(err)
        io.WriteString(w, fmt.Sprint(token))

        // Fetch user from Xee
        user, err := sdk.GetMe(token.AccessToken)
        checkError(err)
        io.WriteString(w, fmt.Sprint(user))

        // Fetch User cars
        cars, err := sdk.FindCars(user.ID, token.AccessToken)
        checkError(err)
        io.WriteString(w, fmt.Sprint(cars))

        // Fetch Car Status
        status, err := sdk.FindCarStatus(cars[0].ID, token.AccessToken)
        checkError(err)
        io.WriteString(w, fmt.Sprint(status))

        // Fetch Car locations
        locations, err := sdk.FindLocations(cars[0].ID, token.AccessToken, nil, nil, nil)
        checkError(err)
        io.WriteString(w, fmt.Sprint(locations))
    }

    // Handle authorization code
    func home(w http.ResponseWriter, r *http.Request) {
        url := sdk.GetAuthURI("1")
        http.Redirect(w, r, url, http.StatusFound)
    }

    func checkError(err error) {
        if err != nil {
            panic(err)
        }
    }

    func main() {
        // Create a SDK instance
        sdk = xee.NewSDK("myclient", "mysecret", "http://127.0.0.1:8000/xee-callback")

        http.HandleFunc("/", home)
        http.HandleFunc("/xee-callback", xeeCallback)
        http.ListenAndServe(":8000", nil)
    }
