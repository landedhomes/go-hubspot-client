package hubspot

import (
	"fmt"
)

// GetContactPropertyByName returns a contact property metadata by name
func (c *Client) GetContactPropertyByName(propertyName string) (Response, error) {
	response, err := SendRequest(Request{
		URL:          fmt.Sprintf("https://api.hubapi.com/properties/v1/contacts/properties/named/%s?hapikey=%s", propertyName, c.apiKey),
		Method:       "GET",
		OkStatusCode: 200,
	})

	return response, err
}
