// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/attendancedaysyncs"
	"github.com/vmkevv/rigelapi/ent/predicate"
)

// AttendanceDaySyncsDelete is the builder for deleting a AttendanceDaySyncs entity.
type AttendanceDaySyncsDelete struct {
	config
	hooks    []Hook
	mutation *AttendanceDaySyncsMutation
}

// Where appends a list predicates to the AttendanceDaySyncsDelete builder.
func (adsd *AttendanceDaySyncsDelete) Where(ps ...predicate.AttendanceDaySyncs) *AttendanceDaySyncsDelete {
	adsd.mutation.Where(ps...)
	return adsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (adsd *AttendanceDaySyncsDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(adsd.hooks) == 0 {
		affected, err = adsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttendanceDaySyncsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			adsd.mutation = mutation
			affected, err = adsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(adsd.hooks) - 1; i >= 0; i-- {
			if adsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = adsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, adsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (adsd *AttendanceDaySyncsDelete) ExecX(ctx context.Context) int {
	n, err := adsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (adsd *AttendanceDaySyncsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: attendancedaysyncs.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: attendancedaysyncs.FieldID,
			},
		},
	}
	if ps := adsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, adsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AttendanceDaySyncsDeleteOne is the builder for deleting a single AttendanceDaySyncs entity.
type AttendanceDaySyncsDeleteOne struct {
	adsd *AttendanceDaySyncsDelete
}

// Exec executes the deletion query.
func (adsdo *AttendanceDaySyncsDeleteOne) Exec(ctx context.Context) error {
	n, err := adsdo.adsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{attendancedaysyncs.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (adsdo *AttendanceDaySyncsDeleteOne) ExecX(ctx context.Context) {
	adsdo.adsd.ExecX(ctx)
}
