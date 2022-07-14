// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/bug/ent/organizationalunit"
	"entgo.io/bug/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrganizationalUnitDelete is the builder for deleting a OrganizationalUnit entity.
type OrganizationalUnitDelete struct {
	config
	hooks    []Hook
	mutation *OrganizationalUnitMutation
}

// Where appends a list predicates to the OrganizationalUnitDelete builder.
func (oud *OrganizationalUnitDelete) Where(ps ...predicate.OrganizationalUnit) *OrganizationalUnitDelete {
	oud.mutation.Where(ps...)
	return oud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oud *OrganizationalUnitDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oud.hooks) == 0 {
		affected, err = oud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrganizationalUnitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oud.mutation = mutation
			affected, err = oud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oud.hooks) - 1; i >= 0; i-- {
			if oud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (oud *OrganizationalUnitDelete) ExecX(ctx context.Context) int {
	n, err := oud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oud *OrganizationalUnitDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: organizationalunit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: organizationalunit.FieldID,
			},
		},
	}
	if ps := oud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, oud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// OrganizationalUnitDeleteOne is the builder for deleting a single OrganizationalUnit entity.
type OrganizationalUnitDeleteOne struct {
	oud *OrganizationalUnitDelete
}

// Exec executes the deletion query.
func (oudo *OrganizationalUnitDeleteOne) Exec(ctx context.Context) error {
	n, err := oudo.oud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{organizationalunit.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oudo *OrganizationalUnitDeleteOne) ExecX(ctx context.Context) {
	oudo.oud.ExecX(ctx)
}
