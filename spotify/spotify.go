// Package spotify:
// go-spotify provides an easy-to-use API
// to access Spotify's WEB API
package spotify

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/parnurzeal/gorequest"
)

// Spotify struct which we use
// to wrap our request operations.
type Spotify struct {
	clientID     string
	clientSecret string
	//	redirectURI        string
	accessToken string
}

const (
	BASE_URL     = "https://api.spotify.com"
	ACCOUNTS_URL = "https://accounts.spotify.com/api/token"
	API_VERSION  = "v1"
)

// Creates a New Spotify API object with the
// clientID and clientSecret
// Usage:
// 	spotify.New("XXX","YYY")
func New(clientID, clientSecret string) Spotify {

	return initialize(clientID, clientSecret)
}

func initialize(clientID, clientSecret string) Spotify {
	spot := Spotify{clientID: clientID, clientSecret: clientSecret}
	return spot
}

// Authorizes your application against Spotify
func (spotify *Spotify) Authorize() (bool, []error) {

	result := false

	// Get Encoded Access Keys for Authentication
	auth := fmt.Sprintf("Basic %s", spotify.getEncodedKeys())

	// create a new request to get our access_token
	// and send our Keys on Authorization Header
	request := gorequest.New()
	request.Post(ACCOUNTS_URL)
	request.Set("Authorization", auth)
	request.Send("grant_type=client_credentials")

	_, body, errs := request.End()

	// Parse response to simplejson object
	js, err := simplejson.NewJson([]byte(body))
	if err != nil {
		fmt.Println("[Authorize] Error parsing Json!")
		errs = []error{err}
	}

	// check whether we got the access_token or not.
	jsToken, exists := js.CheckGet("access_token")
	if exists {
		// If we got it then assign it to the spotify object.
		spotify.accessToken, err = jsToken.String()
		if err != nil {
			fmt.Println("[Authorize] Error Getting Access Token from Json!")
			errs = []error{err}
		}
		result = true
	}

	return result, errs
}

// Creates a new GET Request to Spotify and returns
// the response as a map[string]interface{}.
//
// format: target endpoint format like "albums/%s" - string
//
// data: content to be sent with the request - map[string]interface{}
//
// args: Arguments to be used based on format
//
// Usage:
// 	spotify.Get("albums/%s",nil,0sNOF9WDwhWunNAHPD3Baj)
func (spotify *Spotify) Get(format string, data map[string]interface{}, args ...interface{}) ([]byte, []error) {
	return spotify.Request("GET", format, data, args...)
}

// This version does not support User Authentication so
// POST/PUT/DELETE Request won't work.
//
// Creates a new POST Request to Spotify and returns
// the response as a map[string]interface{}.
//
// format: target endpoint format like "users/%s/playlists" - string
//
// data: content to be sent with the request - map[string]interface{}
//
// args: Arguments to be used based on format
//
// Usage:
// 	spotify.Post("users/%s/playlists",map[string]interface{},"wizzler")
func (spotify *Spotify) Post(format string, data map[string]interface{}, args ...interface{}) ([]byte, []error) {
	return spotify.Request("POST", format, data, args...)
}

// This version does not support User Authentication so
// POST/PUT/DELETE Request won't work.
//
// Creates a new PUT Request to Spotify and returns
// the response as a map[string]interface{}.
//
// format: target endpoint format like "me/tracks?ids=%s" - string
//
// data: content to be sent with the request - map[string]interface{}
//
// args: Arguments to be used based on format
//
// Usage:
// 	spotify.Put("me/tracks?ids=%s",nil,"4iV5W9uYEdYUVa79Axb7Rh")
func (spotify *Spotify) Put(format string, data map[string]interface{}, args ...interface{}) ([]byte, []error) {
	return spotify.Request("PUT", format, data, args...)
}

// This version does not support User Authentication so
// POST/PUT/DELETE Request won't work.
//
// Creates a new DELETE Request to Spotify and returns
// the response as a map[string]interface{}.
//
// format: target endpoint format like "me/tracks?ids=%s" - string
//
// args: Arguments to be used based on format
//
// Usage:
// 	spotify.Delete("me/tracks?ids=%s","4iV5W9uYEdYUVa79Axb7Rh")
func (spotify *Spotify) Delete(format string, args ...interface{}) ([]byte, []error) {
	return spotify.Request("DELETE", format, nil, args...)
}

// Creates a new Request to Spotify and returns
// the response as a map[string]interface{}.
//
// method: GET/POST/PUT/DELETE - string
//
// format: target endpoint format like "albums/%s" - string
//
// data: content to be sent with the request - map[string]interface{}
//
// args: Arguments to be used based on format
//
// Usage:
// 	spotify.request("GET","albums/%s",nil,"0sNOF9WDwhWunNAHPD3Baj")
func (spotify *Spotify) Request(method, format string, data map[string]interface{}, args ...interface{}) ([]byte, []error) {

	// create endpoint based on passed format
	endpoint := fmt.Sprintf(format, args...)

	targetURL := spotify.createTargetURL(endpoint)

	request := gorequest.New()

	// Check method type to call corresponding
	// go-request method
	if method == "GET" {
		request.Get(targetURL)
	}
	if method == "POST" {
		request.Post(targetURL)
	}
	if method == "PUT" {
		request.Put(targetURL)
	}
	if method == "DELETE" {
		request.Delete(targetURL)
	}

	request.Set("Authorization", fmt.Sprintf("Bearer %s", spotify.accessToken))

	// Add the data to the request if it
	// isn't null
	if data != nil {
		jsonData, _ := getJsonBytesFromMap(data)
		if jsonData != nil {
			request.Send(string(jsonData))
		}
	}

	_, body, errs := request.End()

	result := []byte(body)
	if unauthorizedResponse(result) {
		result = nil
		errs = []error{
			errors.New("Authorization Error. Make sure you called Spotify.Authorize() method!"),
			errors.New(body)}
	}

	return result, errs
}

// Checks for the response content to see if we
// received a not authorized error.
func unauthorizedResponse(body []byte) bool {

	// Parse response to simplejson object
	js, err := simplejson.NewJson(body)
	if err != nil {
		fmt.Println("[unauthorizedResponse] Error parsing Json!")
		return true
	}

	// check whether we got an error or not.
	_, exists := js.CheckGet("error")
	if exists {
		return true
	}

	return false
}

// Creates target URL for making a Spotify Request
// to a given endpoint
func (spotify *Spotify) createTargetURL(endpoint string) string {
	result := fmt.Sprintf("%s/%s/%s", BASE_URL, API_VERSION, endpoint)
	return result
}

// returns base64 encoded authorization
// Keys for Spotify.
func (spotify *Spotify) getEncodedKeys() string {

	data := fmt.Sprintf("%v:%v", spotify.clientID, spotify.clientSecret)
	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	return encoded
}

// Extracts Json Bytes from map[string]interface
func getJsonBytesFromMap(data map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Invalid data object, can't parse to json:")
		fmt.Println("Error:", err)
		fmt.Println("Data:", data)
		return nil, err
	}
	return jsonData, nil
}
