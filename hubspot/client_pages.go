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
	IsDraft string `json:"is_draft,omitempty"`
	MetaDescription string `json:"meta_description,omitempty"`
	MetaKeywords string `json:"meta_keywords,omitempty"`
	Name string `json:"name,omitempty"`
	Title string `json:"html_title,omitempty"`
	Password string `json:"password,omitempty"`
	PublishDate int64 `json:"publish_date,omitempty"`
	PublishImmediately string `json:"publish_immediately,omitempty"`
	Slug string `json:"slug,omitempty"`
	Subcategory string `json:"subcategory,omitempty"`
	WidgetContainers string `json:"widget_containers,omitempty"`
	Widgets string `json:"widgets,omitempty"`
	Template string `json:"template_path,omitempty"`
}

// PublishPageBody is the request body for the Hubspot Publish/Unpublish Page API
// Action takes in a string of either
// 1) push-buffer-live
// 2) schedule-publish
// 3) cancel-publish
type PublishPageBody struct{
	Action string `json:"action"`
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

// PublishPage publish a page according to the Action given in the request body
func (c *Client) PublishPage(req *PublishPageBody, id string) (Response, error) {
	body, err := json.Marshal(req)

	fmt.Println(string(body))

	if err != nil {
		return Response{}, fmt.Errorf("Invalid request: %s", err.Error())
	}

	response, err := SendRequest(Request{
		URL:			fmt.Sprintf("https://api.hubapi.com/content/api/v2/pages/%s/publish-action?hapikey=%s", id, c.apiKey),
		Method:			"POST",
		Body:			body,
		OkStatusCode:	204,
	})

	return response, err
}

// DeletePage deletes a page from Hubspot given the Page ID
func (c *Client) DeletePage(id string) (Response, error) {
	response, err := SendRequest(Request{
		URL:			fmt.Sprintf("https://api.hubspot.com/content/api/v2/pages/%s?hapikey=%s", id, c.apiKey),
		Method:			"DELETE",
		OkStatusCode:	204,
	})

	return response, err
}