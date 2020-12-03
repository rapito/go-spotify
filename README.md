go-spotify
==========

[![GoDoc](https://godoc.org/github.com/rapito/go-spotify/spotify?status.svg)](https://godoc.org/github.com/rapito/go-spotify/spotify)  [![baby-gopher](https://raw.github.com/drnic/babygopher-site/gh-pages/images/babygopher-logo-small.png)](http://www.babygopher.org)

Simple Go library for the Spotify Web API


Installation
------------
```
go get github.com/rapito/go-spotify/spotify
```

How-to-use
----------

- Get Requests

```
    import "fmt"
    import "github.com/rapito/go-spotify/spotify"
    ...
    
    spot := spotify.New(clientID,clientSecret)
    result, _ := spot.Get("albums/%s", nil, "0sNOF9WDwhWunNAHPD3Baj")
    
    fmt.Println(string(result))
```

- Check out the *examples* folder for simple usage.
- Read some of the tests at *spotify_test.go* for more examples.

Contribution
------------
 
 - You may fork this library and modify it as you please.
 - You can make a pull request and I will be happy to check it out and merge it.
 - If you find a bug, create an issue and I will do my best to fix it (someday). 

Constraints
-------------

Right now this API version can only make successful GET requests since 
it just authenticates the application and not an specific user. 
This means you can use limited endpoints of spotify WEB API, as 
documented on their web page:

```
The access token allows you to make requests to the Spotify Web 
API endpoints that do not require user authorization such as the Get a 
track endpoint...
```

To see which endpoints you have access to, check out spotify web-api 
[documentation.](https://developer.spotify.com/web-api/)

Buy me a Drink
-------------
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate?hosted_button_id=FFC6KNAX9SKZU)

Links
-----

While I was on my *go-trip* to create this api, I found some awesome libs/links which made 
my life easier.
Check them out, hopefully they'll do the same for you:
 
 - http://github.com/parnurzeal/gorequest
 - http://github.com/bmizerany/assert
 - http://github.com/bitly/go-simplejson
 - http://github.com/avelino/awesome-go
 
 Other APIs
 ----------
 
 - http://github.com/rapito/go-shopify
 
 
