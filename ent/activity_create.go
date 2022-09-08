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
	"github.com/vmkevv/rigelapi/ent/area"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/score"
	"github.com/vmkevv/rigelapi/ent/scoresync"
)

// ActivityCreate is the builder for creating a Activity entity.
type ActivityCreate struct {
	config
	mutation *ActivityMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *ActivityCreate) SetName(s string) *ActivityCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetDate sets the "date" field.
func (ac *ActivityCreate) SetDate(t time.Time) *ActivityCreate {
	ac.mutation.SetDate(t)
	return ac
}

// SetID sets the "id" field.
func (ac *ActivityCreate) SetID(s string) *ActivityCreate {
	ac.mutation.SetID(s)
	return ac
}

// AddScoreIDs adds the "scores" edge to the Score entity by IDs.
func (ac *ActivityCreate) AddScoreIDs(ids ...string) *ActivityCreate {
	ac.mutation.AddScoreIDs(ids...)
	return ac
}

// AddScores adds the "scores" edges to the Score entity.
func (ac *ActivityCreate) AddScores(s ...*Score) *ActivityCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ac.AddScoreIDs(ids...)
}

// AddScoreSyncIDs adds the "scoreSyncs" edge to the ScoreSync entity by IDs.
func (ac *ActivityCreate) AddScoreSyncIDs(ids ...string) *ActivityCreate {
	ac.mutation.AddScoreSyncIDs(ids...)
	return ac
}

// AddScoreSyncs adds the "scoreSyncs" edges to the ScoreSync entity.
func (ac *ActivityCreate) AddScoreSyncs(s ...*ScoreSync) *ActivityCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ac.AddScoreSyncIDs(ids...)
}

// SetAreaID sets the "area" edge to the Area entity by ID.
func (ac *ActivityCreate) SetAreaID(id string) *ActivityCreate {
	ac.mutation.SetAreaID(id)
	return ac
}

// SetNillableAreaID sets the "area" edge to the Area entity by ID if the given value is not nil.
func (ac *ActivityCreate) SetNillableAreaID(id *string) *ActivityCreate {
	if id != nil {
		ac = ac.SetAreaID(*id)
	}
	return ac
}

// SetArea sets the "area" edge to the Area entity.
func (ac *ActivityCreate) SetArea(a *Area) *ActivityCreate {
	return ac.SetAreaID(a.ID)
}

// SetClassPeriodID sets the "classPeriod" edge to the ClassPeriod entity by ID.
func (ac *ActivityCreate) SetClassPeriodID(id string) *ActivityCreate {
	ac.mutation.SetClassPeriodID(id)
	return ac
}

// SetNillableClassPeriodID sets the "classPeriod" edge to the ClassPeriod entity by ID if the given value is not nil.
func (ac *ActivityCreate) SetNillableClassPeriodID(id *string) *ActivityCreate {
	if id != nil {
		ac = ac.SetClassPeriodID(*id)
	}
	return ac
}

// SetClassPeriod sets the "classPeriod" edge to the ClassPeriod entity.
func (ac *ActivityCreate) SetClassPeriod(c *ClassPeriod) *ActivityCreate {
	return ac.SetClassPeriodID(c.ID)
}

// Mutation returns the ActivityMutation object of the builder.
func (ac *ActivityCreate) Mutation() *ActivityMutation {
	return ac.mutation
}

// Save creates the Activity in the database.
func (ac *ActivityCreate) Save(ctx context.Context) (*Activity, error) {
	var (
		err  error
		node *Activity
	)
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Activity)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ActivityMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ActivityCreate) SaveX(ctx context.Context) *Activity {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ActivityCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ActivityCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ActivityCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Activity.name"`)}
	}
	if _, ok := ac.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "Activity.date"`)}
	}
	return nil
}

func (ac *ActivityCreate) sqlSave(ctx context.Context) (*Activity, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Activity.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (ac *ActivityCreate) createSpec() (*Activity, *sqlgraph.CreateSpec) {
	var (
		_node = &Activity{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: activity.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: activity.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activity.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: activity.FieldDate,
		})
		_node.Date = value
	}
	if nodes := ac.mutation.ScoresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activity.ScoresTable,
			Columns: []string{activity.ScoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: score.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ScoreSyncsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activity.ScoreSyncsTable,
			Columns: []string{activity.ScoreSyncsColumn},
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
	if nodes := ac.mutation.AreaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   activity.AreaTable,
			Columns: []string{activity.AreaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.area_activities = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ClassPeriodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   activity.ClassPeriodTable,
			Columns: []string{activity.ClassPeriodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: classperiod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.class_period_activities = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ActivityCreateBulk is the builder for creating many Activity entities in bulk.
type ActivityCreateBulk struct {
	config
	builders []*ActivityCreate
}

// Save creates the Activity entities in the database.
func (acb *ActivityCreateBulk) Save(ctx context.Context) ([]*Activity, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Activity, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ActivityMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ActivityCreateBulk) SaveX(ctx context.Context) []*Activity {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ActivityCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ActivityCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}