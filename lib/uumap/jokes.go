package uumap

import (
	"context"

	"entgo.io/ent/dialect/sql"
)

func (n Network) Joke(ctx context.Context) (string, error) {

	joke, err := n.client.Joke.
		Query().
		Order(func(s *sql.Selector) {
			s.OrderBy("RANDOM()")
		}).
		First(ctx)

	if err != nil {
		return "", err
	}
	return joke.Body, nil
}
