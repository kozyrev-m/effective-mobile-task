// Package agent is used to receive additional data about a person and to add this data to a person.
package agent

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Client.
type Client struct {
	client *resty.Client
}

// New - client constructor.
func New() *Client {
	return &Client{client: resty.New()}
}

// getAge gets person age by name.
func (cl *Client) getAge(ctx context.Context, name string) (int, error) {
	res, err := cl.client.R().SetContext(ctx).Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
	if err != nil {
		return 0, err
	}

	params := &ageResponseBody{}
	if err := json.Unmarshal(res.Body(), params); err != nil {
		return 0, err
	}

	fmt.Printf("resp: %v\n", params)
	return params.Age, nil
}

// getGender gets person gender by name.
func (cl *Client) getGender(ctx context.Context, name string) (string, error) {
	res, err := cl.client.R().SetContext(ctx).Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name))
	if err != nil {
		return "", err
	}

	params := &genderResponseBody{}
	if err := json.Unmarshal(res.Body(), params); err != nil {
		return "", err
	}

	return params.Gender, nil
}

// getNationalize gets person nationalize by name.
func (cl *Client) getNationalize(ctx context.Context, name string) (string, error) {
	res, err := cl.client.R().SetContext(ctx).Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))
	if err != nil {
		return "", err
	}

	params := &nationalizeResponseBody{}
	if err := json.Unmarshal(res.Body(), params); err != nil {
		return "", err
	}

	fmt.Printf("resp: %v\n", params)

	country := selectCountry(params.Country)
	fmt.Printf("country: %v\n", country)

	return country, nil
}
