package agent

import (
	"context"

	"github.com/kozyrev-m/effective-mobile-task/internal/entities"
	"golang.org/x/sync/errgroup"
)

type ResultAge struct {
	Age int
	Err error
}

type ResultGender struct {
	Gender string
	Err    error
}

type ResultNationality struct {
	Nationality string
	Err         error
}

// ReceiveAndSet is used to get additional person data and to set this data to a person.
func (cl *Client) ReceiveAndSet(ctx context.Context, person entities.Person) (*entities.Person, error) {
	chAge := cl.asyncRequestAge(ctx, *person.Name)
	chGender := cl.asyncRequestGender(ctx, *person.Name)
	chNationality := cl.asyncRequestNationality(ctx, *person.Name)

	g := new(errgroup.Group)

	g.Go(func() error {
		return setAge(ctx, chAge, person.Age)
	})

	g.Go(func() error {
		return setGender(ctx, chGender, person.Gender)
	})

	g.Go(func() error {
		return setNationality(ctx, chNationality, person.Nationality)
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &person, nil
}

func (cl *Client) asyncRequestAge(ctx context.Context, name string) chan ResultAge {
	resCh := make(chan ResultAge)

	go func() {
		defer close(resCh)

		age, err := cl.getAge(ctx, name)

		resCh <- ResultAge{Age: age, Err: err}

		return
	}()

	return resCh
}

func (cl *Client) asyncRequestGender(ctx context.Context, name string) chan ResultGender {
	resCh := make(chan ResultGender)

	go func() {
		defer close(resCh)

		gender, err := cl.getGender(ctx, name)

		resCh <- ResultGender{Gender: gender, Err: err}
	}()

	return resCh
}

func (cl *Client) asyncRequestNationality(ctx context.Context, name string) chan ResultNationality {
	resCh := make(chan ResultNationality)

	go func() {
		defer close(resCh)

		nationality, err := cl.getNationalize(ctx, name)

		resCh <- ResultNationality{Nationality: nationality, Err: err}
	}()

	return resCh
}

func setAge(ctx context.Context, chAge chan ResultAge, age *int) error {
	resAge := <-chAge
	if resAge.Err != nil {
		return resAge.Err
	}
	*age = resAge.Age

	return nil
}

func setGender(ctx context.Context, chGender chan ResultGender, gender *string) error {
	resGender := <-chGender
	if resGender.Err != nil {
		return resGender.Err
	}
	*gender = resGender.Gender

	return nil
}

func setNationality(ctx context.Context, chNationality chan ResultNationality, nationality *string) error {
	resNationality := <-chNationality
	if resNationality.Err != nil {
		return resNationality.Err
	}
	*nationality = resNationality.Nationality

	return nil
}
