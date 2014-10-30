// Package spotify:
// go-spotify provides an easy-to-use API
// to access Spotify's WEB API
package spotify

import (
	//	"encoding/json"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/parnurzeal/gorequest"
	"encoding/base64"
	"fmt"
)

// Spotify struct which we use
// to wrap our request operations.
type Spotify struct {
	clientID           string
	clientSecret       string
	//	redirectURI        string
	accessToken        string
}

const (
	//	baseURL     = "https://api.spotify.com"
	accountsURL = "https://accounts.spotify.com/api/token"
)

// Creates a New Spotify API object with the
// clientID and clientSecret
// Usage:
// 	spotify.New("XXX","YYY")
func New(clientID, clientSecret string) Spotify {

	return createSpotify(clientID, clientSecret)
}

//
//// Creates a New Spotify API object with the
//// clientID, clientSecret and redirectURI
//// Usage:
//// 	spotify.New("XXX","YYY", "http://ZZZ")
//func New(clientID, clientSecret, redirectURI string) Spotify {
//
//	return createSpotify(clientID, clientSecret, redirectURI)
//}

func createSpotify(clientID, clientSecret string) Spotify {
	shop := Spotify{clientID: clientID, clientSecret: clientSecret}
	return shop
}

// Authorizes your application against Spotify
func (spotify *Spotify) Authorize() (bool, []error) {

	result := false

	request := gorequest.New()
	request.Post(accountsURL)
	auth := fmt.Sprintf("Basic %s", spotify.getEncodedKeys())
	request.Set("Authorization", auth)
	request.Send("grant_type=client_credentials")

	_, body, errs := request.End()

	js, err := simplejson.NewJson([]byte(body))
	if err != nil {
		fmt.Println("[Authorize] Error parsing Json!")
		errs = []error{err}
	}

	jsToken, exists := js.CheckGet("access_token")

	if exists {
		spotify.accessToken, err = jsToken.String()
		if err != nil {
			fmt.Println("[Authorize] Error Getting Access Token from Json!")
			errs = []error{err}
		}
		result = true
	}

	return result, errs
}

// returns base64 encoded authorization
// Keys for Spotify.
func (spotify *Spotify) getEncodedKeys() string {

	data := fmt.Sprintf("%v:%v", spotify.clientID, spotify.clientSecret)
	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	return encoded;
}
