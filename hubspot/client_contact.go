package hubspot

import (
	"encoding/json"
	"fmt"
)

type Properties struct {
	Properties []Property `json:"properties"`
}

type Property struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}

func (c *Client) GetContactByVID(vid string, properties []string) (Response, error) {
	var propertyQuery string

	for _, property := range properties {
		propertyQuery = propertyQuery + fmt.Sprintf("&property=%s", property)
	}

	response, err := SendRequest(Request{
		URL:          fmt.Sprintf("https://api.hubapi.com/contacts/v1/contact/vid/%s/profile?hapikey=%s%s", vid, c.apiKey, propertyQuery),
		Method:       "GET",
		OkStatusCode: 200,
	})
	return response, err
}

func (c *Client) UpdateContact(contactVID string, properties *Properties) (Response, error) {
	body, err := json.Marshal(properties)

	if err != nil {
		return Response{}, fmt.Errorf("Invalid request: %s", err.Error())
	}

	response, err := SendRequest(Request{
		URL:          fmt.Sprintf("https://api.hubapi.com/contacts/v1/contact/vid/%s/profile?hapikey=%s", contactVID, c.apiKey),
		Method:       "POST",
		Body:         body,
		OkStatusCode: 204,
	})

	return response, err
}
