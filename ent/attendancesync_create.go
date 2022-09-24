// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/attendancesync"
	"github.com/vmkevv/rigelapi/ent/teacher"
)

// AttendanceSyncCreate is the builder for creating a AttendanceSync entity.
type AttendanceSyncCreate struct {
	config
	mutation *AttendanceSyncMutation
	hooks    []Hook
}

// SetLastSyncID sets the "last_sync_id" field.
func (asc *AttendanceSyncCreate) SetLastSyncID(s string) *AttendanceSyncCreate {
	asc.mutation.SetLastSyncID(s)
	return asc
}

// SetID sets the "id" field.
func (asc *AttendanceSyncCreate) SetID(s string) *AttendanceSyncCreate {
	asc.mutation.SetID(s)
	return asc
}

// SetTeacherID sets the "teacher" edge to the Teacher entity by ID.
func (asc *AttendanceSyncCreate) SetTeacherID(id string) *AttendanceSyncCreate {
	asc.mutation.SetTeacherID(id)
	return asc
}

// SetNillableTeacherID sets the "teacher" edge to the Teacher entity by ID if the given value is not nil.
func (asc *AttendanceSyncCreate) SetNillableTeacherID(id *string) *AttendanceSyncCreate {
	if id != nil {
		asc = asc.SetTeacherID(*id)
	}
	return asc
}

// SetTeacher sets the "teacher" edge to the Teacher entity.
func (asc *AttendanceSyncCreate) SetTeacher(t *Teacher) *AttendanceSyncCreate {
	return asc.SetTeacherID(t.ID)
}

// Mutation returns the AttendanceSyncMutation object of the builder.
func (asc *AttendanceSyncCreate) Mutation() *AttendanceSyncMutation {
	return asc.mutation
}

// Save creates the AttendanceSync in the database.
func (asc *AttendanceSyncCreate) Save(ctx context.Context) (*AttendanceSync, error) {
	var (
		err  error
		node *AttendanceSync
	)
	if len(asc.hooks) == 0 {
		if err = asc.check(); err != nil {
			return nil, err
		}
		node, err = asc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttendanceSyncMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = asc.check(); err != nil {
				return nil, err
			}
			asc.mutation = mutation
			if node, err = asc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(asc.hooks) - 1; i >= 0; i-- {
			if asc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, asc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AttendanceSync)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AttendanceSyncMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (asc *AttendanceSyncCreate) SaveX(ctx context.Context) *AttendanceSync {
	v, err := asc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (asc *AttendanceSyncCreate) Exec(ctx context.Context) error {
	_, err := asc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asc *AttendanceSyncCreate) ExecX(ctx context.Context) {
	if err := asc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asc *AttendanceSyncCreate) check() error {
	if _, ok := asc.mutation.LastSyncID(); !ok {
		return &ValidationError{Name: "last_sync_id", err: errors.New(`ent: missing required field "AttendanceSync.last_sync_id"`)}
	}
	return nil
}

func (asc *AttendanceSyncCreate) sqlSave(ctx context.Context) (*AttendanceSync, error) {
	_node, _spec := asc.createSpec()
	if err := sqlgraph.CreateNode(ctx, asc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected AttendanceSync.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (asc *AttendanceSyncCreate) createSpec() (*AttendanceSync, *sqlgraph.CreateSpec) {
	var (
		_node = &AttendanceSync{config: asc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: attendancesync.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: attendancesync.FieldID,
			},
		}
	)
	if id, ok := asc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := asc.mutation.LastSyncID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: attendancesync.FieldLastSyncID,
		})
		_node.LastSyncID = value
	}
	if nodes := asc.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendancesync.TeacherTable,
			Columns: []string{attendancesync.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: teacher.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.teacher_attendance_syncs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AttendanceSyncCreateBulk is the builder for creating many AttendanceSync entities in bulk.
type AttendanceSyncCreateBulk struct {
	config
	builders []*AttendanceSyncCreate
}

// Save creates the AttendanceSync entities in the database.
func (ascb *AttendanceSyncCreateBulk) Save(ctx context.Context) ([]*AttendanceSync, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ascb.builders))
	nodes := make([]*AttendanceSync, len(ascb.builders))
	mutators := make([]Mutator, len(ascb.builders))
	for i := range ascb.builders {
		func(i int, root context.Context) {
			builder := ascb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttendanceSyncMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ascb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ascb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ascb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ascb *AttendanceSyncCreateBulk) SaveX(ctx context.Context) []*AttendanceSync {
	v, err := ascb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ascb *AttendanceSyncCreateBulk) Exec(ctx context.Context) error {
	_, err := ascb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ascb *AttendanceSyncCreateBulk) ExecX(ctx context.Context) {
	if err := ascb.Exec(ctx); err != nil {
		panic(err)
	}
}
