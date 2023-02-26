package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type HostName string

type Host struct {
	ent.Schema
}

func (Host) Fields() []ent.Field {

	return []ent.Field{
		field.String("name").GoType(HostName("")),
		field.String("entry").Default(""),
		field.String("machine_type").Default(""),
		field.String("organization").Default(""),
		field.String("contact").Default(""),
		field.String("contact_address").Default(""),
		field.String("country").Default(""),
		field.String("location").Default(""),
		field.String("geo_location").Default(""),
		field.JSON("phone", []PhoneNumber{}).Optional(),
		field.JSON("neighbours", []HostName{}).Optional(),
	}
}

func (Host) Edges() []ent.Edge {
	return nil
}
