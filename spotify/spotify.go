// Package spotify:
// go-spotify provides an easy-to-use API
// to access Spotify's WEB API
package spotify

import (
//	"fmt"
//	"encoding/json"
//	"github.com/parnurzeal/gorequest"
)

// Spotify store struct which we use
// to wrap our request operations.
type Spotify struct {
	clientID        string
	clientSecret    string
}

const (
//	baseURL = "https://api.spotify.com"
)

// Creates a New Spotify API object with the
// clientID and clientSecret
// Usage:
// 	spotify.New("XXX","YYY")
func New(clientID, clientSecret string) Spotify {

	shop := Spotify{clientID: clientID, clientSecret: clientSecret}

	return shop
}
