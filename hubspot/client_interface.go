package hubspot

// ClientInterface is use to store function declarations for all Hubspot API calls on the Client Struct
type ClientInterface interface {
	// Pages
	SavePage(request *PageBody) (Response, error)

	// Contact Property

}

var _ ClientInterface = &Client{}