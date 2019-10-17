package hubspot

import (
	"fmt"
	"encoding/json"
)

type PageBody struct{
	Name string `json:"name"`
	Slug string `json:"slug"`
	Template string `json:"template_path"`
}

func (c *Client) SavePage() (Response, error) {
	req := PageBody{
		Name: "Test Event Page",
		Slug: "events/adsalkfjdask",
		Template: "landed/Templates/CX Event Page/event-page-template.html",
	}

	body, err := json.Marshal(req)

	if err != nil {
		return Response{}, fmt.Errorf("Invalid request: %s", err.Error())
	}

	response, err := SendRequest(Request{
		URL:			fmt.Sprintf("https://api.hubapi.com/content/api/v2/pages?hapikey=%s", c.apiKey),
		Method:			"POST",
		Body:			body,
		OkStatusCode: 	201,
	})

	return response, err
}