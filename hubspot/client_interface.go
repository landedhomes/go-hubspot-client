package hubspot

type ClientInterface interface {
	// Pages
	SavePage() (Response, error)

	// Contact Property
	
}

var _ ClientInterface = &Client{}