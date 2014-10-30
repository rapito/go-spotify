package spotify

// Import Testing frameworks needed
import (
	"testing"
	"github.com/bmizerany/assert"
)

// Create out api variables for easy access
const (
	clientID     = "your-clientID-here"
	clientSecret = "your-secret-clientSecret-here"
)

var spotify = New(clientID, clientSecret)

// Should create a new Spotify Object.
func TestNew(t *testing.T) {

	assert.T(t, spotify.clientID == clientID, "clientID should be the same")
	assert.T(t, spotify.clientSecret == clientSecret, "clientSecret should be the same")
}

