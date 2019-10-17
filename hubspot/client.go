package hubspot

// Client provides methods to interact with Hubspot API
type Client struct {
	apiKey string
}

// NewClient instantiates a new client to interact with Hubspot
func NewClient(api string) *Client {
	return &Client{
		apiKey:	api,
	}
}