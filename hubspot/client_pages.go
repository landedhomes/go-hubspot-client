package hubspot

import (
	"fmt"
	"encoding/json"
)

// PageBody is the request body for POST/PUT requests on Hubspot Page API calls
type PageBody struct{
	Campaign string `json:"campaign,omitempty"`
	CampaignName string `json:"campaign_name,omitempty"`
	FooterHTML string `json:"footer_html,omitempty"`
	HeadHTML string `json:"head_html,omitempty"`
	IsDraft bool `json:"is_draft,omitempty"`
	MetaDescription string `json:"meta_description,omitempty"`
	MetaKeywords string `json:"meta_keywords,omitempty"`
	Name string `json:"name,omitempty"`
	Title string `json:"html_title,omitempty"`
	Password string `json:"password,omitempty"`
	PublishDate int64 `json:"publish_date,omitempty"`
	PublishImmediately bool `json:"publish_immediately,omitempty"`
	Slug string `json:"slug,omitempty"`
	Subcategory string `json:"subcategory,omitempty"`
	WidgetContainers string `json:"widget_containers,omitempty"`
	Widgets string `json:"widgets,omitempty"`
	Template string `json:"template_path,omitempty"`
}

// SavePage sends the user PageBody as a request to the Hubspot SavePage API
func (c *Client) SavePage(req *PageBody) (Response, error) {
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