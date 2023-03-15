package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type VirtualUser struct {
	ent.Schema
}

func (VirtualUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("user").GoType(Uname("")),
		field.String("password").GoType(Password("")),
	}
}

func (VirtualUser) Edges() []ent.Edge {
	return nil
}
