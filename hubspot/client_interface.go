package hubspot

// ClientInterface is use to store function declarations for all Hubspot API calls on the Client Struct
type ClientInterface interface {
	// Pages
	SavePage(request *PageBody) (Response, error)
	PublishPage(request *PublishPageBody, id string) (Response, error)

	// Contact List Property
	GetContactsInList(listID string) (Response, error)
	CreateList(req *ListBody) (Response, error)

	// Contact Property
	GetContactPropertyByName(contactName string) (Response, error)
}

var _ ClientInterface = &Client{}