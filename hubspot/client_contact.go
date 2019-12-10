package hubspot

import (
	"fmt"
)

func (c *Client) GetContactByVID(vid string, properties []string) (Response, error) {
	var propertyQuery string
	
	for _, property := range properties {
		propertyQuery = propertyQuery + fmt.Sprintf("&property=%s", property) 
	}

	fmt.Printf("%s\n", propertyQuery)

	response, err := SendRequest(Request{
		URL:			fmt.Sprintf("https://api.hubapi.com/contacts/v1/contact/vid/%s/profile?hapikey=%s%s", vid, c.apiKey, propertyQuery),
		Method:			"GET",
		OkStatusCode:	200,
	})
	return response, err
}