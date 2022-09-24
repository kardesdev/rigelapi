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
	"github.com/vmkevv/rigelapi/ent/attendance"
	"github.com/vmkevv/rigelapi/ent/attendanceday"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/predicate"
)

// AttendanceDayUpdate is the builder for updating AttendanceDay entities.
type AttendanceDayUpdate struct {
	config
	hooks    []Hook
	mutation *AttendanceDayMutation
}

// Where appends a list predicates to the AttendanceDayUpdate builder.
func (adu *AttendanceDayUpdate) Where(ps ...predicate.AttendanceDay) *AttendanceDayUpdate {
	adu.mutation.Where(ps...)
	return adu
}

// SetDay sets the "day" field.
func (adu *AttendanceDayUpdate) SetDay(t time.Time) *AttendanceDayUpdate {
	adu.mutation.SetDay(t)
	return adu
}

// AddAttendanceIDs adds the "attendances" edge to the Attendance entity by IDs.
func (adu *AttendanceDayUpdate) AddAttendanceIDs(ids ...string) *AttendanceDayUpdate {
	adu.mutation.AddAttendanceIDs(ids...)
	return adu
}

// AddAttendances adds the "attendances" edges to the Attendance entity.
func (adu *AttendanceDayUpdate) AddAttendances(a ...*Attendance) *AttendanceDayUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return adu.AddAttendanceIDs(ids...)
}

// SetClassPeriodID sets the "classPeriod" edge to the ClassPeriod entity by ID.
func (adu *AttendanceDayUpdate) SetClassPeriodID(id string) *AttendanceDayUpdate {
	adu.mutation.SetClassPeriodID(id)
	return adu
}

// SetNillableClassPeriodID sets the "classPeriod" edge to the ClassPeriod entity by ID if the given value is not nil.
func (adu *AttendanceDayUpdate) SetNillableClassPeriodID(id *string) *AttendanceDayUpdate {
	if id != nil {
		adu = adu.SetClassPeriodID(*id)
	}
	return adu
}

// SetClassPeriod sets the "classPeriod" edge to the ClassPeriod entity.
func (adu *AttendanceDayUpdate) SetClassPeriod(c *ClassPeriod) *AttendanceDayUpdate {
	return adu.SetClassPeriodID(c.ID)
}

// Mutation returns the AttendanceDayMutation object of the builder.
func (adu *AttendanceDayUpdate) Mutation() *AttendanceDayMutation {
	return adu.mutation
}

// ClearAttendances clears all "attendances" edges to the Attendance entity.
func (adu *AttendanceDayUpdate) ClearAttendances() *AttendanceDayUpdate {
	adu.mutation.ClearAttendances()
	return adu
}

// RemoveAttendanceIDs removes the "attendances" edge to Attendance entities by IDs.
func (adu *AttendanceDayUpdate) RemoveAttendanceIDs(ids ...string) *AttendanceDayUpdate {
	adu.mutation.RemoveAttendanceIDs(ids...)
	return adu
}

// RemoveAttendances removes "attendances" edges to Attendance entities.
func (adu *AttendanceDayUpdate) RemoveAttendances(a ...*Attendance) *AttendanceDayUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return adu.RemoveAttendanceIDs(ids...)
}

// ClearClassPeriod clears the "classPeriod" edge to the ClassPeriod entity.
func (adu *AttendanceDayUpdate) ClearClassPeriod() *AttendanceDayUpdate {
	adu.mutation.ClearClassPeriod()
	return adu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (adu *AttendanceDayUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(adu.hooks) == 0 {
		affected, err = adu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttendanceDayMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			adu.mutation = mutation
			affected, err = adu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(adu.hooks) - 1; i >= 0; i-- {
			if adu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = adu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, adu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (adu *AttendanceDayUpdate) SaveX(ctx context.Context) int {
	affected, err := adu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (adu *AttendanceDayUpdate) Exec(ctx context.Context) error {
	_, err := adu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (adu *AttendanceDayUpdate) ExecX(ctx context.Context) {
	if err := adu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (adu *AttendanceDayUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attendanceday.Table,
			Columns: attendanceday.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: attendanceday.FieldID,
			},
		},
	}
	if ps := adu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := adu.mutation.Day(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: attendanceday.FieldDay,
		})
	}
	if adu.mutation.AttendancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attendanceday.AttendancesTable,
			Columns: []string{attendanceday.AttendancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: attendance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := adu.mutation.RemovedAttendancesIDs(); len(nodes) > 0 && !adu.mutation.AttendancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attendanceday.AttendancesTable,
			Columns: []string{attendanceday.AttendancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: attendance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := adu.mutation.AttendancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attendanceday.AttendancesTable,
			Columns: []string{attendanceday.AttendancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: attendance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if adu.mutation.ClassPeriodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendanceday.ClassPeriodTable,
			Columns: []string{attendanceday.ClassPeriodColumn},
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
	if nodes := adu.mutation.ClassPeriodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendanceday.ClassPeriodTable,
			Columns: []string{attendanceday.ClassPeriodColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, adu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attendanceday.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AttendanceDayUpdateOne is the builder for updating a single AttendanceDay entity.
type AttendanceDayUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttendanceDayMutation
}

// SetDay sets the "day" field.
func (aduo *AttendanceDayUpdateOne) SetDay(t time.Time) *AttendanceDayUpdateOne {
	aduo.mutation.SetDay(t)
	return aduo
}

// AddAttendanceIDs adds the "attendances" edge to the Attendance entity by IDs.
func (aduo *AttendanceDayUpdateOne) AddAttendanceIDs(ids ...string) *AttendanceDayUpdateOne {
	aduo.mutation.AddAttendanceIDs(ids...)
	return aduo
}

// AddAttendances adds the "attendances" edges to the Attendance entity.
func (aduo *AttendanceDayUpdateOne) AddAttendances(a ...*Attendance) *AttendanceDayUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aduo.AddAttendanceIDs(ids...)
}

// SetClassPeriodID sets the "classPeriod" edge to the ClassPeriod entity by ID.
func (aduo *AttendanceDayUpdateOne) SetClassPeriodID(id string) *AttendanceDayUpdateOne {
	aduo.mutation.SetClassPeriodID(id)
	return aduo
}

// SetNillableClassPeriodID sets the "classPeriod" edge to the ClassPeriod entity by ID if the given value is not nil.
func (aduo *AttendanceDayUpdateOne) SetNillableClassPeriodID(id *string) *AttendanceDayUpdateOne {
	if id != nil {
		aduo = aduo.SetClassPeriodID(*id)
	}
	return aduo
}

// SetClassPeriod sets the "classPeriod" edge to the ClassPeriod entity.
func (aduo *AttendanceDayUpdateOne) SetClassPeriod(c *ClassPeriod) *AttendanceDayUpdateOne {
	return aduo.SetClassPeriodID(c.ID)
}

// Mutation returns the AttendanceDayMutation object of the builder.
func (aduo *AttendanceDayUpdateOne) Mutation() *AttendanceDayMutation {
	return aduo.mutation
}

// ClearAttendances clears all "attendances" edges to the Attendance entity.
func (aduo *AttendanceDayUpdateOne) ClearAttendances() *AttendanceDayUpdateOne {
	aduo.mutation.ClearAttendances()
	return aduo
}

// RemoveAttendanceIDs removes the "attendances" edge to Attendance entities by IDs.
func (aduo *AttendanceDayUpdateOne) RemoveAttendanceIDs(ids ...string) *AttendanceDayUpdateOne {
	aduo.mutation.RemoveAttendanceIDs(ids...)
	return aduo
}

// RemoveAttendances removes "attendances" edges to Attendance entities.
func (aduo *AttendanceDayUpdateOne) RemoveAttendances(a ...*Attendance) *AttendanceDayUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aduo.RemoveAttendanceIDs(ids...)
}

// ClearClassPeriod clears the "classPeriod" edge to the ClassPeriod entity.
func (aduo *AttendanceDayUpdateOne) ClearClassPeriod() *AttendanceDayUpdateOne {
	aduo.mutation.ClearClassPeriod()
	return aduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aduo *AttendanceDayUpdateOne) Select(field string, fields ...string) *AttendanceDayUpdateOne {
	aduo.fields = append([]string{field}, fields...)
	return aduo
}

// Save executes the query and returns the updated AttendanceDay entity.
func (aduo *AttendanceDayUpdateOne) Save(ctx context.Context) (*AttendanceDay, error) {
	var (
		err  error
		node *AttendanceDay
	)
	if len(aduo.hooks) == 0 {
		node, err = aduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttendanceDayMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aduo.mutation = mutation
			node, err = aduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aduo.hooks) - 1; i >= 0; i-- {
			if aduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, aduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AttendanceDay)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AttendanceDayMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aduo *AttendanceDayUpdateOne) SaveX(ctx context.Context) *AttendanceDay {
	node, err := aduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aduo *AttendanceDayUpdateOne) Exec(ctx context.Context) error {
	_, err := aduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aduo *AttendanceDayUpdateOne) ExecX(ctx context.Context) {
	if err := aduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aduo *AttendanceDayUpdateOne) sqlSave(ctx context.Context) (_node *AttendanceDay, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attendanceday.Table,
			Columns: attendanceday.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: attendanceday.FieldID,
			},
		},
	}
	id, ok := aduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AttendanceDay.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attendanceday.FieldID)
		for _, f := range fields {
			if !attendanceday.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attendanceday.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aduo.mutation.Day(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: attendanceday.FieldDay,
		})
	}
	if aduo.mutation.AttendancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attendanceday.AttendancesTable,
			Columns: []string{attendanceday.AttendancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: attendance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aduo.mutation.RemovedAttendancesIDs(); len(nodes) > 0 && !aduo.mutation.AttendancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attendanceday.AttendancesTable,
			Columns: []string{attendanceday.AttendancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: attendance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aduo.mutation.AttendancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attendanceday.AttendancesTable,
			Columns: []string{attendanceday.AttendancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: attendance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aduo.mutation.ClassPeriodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendanceday.ClassPeriodTable,
			Columns: []string{attendanceday.ClassPeriodColumn},
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
	if nodes := aduo.mutation.ClassPeriodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendanceday.ClassPeriodTable,
			Columns: []string{attendanceday.ClassPeriodColumn},
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
	_node = &AttendanceDay{config: aduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attendanceday.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
