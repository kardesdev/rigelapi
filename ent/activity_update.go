// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/activity"
	"github.com/vmkevv/rigelapi/ent/area"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/score"
	"github.com/vmkevv/rigelapi/ent/scoresync"
)

// ActivityUpdate is the builder for updating Activity entities.
type ActivityUpdate struct {
	config
	hooks    []Hook
	mutation *ActivityMutation
}

// Where appends a list predicates to the ActivityUpdate builder.
func (au *ActivityUpdate) Where(ps ...predicate.Activity) *ActivityUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *ActivityUpdate) SetName(s string) *ActivityUpdate {
	au.mutation.SetName(s)
	return au
}

// SetDate sets the "date" field.
func (au *ActivityUpdate) SetDate(t time.Time) *ActivityUpdate {
	au.mutation.SetDate(t)
	return au
}

// AddScoreIDs adds the "scores" edge to the Score entity by IDs.
func (au *ActivityUpdate) AddScoreIDs(ids ...string) *ActivityUpdate {
	au.mutation.AddScoreIDs(ids...)
	return au
}

// AddScores adds the "scores" edges to the Score entity.
func (au *ActivityUpdate) AddScores(s ...*Score) *ActivityUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return au.AddScoreIDs(ids...)
}

// AddScoreSyncIDs adds the "scoreSyncs" edge to the ScoreSync entity by IDs.
func (au *ActivityUpdate) AddScoreSyncIDs(ids ...string) *ActivityUpdate {
	au.mutation.AddScoreSyncIDs(ids...)
	return au
}

// AddScoreSyncs adds the "scoreSyncs" edges to the ScoreSync entity.
func (au *ActivityUpdate) AddScoreSyncs(s ...*ScoreSync) *ActivityUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return au.AddScoreSyncIDs(ids...)
}

// SetAreaID sets the "area" edge to the Area entity by ID.
func (au *ActivityUpdate) SetAreaID(id string) *ActivityUpdate {
	au.mutation.SetAreaID(id)
	return au
}

// SetNillableAreaID sets the "area" edge to the Area entity by ID if the given value is not nil.
func (au *ActivityUpdate) SetNillableAreaID(id *string) *ActivityUpdate {
	if id != nil {
		au = au.SetAreaID(*id)
	}
	return au
}

// SetArea sets the "area" edge to the Area entity.
func (au *ActivityUpdate) SetArea(a *Area) *ActivityUpdate {
	return au.SetAreaID(a.ID)
}

// SetClassPeriodID sets the "classPeriod" edge to the ClassPeriod entity by ID.
func (au *ActivityUpdate) SetClassPeriodID(id string) *ActivityUpdate {
	au.mutation.SetClassPeriodID(id)
	return au
}

// SetNillableClassPeriodID sets the "classPeriod" edge to the ClassPeriod entity by ID if the given value is not nil.
func (au *ActivityUpdate) SetNillableClassPeriodID(id *string) *ActivityUpdate {
	if id != nil {
		au = au.SetClassPeriodID(*id)
	}
	return au
}

// SetClassPeriod sets the "classPeriod" edge to the ClassPeriod entity.
func (au *ActivityUpdate) SetClassPeriod(c *ClassPeriod) *ActivityUpdate {
	return au.SetClassPeriodID(c.ID)
}

// Mutation returns the ActivityMutation object of the builder.
func (au *ActivityUpdate) Mutation() *ActivityMutation {
	return au.mutation
}

// ClearScores clears all "scores" edges to the Score entity.
func (au *ActivityUpdate) ClearScores() *ActivityUpdate {
	au.mutation.ClearScores()
	return au
}

// RemoveScoreIDs removes the "scores" edge to Score entities by IDs.
func (au *ActivityUpdate) RemoveScoreIDs(ids ...string) *ActivityUpdate {
	au.mutation.RemoveScoreIDs(ids...)
	return au
}

// RemoveScores removes "scores" edges to Score entities.
func (au *ActivityUpdate) RemoveScores(s ...*Score) *ActivityUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return au.RemoveScoreIDs(ids...)
}

// ClearScoreSyncs clears all "scoreSyncs" edges to the ScoreSync entity.
func (au *ActivityUpdate) ClearScoreSyncs() *ActivityUpdate {
	au.mutation.ClearScoreSyncs()
	return au
}

// RemoveScoreSyncIDs removes the "scoreSyncs" edge to ScoreSync entities by IDs.
func (au *ActivityUpdate) RemoveScoreSyncIDs(ids ...string) *ActivityUpdate {
	au.mutation.RemoveScoreSyncIDs(ids...)
	return au
}

// RemoveScoreSyncs removes "scoreSyncs" edges to ScoreSync entities.
func (au *ActivityUpdate) RemoveScoreSyncs(s ...*ScoreSync) *ActivityUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return au.RemoveScoreSyncIDs(ids...)
}

// ClearArea clears the "area" edge to the Area entity.
func (au *ActivityUpdate) ClearArea() *ActivityUpdate {
	au.mutation.ClearArea()
	return au
}

// ClearClassPeriod clears the "classPeriod" edge to the ClassPeriod entity.
func (au *ActivityUpdate) ClearClassPeriod() *ActivityUpdate {
	au.mutation.ClearClassPeriod()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ActivityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *ActivityUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ActivityUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ActivityUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ActivityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   activity.Table,
			Columns: activity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: activity.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activity.FieldName,
		})
	}
	if value, ok := au.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: activity.FieldDate,
		})
	}
	if au.mutation.ScoresCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedScoresIDs(); len(nodes) > 0 && !au.mutation.ScoresCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ScoresIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.ScoreSyncsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedScoreSyncsIDs(); len(nodes) > 0 && !au.mutation.ScoreSyncsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ScoreSyncsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.AreaCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AreaIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.ClassPeriodCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ClassPeriodIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ActivityUpdateOne is the builder for updating a single Activity entity.
type ActivityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ActivityMutation
}

// SetName sets the "name" field.
func (auo *ActivityUpdateOne) SetName(s string) *ActivityUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetDate sets the "date" field.
func (auo *ActivityUpdateOne) SetDate(t time.Time) *ActivityUpdateOne {
	auo.mutation.SetDate(t)
	return auo
}

// AddScoreIDs adds the "scores" edge to the Score entity by IDs.
func (auo *ActivityUpdateOne) AddScoreIDs(ids ...string) *ActivityUpdateOne {
	auo.mutation.AddScoreIDs(ids...)
	return auo
}

// AddScores adds the "scores" edges to the Score entity.
func (auo *ActivityUpdateOne) AddScores(s ...*Score) *ActivityUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return auo.AddScoreIDs(ids...)
}

// AddScoreSyncIDs adds the "scoreSyncs" edge to the ScoreSync entity by IDs.
func (auo *ActivityUpdateOne) AddScoreSyncIDs(ids ...string) *ActivityUpdateOne {
	auo.mutation.AddScoreSyncIDs(ids...)
	return auo
}

// AddScoreSyncs adds the "scoreSyncs" edges to the ScoreSync entity.
func (auo *ActivityUpdateOne) AddScoreSyncs(s ...*ScoreSync) *ActivityUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return auo.AddScoreSyncIDs(ids...)
}

// SetAreaID sets the "area" edge to the Area entity by ID.
func (auo *ActivityUpdateOne) SetAreaID(id string) *ActivityUpdateOne {
	auo.mutation.SetAreaID(id)
	return auo
}

// SetNillableAreaID sets the "area" edge to the Area entity by ID if the given value is not nil.
func (auo *ActivityUpdateOne) SetNillableAreaID(id *string) *ActivityUpdateOne {
	if id != nil {
		auo = auo.SetAreaID(*id)
	}
	return auo
}

// SetArea sets the "area" edge to the Area entity.
func (auo *ActivityUpdateOne) SetArea(a *Area) *ActivityUpdateOne {
	return auo.SetAreaID(a.ID)
}

// SetClassPeriodID sets the "classPeriod" edge to the ClassPeriod entity by ID.
func (auo *ActivityUpdateOne) SetClassPeriodID(id string) *ActivityUpdateOne {
	auo.mutation.SetClassPeriodID(id)
	return auo
}

// SetNillableClassPeriodID sets the "classPeriod" edge to the ClassPeriod entity by ID if the given value is not nil.
func (auo *ActivityUpdateOne) SetNillableClassPeriodID(id *string) *ActivityUpdateOne {
	if id != nil {
		auo = auo.SetClassPeriodID(*id)
	}
	return auo
}

// SetClassPeriod sets the "classPeriod" edge to the ClassPeriod entity.
func (auo *ActivityUpdateOne) SetClassPeriod(c *ClassPeriod) *ActivityUpdateOne {
	return auo.SetClassPeriodID(c.ID)
}

// Mutation returns the ActivityMutation object of the builder.
func (auo *ActivityUpdateOne) Mutation() *ActivityMutation {
	return auo.mutation
}

// ClearScores clears all "scores" edges to the Score entity.
func (auo *ActivityUpdateOne) ClearScores() *ActivityUpdateOne {
	auo.mutation.ClearScores()
	return auo
}

// RemoveScoreIDs removes the "scores" edge to Score entities by IDs.
func (auo *ActivityUpdateOne) RemoveScoreIDs(ids ...string) *ActivityUpdateOne {
	auo.mutation.RemoveScoreIDs(ids...)
	return auo
}

// RemoveScores removes "scores" edges to Score entities.
func (auo *ActivityUpdateOne) RemoveScores(s ...*Score) *ActivityUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return auo.RemoveScoreIDs(ids...)
}

// ClearScoreSyncs clears all "scoreSyncs" edges to the ScoreSync entity.
func (auo *ActivityUpdateOne) ClearScoreSyncs() *ActivityUpdateOne {
	auo.mutation.ClearScoreSyncs()
	return auo
}

// RemoveScoreSyncIDs removes the "scoreSyncs" edge to ScoreSync entities by IDs.
func (auo *ActivityUpdateOne) RemoveScoreSyncIDs(ids ...string) *ActivityUpdateOne {
	auo.mutation.RemoveScoreSyncIDs(ids...)
	return auo
}

// RemoveScoreSyncs removes "scoreSyncs" edges to ScoreSync entities.
func (auo *ActivityUpdateOne) RemoveScoreSyncs(s ...*ScoreSync) *ActivityUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return auo.RemoveScoreSyncIDs(ids...)
}

// ClearArea clears the "area" edge to the Area entity.
func (auo *ActivityUpdateOne) ClearArea() *ActivityUpdateOne {
	auo.mutation.ClearArea()
	return auo
}

// ClearClassPeriod clears the "classPeriod" edge to the ClassPeriod entity.
func (auo *ActivityUpdateOne) ClearClassPeriod() *ActivityUpdateOne {
	auo.mutation.ClearClassPeriod()
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ActivityUpdateOne) Select(field string, fields ...string) *ActivityUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Activity entity.
func (auo *ActivityUpdateOne) Save(ctx context.Context) (*Activity, error) {
	var (
		err  error
		node *Activity
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (auo *ActivityUpdateOne) SaveX(ctx context.Context) *Activity {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ActivityUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ActivityUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ActivityUpdateOne) sqlSave(ctx context.Context) (_node *Activity, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   activity.Table,
			Columns: activity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: activity.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Activity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, activity.FieldID)
		for _, f := range fields {
			if !activity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != activity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activity.FieldName,
		})
	}
	if value, ok := auo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: activity.FieldDate,
		})
	}
	if auo.mutation.ScoresCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedScoresIDs(); len(nodes) > 0 && !auo.mutation.ScoresCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ScoresIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.ScoreSyncsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedScoreSyncsIDs(); len(nodes) > 0 && !auo.mutation.ScoreSyncsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ScoreSyncsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.AreaCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AreaIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.ClassPeriodCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ClassPeriodIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Activity{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
