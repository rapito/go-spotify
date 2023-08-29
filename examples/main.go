// Example Usage of go-spotify API
package main

import (
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/rapito/go-spotify/spotify"
	"github.com/spf13/viper"
)

func main() {

	// Create a new spotify object

	// Make .env file variables accessible in file
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	spot := spotify.New(viper.GetString("CLIENT_ID"), viper.GetString("CLIENT_SECRET"))

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
