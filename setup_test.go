package main

import (
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

var jsonToReturn = `
{
  "ts": 1753630018238,
  "tsj": 1753630015670,
  "date": "Jul 27th 2025, 11:26:55 am NY",
  "items": [
    {
      "curr": "USD",
      "xauPrice": 3336.665,
      "xagPrice": 38.184,
      "chgXau": -32.38,
      "chgXag": -0.806,
      "pcXau": -0.9611,
      "pcXag": -2.0672,
      "xauClose": 3369.045,
      "xagClose": 38.99005
    }
  ]
}
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
