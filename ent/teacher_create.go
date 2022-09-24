// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/activitysync"
	"github.com/vmkevv/rigelapi/ent/attendancesync"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiodsync"
	"github.com/vmkevv/rigelapi/ent/scoresync"
	"github.com/vmkevv/rigelapi/ent/studentsync"
	"github.com/vmkevv/rigelapi/ent/teacher"
)

// TeacherCreate is the builder for creating a Teacher entity.
type TeacherCreate struct {
	config
	mutation *TeacherMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TeacherCreate) SetName(s string) *TeacherCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetLastName sets the "last_name" field.
func (tc *TeacherCreate) SetLastName(s string) *TeacherCreate {
	tc.mutation.SetLastName(s)
	return tc
}

// SetEmail sets the "email" field.
func (tc *TeacherCreate) SetEmail(s string) *TeacherCreate {
	tc.mutation.SetEmail(s)
	return tc
}

// SetPassword sets the "password" field.
func (tc *TeacherCreate) SetPassword(s string) *TeacherCreate {
	tc.mutation.SetPassword(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TeacherCreate) SetID(s string) *TeacherCreate {
	tc.mutation.SetID(s)
	return tc
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (tc *TeacherCreate) AddClassIDs(ids ...string) *TeacherCreate {
	tc.mutation.AddClassIDs(ids...)
	return tc
}

// AddClasses adds the "classes" edges to the Class entity.
func (tc *TeacherCreate) AddClasses(c ...*Class) *TeacherCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tc.AddClassIDs(ids...)
}

// AddScoreSyncIDs adds the "scoreSyncs" edge to the ScoreSync entity by IDs.
func (tc *TeacherCreate) AddScoreSyncIDs(ids ...string) *TeacherCreate {
	tc.mutation.AddScoreSyncIDs(ids...)
	return tc
}

// AddScoreSyncs adds the "scoreSyncs" edges to the ScoreSync entity.
func (tc *TeacherCreate) AddScoreSyncs(s ...*ScoreSync) *TeacherCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tc.AddScoreSyncIDs(ids...)
}

// AddStudentSyncIDs adds the "studentSyncs" edge to the StudentSync entity by IDs.
func (tc *TeacherCreate) AddStudentSyncIDs(ids ...string) *TeacherCreate {
	tc.mutation.AddStudentSyncIDs(ids...)
	return tc
}

// AddStudentSyncs adds the "studentSyncs" edges to the StudentSync entity.
func (tc *TeacherCreate) AddStudentSyncs(s ...*StudentSync) *TeacherCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tc.AddStudentSyncIDs(ids...)
}

// AddActivitySyncIDs adds the "activitySyncs" edge to the ActivitySync entity by IDs.
func (tc *TeacherCreate) AddActivitySyncIDs(ids ...string) *TeacherCreate {
	tc.mutation.AddActivitySyncIDs(ids...)
	return tc
}

// AddActivitySyncs adds the "activitySyncs" edges to the ActivitySync entity.
func (tc *TeacherCreate) AddActivitySyncs(a ...*ActivitySync) *TeacherCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tc.AddActivitySyncIDs(ids...)
}

// AddAttendanceSyncIDs adds the "attendanceSyncs" edge to the AttendanceSync entity by IDs.
func (tc *TeacherCreate) AddAttendanceSyncIDs(ids ...string) *TeacherCreate {
	tc.mutation.AddAttendanceSyncIDs(ids...)
	return tc
}

// AddAttendanceSyncs adds the "attendanceSyncs" edges to the AttendanceSync entity.
func (tc *TeacherCreate) AddAttendanceSyncs(a ...*AttendanceSync) *TeacherCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tc.AddAttendanceSyncIDs(ids...)
}

// AddClassPeriodSyncIDs adds the "classPeriodSyncs" edge to the ClassPeriodSync entity by IDs.
func (tc *TeacherCreate) AddClassPeriodSyncIDs(ids ...string) *TeacherCreate {
	tc.mutation.AddClassPeriodSyncIDs(ids...)
	return tc
}

// AddClassPeriodSyncs adds the "classPeriodSyncs" edges to the ClassPeriodSync entity.
func (tc *TeacherCreate) AddClassPeriodSyncs(c ...*ClassPeriodSync) *TeacherCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tc.AddClassPeriodSyncIDs(ids...)
}

// Mutation returns the TeacherMutation object of the builder.
func (tc *TeacherCreate) Mutation() *TeacherMutation {
	return tc.mutation
}

// Save creates the Teacher in the database.
func (tc *TeacherCreate) Save(ctx context.Context) (*Teacher, error) {
	var (
		err  error
		node *Teacher
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeacherMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Teacher)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TeacherMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TeacherCreate) SaveX(ctx context.Context) *Teacher {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TeacherCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TeacherCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TeacherCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Teacher.name"`)}
	}
	if _, ok := tc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "Teacher.last_name"`)}
	}
	if _, ok := tc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Teacher.email"`)}
	}
	if _, ok := tc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Teacher.password"`)}
	}
	return nil
}

func (tc *TeacherCreate) sqlSave(ctx context.Context) (*Teacher, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Teacher.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (tc *TeacherCreate) createSpec() (*Teacher, *sqlgraph.CreateSpec) {
	var (
		_node = &Teacher{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: teacher.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: teacher.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teacher.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teacher.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := tc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teacher.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := tc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teacher.FieldPassword,
		})
		_node.Password = value
	}
	if nodes := tc.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ClassesTable,
			Columns: []string{teacher.ClassesColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ScoreSyncsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ScoreSyncsTable,
			Columns: []string{teacher.ScoreSyncsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: scoresync.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.StudentSyncsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.StudentSyncsTable,
			Columns: []string{teacher.StudentSyncsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: studentsync.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ActivitySyncsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ActivitySyncsTable,
			Columns: []string{teacher.ActivitySyncsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: activitysync.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.AttendanceSyncsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.AttendanceSyncsTable,
			Columns: []string{teacher.AttendanceSyncsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: attendancesync.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ClassPeriodSyncsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ClassPeriodSyncsTable,
			Columns: []string{teacher.ClassPeriodSyncsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: classperiodsync.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TeacherCreateBulk is the builder for creating many Teacher entities in bulk.
type TeacherCreateBulk struct {
	config
	builders []*TeacherCreate
}

// Save creates the Teacher entities in the database.
func (tcb *TeacherCreateBulk) Save(ctx context.Context) ([]*Teacher, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Teacher, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeacherMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TeacherCreateBulk) SaveX(ctx context.Context) []*Teacher {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TeacherCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TeacherCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
