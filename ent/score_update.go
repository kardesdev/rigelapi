// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/activity"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/score"
	"github.com/vmkevv/rigelapi/ent/student"
)

// ScoreUpdate is the builder for updating Score entities.
type ScoreUpdate struct {
	config
	hooks    []Hook
	mutation *ScoreMutation
}

// Where appends a list predicates to the ScoreUpdate builder.
func (su *ScoreUpdate) Where(ps ...predicate.Score) *ScoreUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetPoints sets the "points" field.
func (su *ScoreUpdate) SetPoints(i int) *ScoreUpdate {
	su.mutation.ResetPoints()
	su.mutation.SetPoints(i)
	return su
}

// AddPoints adds i to the "points" field.
func (su *ScoreUpdate) AddPoints(i int) *ScoreUpdate {
	su.mutation.AddPoints(i)
	return su
}

// SetActivityID sets the "activity" edge to the Activity entity by ID.
func (su *ScoreUpdate) SetActivityID(id string) *ScoreUpdate {
	su.mutation.SetActivityID(id)
	return su
}

// SetNillableActivityID sets the "activity" edge to the Activity entity by ID if the given value is not nil.
func (su *ScoreUpdate) SetNillableActivityID(id *string) *ScoreUpdate {
	if id != nil {
		su = su.SetActivityID(*id)
	}
	return su
}

// SetActivity sets the "activity" edge to the Activity entity.
func (su *ScoreUpdate) SetActivity(a *Activity) *ScoreUpdate {
	return su.SetActivityID(a.ID)
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (su *ScoreUpdate) SetStudentID(id string) *ScoreUpdate {
	su.mutation.SetStudentID(id)
	return su
}

// SetNillableStudentID sets the "student" edge to the Student entity by ID if the given value is not nil.
func (su *ScoreUpdate) SetNillableStudentID(id *string) *ScoreUpdate {
	if id != nil {
		su = su.SetStudentID(*id)
	}
	return su
}

// SetStudent sets the "student" edge to the Student entity.
func (su *ScoreUpdate) SetStudent(s *Student) *ScoreUpdate {
	return su.SetStudentID(s.ID)
}

// Mutation returns the ScoreMutation object of the builder.
func (su *ScoreUpdate) Mutation() *ScoreMutation {
	return su.mutation
}

// ClearActivity clears the "activity" edge to the Activity entity.
func (su *ScoreUpdate) ClearActivity() *ScoreUpdate {
	su.mutation.ClearActivity()
	return su
}

// ClearStudent clears the "student" edge to the Student entity.
func (su *ScoreUpdate) ClearStudent() *ScoreUpdate {
	su.mutation.ClearStudent()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ScoreUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ScoreUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ScoreUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ScoreUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *ScoreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   score.Table,
			Columns: score.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: score.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Points(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: score.FieldPoints,
		})
	}
	if value, ok := su.mutation.AddedPoints(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: score.FieldPoints,
		})
	}
	if su.mutation.ActivityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.ActivityTable,
			Columns: []string{score.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: activity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ActivityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.ActivityTable,
			Columns: []string{score.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: activity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.StudentTable,
			Columns: []string{score.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: student.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.StudentTable,
			Columns: []string{score.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{score.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ScoreUpdateOne is the builder for updating a single Score entity.
type ScoreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ScoreMutation
}

// SetPoints sets the "points" field.
func (suo *ScoreUpdateOne) SetPoints(i int) *ScoreUpdateOne {
	suo.mutation.ResetPoints()
	suo.mutation.SetPoints(i)
	return suo
}

// AddPoints adds i to the "points" field.
func (suo *ScoreUpdateOne) AddPoints(i int) *ScoreUpdateOne {
	suo.mutation.AddPoints(i)
	return suo
}

// SetActivityID sets the "activity" edge to the Activity entity by ID.
func (suo *ScoreUpdateOne) SetActivityID(id string) *ScoreUpdateOne {
	suo.mutation.SetActivityID(id)
	return suo
}

// SetNillableActivityID sets the "activity" edge to the Activity entity by ID if the given value is not nil.
func (suo *ScoreUpdateOne) SetNillableActivityID(id *string) *ScoreUpdateOne {
	if id != nil {
		suo = suo.SetActivityID(*id)
	}
	return suo
}

// SetActivity sets the "activity" edge to the Activity entity.
func (suo *ScoreUpdateOne) SetActivity(a *Activity) *ScoreUpdateOne {
	return suo.SetActivityID(a.ID)
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (suo *ScoreUpdateOne) SetStudentID(id string) *ScoreUpdateOne {
	suo.mutation.SetStudentID(id)
	return suo
}

// SetNillableStudentID sets the "student" edge to the Student entity by ID if the given value is not nil.
func (suo *ScoreUpdateOne) SetNillableStudentID(id *string) *ScoreUpdateOne {
	if id != nil {
		suo = suo.SetStudentID(*id)
	}
	return suo
}

// SetStudent sets the "student" edge to the Student entity.
func (suo *ScoreUpdateOne) SetStudent(s *Student) *ScoreUpdateOne {
	return suo.SetStudentID(s.ID)
}

// Mutation returns the ScoreMutation object of the builder.
func (suo *ScoreUpdateOne) Mutation() *ScoreMutation {
	return suo.mutation
}

// ClearActivity clears the "activity" edge to the Activity entity.
func (suo *ScoreUpdateOne) ClearActivity() *ScoreUpdateOne {
	suo.mutation.ClearActivity()
	return suo
}

// ClearStudent clears the "student" edge to the Student entity.
func (suo *ScoreUpdateOne) ClearStudent() *ScoreUpdateOne {
	suo.mutation.ClearStudent()
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ScoreUpdateOne) Select(field string, fields ...string) *ScoreUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Score entity.
func (suo *ScoreUpdateOne) Save(ctx context.Context) (*Score, error) {
	var (
		err  error
		node *Score
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Score)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ScoreMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ScoreUpdateOne) SaveX(ctx context.Context) *Score {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ScoreUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ScoreUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *ScoreUpdateOne) sqlSave(ctx context.Context) (_node *Score, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   score.Table,
			Columns: score.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: score.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Score.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, score.FieldID)
		for _, f := range fields {
			if !score.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != score.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Points(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: score.FieldPoints,
		})
	}
	if value, ok := suo.mutation.AddedPoints(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: score.FieldPoints,
		})
	}
	if suo.mutation.ActivityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.ActivityTable,
			Columns: []string{score.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: activity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ActivityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.ActivityTable,
			Columns: []string{score.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: activity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.StudentTable,
			Columns: []string{score.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: student.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.StudentTable,
			Columns: []string{score.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Score{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{score.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}