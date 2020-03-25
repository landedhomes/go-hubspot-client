package oauth

import "golang.org/x/oauth2"

// HubSpotEndpoint is the enpoints for HubSpot API
var HubSpotEndpoint = oauth2.Endpoint{
	AuthURL:   "https://app.hubspot.com/oauth/authorize",
	TokenURL:  "https://api.hubapi.com/oauth/v1/token",
	AuthStyle: oauth2.AuthStyleAutoDetect,
}
