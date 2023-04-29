package schema

import (
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"golang.org/x/crypto/bcrypt"
)

type Uname string
type Password string
type PasswordHash string

func (u Uname) ToLower() Uname {
	return Uname(strings.ToLower(string(u)))
}

func (p Password) Hash() (PasswordHash, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return PasswordHash(hash), err
}

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user").GoType(Uname("")),
		field.String("password").GoType(PasswordHash("")),
		field.String("status").Optional().Default(""),
		field.Time("last_login").Optional().Nillable(),
		field.Time("last_login_attempt").Optional().Nillable(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hosts", Host.Type).Ref("hackers"),
	}
}
