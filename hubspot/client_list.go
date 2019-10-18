package hubspot

import(
	"fmt"
	"encoding/json"
)

// ContactPropertyFilter is the Hubspot conditional logic use to generate Hubspot Contact list
type ContactPropertyFilter struct {
	Operator string `json:"operator"`
	Value string `json:"value"`
	Property string `json:"property"`
	Type string `json:"type"`
}

// ListBody is the request body structure for the Hubspot CreateList API
type ListBody struct {
	Name string `json:"name"`
	Dynamic bool `json:"dynamic,omitempty"`
	Filters [][]ContactPropertyFilter `json:"filters,omitempty"`
}

// GetContactsInList return Contacts associated to a given List ID
func (c *Client) GetContactsInList(listID string) (Response, error) {
	response, err := SendRequest(Request{
		URL:			fmt.Sprintf("https://api.hubapi.com/contacts/v1/lists/%s/contacts/all?hapikey=%s", listID, c.apiKey),
		Method:			"GET",
		OkStatusCode: 	200,
	})

	return response, err
}

// CreateList generate a new Hubspot List given the request body criterion
func (c *Client) CreateList(req *ListBody) (Response, error) {
	body, err := json.Marshal(req)

	if err != nil {
		return Response{}, fmt.Errorf("Invalid request: %s", err.Error())
	}

	fmt.Printf("%s\n", body)

	response, err := SendRequest(Request{
		URL:			fmt.Sprintf("https://api.hubapi.com/contacts/v1/lists?hapikey=%s", c.apiKey),
		Method:			"POST",
		Body:			body,
		OkStatusCode: 	200,
	})

	return response, err
}