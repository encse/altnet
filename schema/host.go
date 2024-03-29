package schema

import (
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type HostName string

func (h HostName) ToUpper() string {
	return strings.ToUpper(string(h))
}

type Host struct {
	ent.Schema
}

func (Host) Fields() []ent.Field {

	return []ent.Field{
		field.String("name").GoType(HostName("")).Unique().Immutable(),
		field.Enum("type").Values("bbs", "uucp", "mil"),
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
	return []ent.Edge{
		edge.To("services", TcpService.Type),
		edge.To("virtualusers", VirtualUser.Type),
		edge.To("hackers", User.Type),
	}
}
