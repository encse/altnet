package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type TcpService struct {
	ent.Schema
}

func (TcpService) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Immutable(),
		field.Int("port").Unique().Immutable(),
		field.String("description").Immutable(),
	}
}

func (TcpService) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hosts", Host.Type).Ref("services"),
	}
}
