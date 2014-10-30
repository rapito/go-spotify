package spotify

// Import Testing frameworks needed
import (
	"testing"
	"github.com/bmizerany/assert"
)

// Create out api variables for easy access
const (
	clientID     = "your-clientID-here"
	clientSecret = "your-client-secret-here"
)

var spotify = New(clientID, clientSecret)

// Should create a new Spotify Object.
func TestNew(t *testing.T) {

	assert.T(t, spotify.clientID == clientID, "clientID should be the same")
	assert.T(t, spotify.clientSecret == clientSecret, "clientSecret should be the same")
}

// Should create a new Spotify Object.
func TestAuthorize(t *testing.T) {

	result, err := spotify.Authorize();
	assert.T(t, result, "should be true")
	assert.T(t, len(err) == 0, "should be nil")

	assert.T(t, len(spotify.accessToken) > 0, "should not be nil")

}

// Should create a new Spotify Object.
func TestRequest(t *testing.T) {

	result, err := spotify.Request("GET", "albums/%s", nil, "0sNOF9WDwhWunNAHPD3Baj");

	assert.T(t, result != nil, "Shouldnt be null")
	assert.T(t, err == nil, "Should be null")

}

// Should create a new Spotify Object.
func TestGet(t *testing.T) {

	result, err := spotify.Get("albums/%s", nil, "0sNOF9WDwhWunNAHPD3Baj")
	assert.T(t, result != nil, "Shouldnt be null")
	assert.T(t, err == nil, "Should be null")

}

// Should create a new Spotify Object.
func TestGetEncodedKeys(t *testing.T) {

	result := spotify.getEncodedKeys();
	assert.T(t, len(result) > 0, "shouldn't be null")

}

// Should create a new Spotify Object.
func TestUnauthorizedResponse(t *testing.T) {

	result := unauthorizedResponse([]byte(`"{error: {	status: 401, message: "Notoken provided"}}"`))
	assert.T(t, result, "should be true")

}

// Should create a new Spotify Object.
func TestCreateTargetURL(t *testing.T) {

	result := spotify.createTargetURL("albums");
	assert.T(t, result == "https://api.spotify.com/v1/albums", "shouldn be same URL")

}

