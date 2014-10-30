// Example Usage of go-spotify API
package main

import (
	"github.com/rapito/go-spotify/spotify"
	simplejson "github.com/bitly/go-simplejson"
	"fmt"
)

func main() {

	// Create a new spotify object
	spot := spotify.New("756100c43d724ae6b791f7804c82b219", "841a197013f847bca78c01b3b69fc72d")

	// Authorize against Spotify first
	authorized, _ := spot.Authorize()
	if authorized {

		// If we ere able to authorize then Get a simple album
		response, _ := spot.Get("albums/%s", nil, "0sNOF9WDwhWunNAHPD3Baj")

		// Parse response to a JSON Object and
		// get the album's name
		json, _ := simplejson.NewJson(response)
		jsonData, exists := json.CheckGet("name")

		if exists {
			albumName, _ := jsonData.String()
			//Let's see what we got :)
			fmt.Println("Album name is ", albumName)
		} else {
			fmt.Println("Something went wrong... You should probably checkout the errors.")
		}

	}

}
