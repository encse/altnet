// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/encse/altnet/ent/joke"
	"github.com/encse/altnet/ent/predicate"
)

// JokeDelete is the builder for deleting a Joke entity.
type JokeDelete struct {
	config
	hooks    []Hook
	mutation *JokeMutation
}

// Where appends a list predicates to the JokeDelete builder.
func (jd *JokeDelete) Where(ps ...predicate.Joke) *JokeDelete {
	jd.mutation.Where(ps...)
	return jd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (jd *JokeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, JokeMutation](ctx, jd.sqlExec, jd.mutation, jd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (jd *JokeDelete) ExecX(ctx context.Context) int {
	n, err := jd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (jd *JokeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(joke.Table, sqlgraph.NewFieldSpec(joke.FieldID, field.TypeInt))
	if ps := jd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, jd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	jd.mutation.done = true
	return affected, err
}

// JokeDeleteOne is the builder for deleting a single Joke entity.
type JokeDeleteOne struct {
	jd *JokeDelete
}

// Where appends a list predicates to the JokeDelete builder.
func (jdo *JokeDeleteOne) Where(ps ...predicate.Joke) *JokeDeleteOne {
	jdo.jd.mutation.Where(ps...)
	return jdo
}

// Exec executes the deletion query.
func (jdo *JokeDeleteOne) Exec(ctx context.Context) error {
	n, err := jdo.jd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{joke.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (jdo *JokeDeleteOne) ExecX(ctx context.Context) {
	if err := jdo.Exec(ctx); err != nil {
		panic(err)
	}
}
