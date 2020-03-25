package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

var (
	// OauthState is a random state for generating auth code url to mitigate CSRF attacks.
	OAuthState string

	// DefaultScopes describes all available HubSpot scopes.
	DefaultScopes = []string{
		"contacts",
		"content",
		"reports",
		"social",
		"automation",
		"actions",
		"timeline",
		"business-intelligence",
		"oauth",
		"forms",
		"files",
		"hubdb",
		"integration-sync",
		"tickets",
		"e-commerce",
		"sales-email-read",
		"forms-uploaded-files",
	}
)

func init() {
	var err error

	OAuthState, err = randState()
	if err != nil {
		panic(err)
	}
}

// randState randomly generates a base64 encoded string of
// 10 bytes
func randState() (state string, err error) {
	buffer := make([]byte, 10)
	_, err = rand.Read(buffer)
	return base64.URLEncoding.EncodeToString(buffer), err
}

// Client provides methods to interact with the HubSpot API
type Client struct {
	OAuth2Config *oauth2.Config
	Timeout      time.Duration
	Scopes       []string
}

// NewClient instantiates a new client to interact with the HubSpot API. Please
// refer to the official HubSPot documentation to obtain the required parameters.
func NewClient(clientID, clientSecret, redirectURL string) Client {
	return Client{
		OAuth2Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint:     HubSpotEndpoint,
			RedirectURL:  redirectURL,
			// Scopes:       DefaultScopes,
		},
		Timeout: 5 * time.Second,
		Scopes:  DefaultScopes,
	}
}

// SetScopes is a variadic function that sets the required scopes
// for your application for authentication with end users.
// Default behaviour is to use all scope available.
func (c *Client) SetScopes(scopes ...string) {
	c.Scopes = scopes
}

// getContext creates a context object
func (c *Client) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), c.Timeout)
}

// GetAuthCodeURL obtains the user authentication URL
func (c *Client) GetAuthCodeURL() string {
	authCodeURL := c.OAuth2Config.AuthCodeURL(OAuthState)
	if len(c.Scopes) > 0 {
		authCodeURL += "&scope=" + strings.Join(c.Scopes, "%20")
	}

	return authCodeURL
}

// GetAccessToken obtains the access token for the authenticated user
func (c *Client) GetAccessToken(code string) (token *oauth2.Token, err error) {
	ctx, cancel := c.getContext()
	defer cancel()

	token, err = c.OAuth2Config.Exchange(ctx, code)

	return
}
