package spotify

// Import Testing frameworks needed
import (
	"testing"
	"github.com/bmizerany/assert"
)

// Create out api variables for easy access
const (
	clientID     = "756100c43d724ae6b791f7804c82b219"
	clientSecret = "841a197013f847bca78c01b3b69fc72d"
	redirectURI  = "http://localhost/callback"

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
func TestGetEncodedKeys(t *testing.T) {


	result := spotify.getEncodedKeys();
	assert.T(t, len(result) > 0, "shouldn't be null")

}

