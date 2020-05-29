package hubspot

import (
	"fmt"
)

func (c *Client) EnrollContact(contactEmail string, workflowID string) (Response, error) {
	response, err := SendRequest(Request{
		URL:          fmt.Sprintf("https://api.hubapi.com/automation/v2/workflows/%s/enrollments/contacts/%s?hapikey=%s", workflowID, contactEmail, c.apiKey),
		Method:       "POST",
		Body:         nil,
		OkStatusCode: 204,
	})

	return response, err
}
