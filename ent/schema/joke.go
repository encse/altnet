package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Joke struct {
	ent.Schema
}

func (Joke) Fields() []ent.Field {

	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("body"),
		field.String("category"),
	}
}
