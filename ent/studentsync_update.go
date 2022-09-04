// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/studentsync"
)

// StudentSyncUpdate is the builder for updating StudentSync entities.
type StudentSyncUpdate struct {
	config
	hooks    []Hook
	mutation *StudentSyncMutation
}

// Where appends a list predicates to the StudentSyncUpdate builder.
func (ssu *StudentSyncUpdate) Where(ps ...predicate.StudentSync) *StudentSyncUpdate {
	ssu.mutation.Where(ps...)
	return ssu
}

// SetLastSyncID sets the "last_sync_id" field.
func (ssu *StudentSyncUpdate) SetLastSyncID(s string) *StudentSyncUpdate {
	ssu.mutation.SetLastSyncID(s)
	return ssu
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (ssu *StudentSyncUpdate) SetClassID(id string) *StudentSyncUpdate {
	ssu.mutation.SetClassID(id)
	return ssu
}

// SetNillableClassID sets the "class" edge to the Class entity by ID if the given value is not nil.
func (ssu *StudentSyncUpdate) SetNillableClassID(id *string) *StudentSyncUpdate {
	if id != nil {
		ssu = ssu.SetClassID(*id)
	}
	return ssu
}

// SetClass sets the "class" edge to the Class entity.
func (ssu *StudentSyncUpdate) SetClass(c *Class) *StudentSyncUpdate {
	return ssu.SetClassID(c.ID)
}

// Mutation returns the StudentSyncMutation object of the builder.
func (ssu *StudentSyncUpdate) Mutation() *StudentSyncMutation {
	return ssu.mutation
}

// ClearClass clears the "class" edge to the Class entity.
func (ssu *StudentSyncUpdate) ClearClass() *StudentSyncUpdate {
	ssu.mutation.ClearClass()
	return ssu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ssu *StudentSyncUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ssu.hooks) == 0 {
		affected, err = ssu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudentSyncMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ssu.mutation = mutation
			affected, err = ssu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ssu.hooks) - 1; i >= 0; i-- {
			if ssu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ssu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ssu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ssu *StudentSyncUpdate) SaveX(ctx context.Context) int {
	affected, err := ssu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ssu *StudentSyncUpdate) Exec(ctx context.Context) error {
	_, err := ssu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssu *StudentSyncUpdate) ExecX(ctx context.Context) {
	if err := ssu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ssu *StudentSyncUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   studentsync.Table,
			Columns: studentsync.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: studentsync.FieldID,
			},
		},
	}
	if ps := ssu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ssu.mutation.LastSyncID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studentsync.FieldLastSyncID,
		})
	}
	if ssu.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studentsync.ClassTable,
			Columns: []string{studentsync.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ssu.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studentsync.ClassTable,
			Columns: []string{studentsync.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ssu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{studentsync.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// StudentSyncUpdateOne is the builder for updating a single StudentSync entity.
type StudentSyncUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudentSyncMutation
}

// SetLastSyncID sets the "last_sync_id" field.
func (ssuo *StudentSyncUpdateOne) SetLastSyncID(s string) *StudentSyncUpdateOne {
	ssuo.mutation.SetLastSyncID(s)
	return ssuo
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (ssuo *StudentSyncUpdateOne) SetClassID(id string) *StudentSyncUpdateOne {
	ssuo.mutation.SetClassID(id)
	return ssuo
}

// SetNillableClassID sets the "class" edge to the Class entity by ID if the given value is not nil.
func (ssuo *StudentSyncUpdateOne) SetNillableClassID(id *string) *StudentSyncUpdateOne {
	if id != nil {
		ssuo = ssuo.SetClassID(*id)
	}
	return ssuo
}

// SetClass sets the "class" edge to the Class entity.
func (ssuo *StudentSyncUpdateOne) SetClass(c *Class) *StudentSyncUpdateOne {
	return ssuo.SetClassID(c.ID)
}

// Mutation returns the StudentSyncMutation object of the builder.
func (ssuo *StudentSyncUpdateOne) Mutation() *StudentSyncMutation {
	return ssuo.mutation
}

// ClearClass clears the "class" edge to the Class entity.
func (ssuo *StudentSyncUpdateOne) ClearClass() *StudentSyncUpdateOne {
	ssuo.mutation.ClearClass()
	return ssuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ssuo *StudentSyncUpdateOne) Select(field string, fields ...string) *StudentSyncUpdateOne {
	ssuo.fields = append([]string{field}, fields...)
	return ssuo
}

// Save executes the query and returns the updated StudentSync entity.
func (ssuo *StudentSyncUpdateOne) Save(ctx context.Context) (*StudentSync, error) {
	var (
		err  error
		node *StudentSync
	)
	if len(ssuo.hooks) == 0 {
		node, err = ssuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudentSyncMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ssuo.mutation = mutation
			node, err = ssuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ssuo.hooks) - 1; i >= 0; i-- {
			if ssuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ssuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ssuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*StudentSync)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from StudentSyncMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ssuo *StudentSyncUpdateOne) SaveX(ctx context.Context) *StudentSync {
	node, err := ssuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ssuo *StudentSyncUpdateOne) Exec(ctx context.Context) error {
	_, err := ssuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssuo *StudentSyncUpdateOne) ExecX(ctx context.Context) {
	if err := ssuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ssuo *StudentSyncUpdateOne) sqlSave(ctx context.Context) (_node *StudentSync, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   studentsync.Table,
			Columns: studentsync.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: studentsync.FieldID,
			},
		},
	}
	id, ok := ssuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "StudentSync.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ssuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, studentsync.FieldID)
		for _, f := range fields {
			if !studentsync.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != studentsync.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ssuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ssuo.mutation.LastSyncID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: studentsync.FieldLastSyncID,
		})
	}
	if ssuo.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studentsync.ClassTable,
			Columns: []string{studentsync.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ssuo.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studentsync.ClassTable,
			Columns: []string{studentsync.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &StudentSync{config: ssuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ssuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{studentsync.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
