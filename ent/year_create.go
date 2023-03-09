// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/area"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/period"
	"github.com/vmkevv/rigelapi/ent/subscription"
	"github.com/vmkevv/rigelapi/ent/year"
)

// YearCreate is the builder for creating a Year entity.
type YearCreate struct {
	config
	mutation *YearMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetValue sets the "value" field.
func (yc *YearCreate) SetValue(i int) *YearCreate {
	yc.mutation.SetValue(i)
	return yc
}

// SetID sets the "id" field.
func (yc *YearCreate) SetID(s string) *YearCreate {
	yc.mutation.SetID(s)
	return yc
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (yc *YearCreate) AddClassIDs(ids ...string) *YearCreate {
	yc.mutation.AddClassIDs(ids...)
	return yc
}

// AddClasses adds the "classes" edges to the Class entity.
func (yc *YearCreate) AddClasses(c ...*Class) *YearCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return yc.AddClassIDs(ids...)
}

// AddPeriodIDs adds the "periods" edge to the Period entity by IDs.
func (yc *YearCreate) AddPeriodIDs(ids ...string) *YearCreate {
	yc.mutation.AddPeriodIDs(ids...)
	return yc
}

// AddPeriods adds the "periods" edges to the Period entity.
func (yc *YearCreate) AddPeriods(p ...*Period) *YearCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return yc.AddPeriodIDs(ids...)
}

// AddAreaIDs adds the "areas" edge to the Area entity by IDs.
func (yc *YearCreate) AddAreaIDs(ids ...string) *YearCreate {
	yc.mutation.AddAreaIDs(ids...)
	return yc
}

// AddAreas adds the "areas" edges to the Area entity.
func (yc *YearCreate) AddAreas(a ...*Area) *YearCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return yc.AddAreaIDs(ids...)
}

// AddSubscriptionIDs adds the "subscriptions" edge to the Subscription entity by IDs.
func (yc *YearCreate) AddSubscriptionIDs(ids ...string) *YearCreate {
	yc.mutation.AddSubscriptionIDs(ids...)
	return yc
}

// AddSubscriptions adds the "subscriptions" edges to the Subscription entity.
func (yc *YearCreate) AddSubscriptions(s ...*Subscription) *YearCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return yc.AddSubscriptionIDs(ids...)
}

// Mutation returns the YearMutation object of the builder.
func (yc *YearCreate) Mutation() *YearMutation {
	return yc.mutation
}

// Save creates the Year in the database.
func (yc *YearCreate) Save(ctx context.Context) (*Year, error) {
	var (
		err  error
		node *Year
	)
	if len(yc.hooks) == 0 {
		if err = yc.check(); err != nil {
			return nil, err
		}
		node, err = yc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*YearMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = yc.check(); err != nil {
				return nil, err
			}
			yc.mutation = mutation
			if node, err = yc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(yc.hooks) - 1; i >= 0; i-- {
			if yc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = yc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, yc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Year)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from YearMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (yc *YearCreate) SaveX(ctx context.Context) *Year {
	v, err := yc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (yc *YearCreate) Exec(ctx context.Context) error {
	_, err := yc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (yc *YearCreate) ExecX(ctx context.Context) {
	if err := yc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (yc *YearCreate) check() error {
	if _, ok := yc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Year.value"`)}
	}
	return nil
}

func (yc *YearCreate) sqlSave(ctx context.Context) (*Year, error) {
	_node, _spec := yc.createSpec()
	if err := sqlgraph.CreateNode(ctx, yc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Year.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (yc *YearCreate) createSpec() (*Year, *sqlgraph.CreateSpec) {
	var (
		_node = &Year{config: yc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: year.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: year.FieldID,
			},
		}
	)
	_spec.OnConflict = yc.conflict
	if id, ok := yc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := yc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: year.FieldValue,
		})
		_node.Value = value
	}
	if nodes := yc.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.ClassesTable,
			Columns: []string{year.ClassesColumn},
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
	if nodes := yc.mutation.PeriodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.PeriodsTable,
			Columns: []string{year.PeriodsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := yc.mutation.AreasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.AreasTable,
			Columns: []string{year.AreasColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := yc.mutation.SubscriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   year.SubscriptionsTable,
			Columns: []string{year.SubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: subscription.FieldID,
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Year.Create().
//		SetValue(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.YearUpsert) {
//			SetValue(v+v).
//		}).
//		Exec(ctx)
func (yc *YearCreate) OnConflict(opts ...sql.ConflictOption) *YearUpsertOne {
	yc.conflict = opts
	return &YearUpsertOne{
		create: yc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Year.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (yc *YearCreate) OnConflictColumns(columns ...string) *YearUpsertOne {
	yc.conflict = append(yc.conflict, sql.ConflictColumns(columns...))
	return &YearUpsertOne{
		create: yc,
	}
}

type (
	// YearUpsertOne is the builder for "upsert"-ing
	//  one Year node.
	YearUpsertOne struct {
		create *YearCreate
	}

	// YearUpsert is the "OnConflict" setter.
	YearUpsert struct {
		*sql.UpdateSet
	}
)

// SetValue sets the "value" field.
func (u *YearUpsert) SetValue(v int) *YearUpsert {
	u.Set(year.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *YearUpsert) UpdateValue() *YearUpsert {
	u.SetExcluded(year.FieldValue)
	return u
}

// AddValue adds v to the "value" field.
func (u *YearUpsert) AddValue(v int) *YearUpsert {
	u.Add(year.FieldValue, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Year.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(year.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *YearUpsertOne) UpdateNewValues() *YearUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(year.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Year.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *YearUpsertOne) Ignore() *YearUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *YearUpsertOne) DoNothing() *YearUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the YearCreate.OnConflict
// documentation for more info.
func (u *YearUpsertOne) Update(set func(*YearUpsert)) *YearUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&YearUpsert{UpdateSet: update})
	}))
	return u
}

// SetValue sets the "value" field.
func (u *YearUpsertOne) SetValue(v int) *YearUpsertOne {
	return u.Update(func(s *YearUpsert) {
		s.SetValue(v)
	})
}

// AddValue adds v to the "value" field.
func (u *YearUpsertOne) AddValue(v int) *YearUpsertOne {
	return u.Update(func(s *YearUpsert) {
		s.AddValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *YearUpsertOne) UpdateValue() *YearUpsertOne {
	return u.Update(func(s *YearUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *YearUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for YearCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *YearUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *YearUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: YearUpsertOne.ID is not supported by MySQL driver. Use YearUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *YearUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// YearCreateBulk is the builder for creating many Year entities in bulk.
type YearCreateBulk struct {
	config
	builders []*YearCreate
	conflict []sql.ConflictOption
}

// Save creates the Year entities in the database.
func (ycb *YearCreateBulk) Save(ctx context.Context) ([]*Year, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ycb.builders))
	nodes := make([]*Year, len(ycb.builders))
	mutators := make([]Mutator, len(ycb.builders))
	for i := range ycb.builders {
		func(i int, root context.Context) {
			builder := ycb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*YearMutation)
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
					_, err = mutators[i+1].Mutate(root, ycb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ycb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ycb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ycb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ycb *YearCreateBulk) SaveX(ctx context.Context) []*Year {
	v, err := ycb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ycb *YearCreateBulk) Exec(ctx context.Context) error {
	_, err := ycb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ycb *YearCreateBulk) ExecX(ctx context.Context) {
	if err := ycb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Year.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.YearUpsert) {
//			SetValue(v+v).
//		}).
//		Exec(ctx)
func (ycb *YearCreateBulk) OnConflict(opts ...sql.ConflictOption) *YearUpsertBulk {
	ycb.conflict = opts
	return &YearUpsertBulk{
		create: ycb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Year.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ycb *YearCreateBulk) OnConflictColumns(columns ...string) *YearUpsertBulk {
	ycb.conflict = append(ycb.conflict, sql.ConflictColumns(columns...))
	return &YearUpsertBulk{
		create: ycb,
	}
}

// YearUpsertBulk is the builder for "upsert"-ing
// a bulk of Year nodes.
type YearUpsertBulk struct {
	create *YearCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Year.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(year.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *YearUpsertBulk) UpdateNewValues() *YearUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(year.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Year.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *YearUpsertBulk) Ignore() *YearUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *YearUpsertBulk) DoNothing() *YearUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the YearCreateBulk.OnConflict
// documentation for more info.
func (u *YearUpsertBulk) Update(set func(*YearUpsert)) *YearUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&YearUpsert{UpdateSet: update})
	}))
	return u
}

// SetValue sets the "value" field.
func (u *YearUpsertBulk) SetValue(v int) *YearUpsertBulk {
	return u.Update(func(s *YearUpsert) {
		s.SetValue(v)
	})
}

// AddValue adds v to the "value" field.
func (u *YearUpsertBulk) AddValue(v int) *YearUpsertBulk {
	return u.Update(func(s *YearUpsert) {
		s.AddValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *YearUpsertBulk) UpdateValue() *YearUpsertBulk {
	return u.Update(func(s *YearUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *YearUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the YearCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for YearCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *YearUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
