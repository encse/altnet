// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/encse/altnet/ent/predicate"
	"github.com/encse/altnet/ent/schema"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// User applies equality check predicate on the "user" field. It's identical to UserEQ.
func User(v schema.Uname) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldEQ(FieldUser, vc))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v schema.PasswordHash) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldEQ(FieldPassword, vc))
}

// LastLogin applies equality check predicate on the "last_login" field. It's identical to LastLoginEQ.
func LastLogin(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastLogin, v))
}

// LastLoginAttempt applies equality check predicate on the "last_login_attempt" field. It's identical to LastLoginAttemptEQ.
func LastLoginAttempt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastLoginAttempt, v))
}

// UserEQ applies the EQ predicate on the "user" field.
func UserEQ(v schema.Uname) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldEQ(FieldUser, vc))
}

// UserNEQ applies the NEQ predicate on the "user" field.
func UserNEQ(v schema.Uname) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldNEQ(FieldUser, vc))
}

// UserIn applies the In predicate on the "user" field.
func UserIn(vs ...schema.Uname) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.User(sql.FieldIn(FieldUser, v...))
}

// UserNotIn applies the NotIn predicate on the "user" field.
func UserNotIn(vs ...schema.Uname) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.User(sql.FieldNotIn(FieldUser, v...))
}

// UserGT applies the GT predicate on the "user" field.
func UserGT(v schema.Uname) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldGT(FieldUser, vc))
}

// UserGTE applies the GTE predicate on the "user" field.
func UserGTE(v schema.Uname) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldGTE(FieldUser, vc))
}

// UserLT applies the LT predicate on the "user" field.
func UserLT(v schema.Uname) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldLT(FieldUser, vc))
}

// UserLTE applies the LTE predicate on the "user" field.
func UserLTE(v schema.Uname) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldLTE(FieldUser, vc))
}

// UserContains applies the Contains predicate on the "user" field.
func UserContains(v schema.Uname) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldContains(FieldUser, vc))
}

// UserHasPrefix applies the HasPrefix predicate on the "user" field.
func UserHasPrefix(v schema.Uname) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldHasPrefix(FieldUser, vc))
}

// UserHasSuffix applies the HasSuffix predicate on the "user" field.
func UserHasSuffix(v schema.Uname) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldHasSuffix(FieldUser, vc))
}

// UserEqualFold applies the EqualFold predicate on the "user" field.
func UserEqualFold(v schema.Uname) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldEqualFold(FieldUser, vc))
}

// UserContainsFold applies the ContainsFold predicate on the "user" field.
func UserContainsFold(v schema.Uname) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldContainsFold(FieldUser, vc))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v schema.PasswordHash) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldEQ(FieldPassword, vc))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v schema.PasswordHash) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldNEQ(FieldPassword, vc))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...schema.PasswordHash) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.User(sql.FieldIn(FieldPassword, v...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...schema.PasswordHash) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.User(sql.FieldNotIn(FieldPassword, v...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v schema.PasswordHash) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldGT(FieldPassword, vc))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v schema.PasswordHash) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldGTE(FieldPassword, vc))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v schema.PasswordHash) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldLT(FieldPassword, vc))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v schema.PasswordHash) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldLTE(FieldPassword, vc))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v schema.PasswordHash) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldContains(FieldPassword, vc))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v schema.PasswordHash) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldHasPrefix(FieldPassword, vc))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v schema.PasswordHash) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldHasSuffix(FieldPassword, vc))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v schema.PasswordHash) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldEqualFold(FieldPassword, vc))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v schema.PasswordHash) predicate.User {
	vc := string(v)
	return predicate.User(sql.FieldContainsFold(FieldPassword, vc))
}

// LastLoginEQ applies the EQ predicate on the "last_login" field.
func LastLoginEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastLogin, v))
}

// LastLoginNEQ applies the NEQ predicate on the "last_login" field.
func LastLoginNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldLastLogin, v))
}

// LastLoginIn applies the In predicate on the "last_login" field.
func LastLoginIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldLastLogin, vs...))
}

// LastLoginNotIn applies the NotIn predicate on the "last_login" field.
func LastLoginNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldLastLogin, vs...))
}

// LastLoginGT applies the GT predicate on the "last_login" field.
func LastLoginGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldLastLogin, v))
}

// LastLoginGTE applies the GTE predicate on the "last_login" field.
func LastLoginGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldLastLogin, v))
}

// LastLoginLT applies the LT predicate on the "last_login" field.
func LastLoginLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldLastLogin, v))
}

// LastLoginLTE applies the LTE predicate on the "last_login" field.
func LastLoginLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldLastLogin, v))
}

// LastLoginAttemptEQ applies the EQ predicate on the "last_login_attempt" field.
func LastLoginAttemptEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastLoginAttempt, v))
}

// LastLoginAttemptNEQ applies the NEQ predicate on the "last_login_attempt" field.
func LastLoginAttemptNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldLastLoginAttempt, v))
}

// LastLoginAttemptIn applies the In predicate on the "last_login_attempt" field.
func LastLoginAttemptIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldLastLoginAttempt, vs...))
}

// LastLoginAttemptNotIn applies the NotIn predicate on the "last_login_attempt" field.
func LastLoginAttemptNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldLastLoginAttempt, vs...))
}

// LastLoginAttemptGT applies the GT predicate on the "last_login_attempt" field.
func LastLoginAttemptGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldLastLoginAttempt, v))
}

// LastLoginAttemptGTE applies the GTE predicate on the "last_login_attempt" field.
func LastLoginAttemptGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldLastLoginAttempt, v))
}

// LastLoginAttemptLT applies the LT predicate on the "last_login_attempt" field.
func LastLoginAttemptLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldLastLoginAttempt, v))
}

// LastLoginAttemptLTE applies the LTE predicate on the "last_login_attempt" field.
func LastLoginAttemptLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldLastLoginAttempt, v))
}

// HasHosts applies the HasEdge predicate on the "hosts" edge.
func HasHosts() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, HostsTable, HostsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHostsWith applies the HasEdge predicate on the "hosts" edge with a given conditions (other predicates).
func HasHostsWith(preds ...predicate.Host) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HostsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, HostsTable, HostsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
