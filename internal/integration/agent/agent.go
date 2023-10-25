// Package agent is used to receive additional data about a person and to add this data to a person.
package agent

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/kozyrev-m/effective-mobile-task/internal/entities"
)

// Client.
type Client struct {
	client *resty.Client
}

// New - client constructor.
func New() *Client {
	return &Client{client: resty.New()}
}

// ReceiveAndSet is used to get additional person data and to set this data to a person.
func (cl *Client) ReceiveAndSet(ctx context.Context, person entities.Person) (*entities.Person, error) {
	age, err := cl.getAge(ctx, *person.Name)
	if err != nil {
		return nil, err
	}
	gender, err := cl.getGender(ctx, *person.Name)
	if err != nil {
		return nil, err
	}
	nationalize, err := cl.getNationalize(ctx, *person.Name)
	if err != nil {
		return nil, err
	}

	// set received data to person.
	*person.Age = age
	*person.Gender = gender
	*person.Nationality = nationalize

	return &person, nil
}

// getAge gets person age by name.
func (cl *Client) getAge(ctx context.Context, name string) (int, error) {

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
		params := &ageResponseBody{}
		res, err := cl.client.R().Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
		if err != nil {
			return 0, err
		}

		if err := json.Unmarshal(res.Body(), params); err != nil {
			return 0, err
		}

		fmt.Printf("resp: %v\n", params)
		return params.Age, nil
	}

}

// getGender gets person gender by name.
func (cl *Client) getGender(ctx context.Context, name string) (string, error) {

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		res, err := cl.client.R().Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name))
		if err != nil {
			return "", err
		}

		params := &genderResponseBody{}
		if err := json.Unmarshal(res.Body(), params); err != nil {
			return "", err
		}

		fmt.Printf("resp: %v\n", params)

		return params.Gender, nil
	}
}

// getNationalize gets person nationalize by name.
func (cl *Client) getNationalize(ctx context.Context, name string) (string, error) {

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		res, err := cl.client.R().Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))
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
}
