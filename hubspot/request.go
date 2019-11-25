package hubspot

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"bytes"
)

// Request is the standard struct use for all HTTP calls
type Request struct {
	URL          string
	Method       string
	Body         []byte
	OkStatusCode int
}

// Response is the standard struct use for all HTTP calls response
type Response struct {
	Body       []byte
	StatusCode int
}

type HubspotError struct {
	Details		HubspotErrorDetails `json:"details`
}

type HubspotErrorDetails struct {
	ErrorType	string	`json:"errorType"`
}

// SendRequest is the helper function use for all HTTP calls
func SendRequest(r Request) (Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(r.Method, r.URL, bytes.NewBuffer(r.Body))
	if err != nil {
		return Response{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Response{}, err
	}

	if resp.StatusCode != r.OkStatusCode {
		var hubspotErrorResponse HubspotError
		err = json.Unmarshal([]byte(body), &hubspotErrorResponse)
		return Response{}, fmt.Errorf(hubspotErrorResponse.Details.ErrorType)
	}
	return Response{Body: body, StatusCode: resp.StatusCode}, nil
}