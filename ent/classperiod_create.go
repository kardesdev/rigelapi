// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/activity"
	"github.com/vmkevv/rigelapi/ent/activitysync"
	"github.com/vmkevv/rigelapi/ent/attendanceday"
	"github.com/vmkevv/rigelapi/ent/attendancedaysyncs"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/period"
)

// ClassPeriodCreate is the builder for creating a ClassPeriod entity.
type ClassPeriodCreate struct {
	config
	mutation *ClassPeriodMutation
	hooks    []Hook
}

// SetStart sets the "start" field.
func (cpc *ClassPeriodCreate) SetStart(t time.Time) *ClassPeriodCreate {
	cpc.mutation.SetStart(t)
	return cpc
}

// SetEnd sets the "end" field.
func (cpc *ClassPeriodCreate) SetEnd(t time.Time) *ClassPeriodCreate {
	cpc.mutation.SetEnd(t)
	return cpc
}

// SetFinished sets the "finished" field.
func (cpc *ClassPeriodCreate) SetFinished(b bool) *ClassPeriodCreate {
	cpc.mutation.SetFinished(b)
	return cpc
}

// SetID sets the "id" field.
func (cpc *ClassPeriodCreate) SetID(s string) *ClassPeriodCreate {
	cpc.mutation.SetID(s)
	return cpc
}

// AddAttendanceDayIDs adds the "attendanceDays" edge to the AttendanceDay entity by IDs.
func (cpc *ClassPeriodCreate) AddAttendanceDayIDs(ids ...string) *ClassPeriodCreate {
	cpc.mutation.AddAttendanceDayIDs(ids...)
	return cpc
}

// AddAttendanceDays adds the "attendanceDays" edges to the AttendanceDay entity.
func (cpc *ClassPeriodCreate) AddAttendanceDays(a ...*AttendanceDay) *ClassPeriodCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpc.AddAttendanceDayIDs(ids...)
}

// AddAttendanceDaySyncIDs adds the "attendanceDaySyncs" edge to the AttendanceDaySyncs entity by IDs.
func (cpc *ClassPeriodCreate) AddAttendanceDaySyncIDs(ids ...string) *ClassPeriodCreate {
	cpc.mutation.AddAttendanceDaySyncIDs(ids...)
	return cpc
}

// AddAttendanceDaySyncs adds the "attendanceDaySyncs" edges to the AttendanceDaySyncs entity.
func (cpc *ClassPeriodCreate) AddAttendanceDaySyncs(a ...*AttendanceDaySyncs) *ClassPeriodCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpc.AddAttendanceDaySyncIDs(ids...)
}

// AddActivityIDs adds the "activities" edge to the Activity entity by IDs.
func (cpc *ClassPeriodCreate) AddActivityIDs(ids ...string) *ClassPeriodCreate {
	cpc.mutation.AddActivityIDs(ids...)
	return cpc
}

// AddActivities adds the "activities" edges to the Activity entity.
func (cpc *ClassPeriodCreate) AddActivities(a ...*Activity) *ClassPeriodCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpc.AddActivityIDs(ids...)
}

// AddActivitySyncIDs adds the "activitySyncs" edge to the ActivitySync entity by IDs.
func (cpc *ClassPeriodCreate) AddActivitySyncIDs(ids ...string) *ClassPeriodCreate {
	cpc.mutation.AddActivitySyncIDs(ids...)
	return cpc
}

// AddActivitySyncs adds the "activitySyncs" edges to the ActivitySync entity.
func (cpc *ClassPeriodCreate) AddActivitySyncs(a ...*ActivitySync) *ClassPeriodCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpc.AddActivitySyncIDs(ids...)
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (cpc *ClassPeriodCreate) SetClassID(id string) *ClassPeriodCreate {
	cpc.mutation.SetClassID(id)
	return cpc
}

// SetNillableClassID sets the "class" edge to the Class entity by ID if the given value is not nil.
func (cpc *ClassPeriodCreate) SetNillableClassID(id *string) *ClassPeriodCreate {
	if id != nil {
		cpc = cpc.SetClassID(*id)
	}
	return cpc
}

// SetClass sets the "class" edge to the Class entity.
func (cpc *ClassPeriodCreate) SetClass(c *Class) *ClassPeriodCreate {
	return cpc.SetClassID(c.ID)
}

// SetPeriodID sets the "period" edge to the Period entity by ID.
func (cpc *ClassPeriodCreate) SetPeriodID(id string) *ClassPeriodCreate {
	cpc.mutation.SetPeriodID(id)
	return cpc
}

// SetNillablePeriodID sets the "period" edge to the Period entity by ID if the given value is not nil.
func (cpc *ClassPeriodCreate) SetNillablePeriodID(id *string) *ClassPeriodCreate {
	if id != nil {
		cpc = cpc.SetPeriodID(*id)
	}
	return cpc
}

// SetPeriod sets the "period" edge to the Period entity.
func (cpc *ClassPeriodCreate) SetPeriod(p *Period) *ClassPeriodCreate {
	return cpc.SetPeriodID(p.ID)
}

// Mutation returns the ClassPeriodMutation object of the builder.
func (cpc *ClassPeriodCreate) Mutation() *ClassPeriodMutation {
	return cpc.mutation
}

// Save creates the ClassPeriod in the database.
func (cpc *ClassPeriodCreate) Save(ctx context.Context) (*ClassPeriod, error) {
	var (
		err  error
		node *ClassPeriod
	)
	if len(cpc.hooks) == 0 {
		if err = cpc.check(); err != nil {
			return nil, err
		}
		node, err = cpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClassPeriodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cpc.check(); err != nil {
				return nil, err
			}
			cpc.mutation = mutation
			if node, err = cpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cpc.hooks) - 1; i >= 0; i-- {
			if cpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cpc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ClassPeriod)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ClassPeriodMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cpc *ClassPeriodCreate) SaveX(ctx context.Context) *ClassPeriod {
	v, err := cpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cpc *ClassPeriodCreate) Exec(ctx context.Context) error {
	_, err := cpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpc *ClassPeriodCreate) ExecX(ctx context.Context) {
	if err := cpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpc *ClassPeriodCreate) check() error {
	if _, ok := cpc.mutation.Start(); !ok {
		return &ValidationError{Name: "start", err: errors.New(`ent: missing required field "ClassPeriod.start"`)}
	}
	if _, ok := cpc.mutation.End(); !ok {
		return &ValidationError{Name: "end", err: errors.New(`ent: missing required field "ClassPeriod.end"`)}
	}
	if _, ok := cpc.mutation.Finished(); !ok {
		return &ValidationError{Name: "finished", err: errors.New(`ent: missing required field "ClassPeriod.finished"`)}
	}
	return nil
}

func (cpc *ClassPeriodCreate) sqlSave(ctx context.Context) (*ClassPeriod, error) {
	_node, _spec := cpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ClassPeriod.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (cpc *ClassPeriodCreate) createSpec() (*ClassPeriod, *sqlgraph.CreateSpec) {
	var (
		_node = &ClassPeriod{config: cpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: classperiod.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: classperiod.FieldID,
			},
		}
	)
	if id, ok := cpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cpc.mutation.Start(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: classperiod.FieldStart,
		})
		_node.Start = value
	}
	if value, ok := cpc.mutation.End(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: classperiod.FieldEnd,
		})
		_node.End = value
	}
	if value, ok := cpc.mutation.Finished(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: classperiod.FieldFinished,
		})
		_node.Finished = value
	}
	if nodes := cpc.mutation.AttendanceDaysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.AttendanceDaysTable,
			Columns: []string{classperiod.AttendanceDaysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: attendanceday.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cpc.mutation.AttendanceDaySyncsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.AttendanceDaySyncsTable,
			Columns: []string{classperiod.AttendanceDaySyncsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: attendancedaysyncs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cpc.mutation.ActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.ActivitiesTable,
			Columns: []string{classperiod.ActivitiesColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cpc.mutation.ActivitySyncsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.ActivitySyncsTable,
			Columns: []string{classperiod.ActivitySyncsColumn},
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
	if nodes := cpc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classperiod.ClassTable,
			Columns: []string{classperiod.ClassColumn},
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
		_node.class_class_periods = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cpc.mutation.PeriodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classperiod.PeriodTable,
			Columns: []string{classperiod.PeriodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: period.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.period_class_periods = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ClassPeriodCreateBulk is the builder for creating many ClassPeriod entities in bulk.
type ClassPeriodCreateBulk struct {
	config
	builders []*ClassPeriodCreate
}

// Save creates the ClassPeriod entities in the database.
func (cpcb *ClassPeriodCreateBulk) Save(ctx context.Context) ([]*ClassPeriod, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cpcb.builders))
	nodes := make([]*ClassPeriod, len(cpcb.builders))
	mutators := make([]Mutator, len(cpcb.builders))
	for i := range cpcb.builders {
		func(i int, root context.Context) {
			builder := cpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClassPeriodMutation)
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
					_, err = mutators[i+1].Mutate(root, cpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cpcb *ClassPeriodCreateBulk) SaveX(ctx context.Context) []*ClassPeriod {
	v, err := cpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cpcb *ClassPeriodCreateBulk) Exec(ctx context.Context) error {
	_, err := cpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpcb *ClassPeriodCreateBulk) ExecX(ctx context.Context) {
	if err := cpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
