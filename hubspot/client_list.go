package hubspot

import (
	"encoding/json"
	"fmt"
)

// ListFilter is the Hubspot conditional logic use to filter Hubspot Contact list properties
type ListFilter struct {
	Operator string `json:"operator"`
	Form     string `json:"form,omitempty"`
	Page     string `json:"page,omitempty"`
	Value    string `json:"value,omitempty"`
	Property string `json:"property,omitempty"`
	Type     string `json:"type,omitempty"`
}

// ListBody is the request body structure for the Hubspot CreateList API
type ListBody struct {
	Name    string         `json:"name"`
	Dynamic bool           `json:"dynamic,omitempty"`
	Filters [][]ListFilter `json:"filters,omitempty"`
}

// GetContactsInList return Contacts associated to a given List ID
func (c *Client) GetContactsInList(listID string) (Response, error) {
	response, err := SendRequest(Request{
		URL:          fmt.Sprintf("https://api.hubapi.com/contacts/v1/lists/%s/contacts/all?hapikey=%s", listID, c.apiKey),
		Method:       "GET",
		OkStatusCode: 200,
	})

	return response, err
}

// CreateList generate a new Hubspot List given the request body criterion
func (c *Client) CreateList(req *ListBody) (Response, error) {
	body, err := json.Marshal(req)

	if err != nil {
		return Response{}, fmt.Errorf("Invalid request: %s", err.Error())
	}

	response, err := SendRequest(Request{
		URL:          fmt.Sprintf("https://api.hubapi.com/contacts/v1/lists?hapikey=%s", c.apiKey),
		Method:       "POST",
		Body:         body,
		OkStatusCode: 200,
	})

	return response, err
}

// UpdateList update an existing list on Hubspot with metadata from req body
func (c *Client) UpdateList(req *ListBody, listID string) (Response, error) {
	body, err := json.Marshal(req)

	if err != nil {
		return Response{}, fmt.Errorf("Invalid request: %s", err.Error())
	}

	response, err := SendRequest(Request{
		URL:          fmt.Sprintf("https://api.hubapi.com/contacts/v1/lists/%s?hapikey=%s", listID, c.apiKey),
		Method:       "POST",
		Body:         body,
		OkStatusCode: 200,
	})

	return response, err
}

// DeleteList remove a Hubspot List given the List ID
func (c *Client) DeleteList(listID string) (Response, error) {
	response, err := SendRequest(Request{
		URL:          fmt.Sprintf("https://api.hubapi.com/contacts/v1/lists/%s?hapikey=%s", listID, c.apiKey),
		Method:       "DELETE",
		OkStatusCode: 204,
	})

	return response, err
}
