// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/attendance"
	"github.com/vmkevv/rigelapi/ent/attendanceday"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/student"
)

// AttendanceQuery is the builder for querying Attendance entities.
type AttendanceQuery struct {
	config
	limit             *int
	offset            *int
	unique            *bool
	order             []OrderFunc
	fields            []string
	predicates        []predicate.Attendance
	withAttendanceDay *AttendanceDayQuery
	withStudent       *StudentQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttendanceQuery builder.
func (aq *AttendanceQuery) Where(ps ...predicate.Attendance) *AttendanceQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AttendanceQuery) Limit(limit int) *AttendanceQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AttendanceQuery) Offset(offset int) *AttendanceQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AttendanceQuery) Unique(unique bool) *AttendanceQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AttendanceQuery) Order(o ...OrderFunc) *AttendanceQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryAttendanceDay chains the current query on the "attendanceDay" edge.
func (aq *AttendanceQuery) QueryAttendanceDay() *AttendanceDayQuery {
	query := &AttendanceDayQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attendance.Table, attendance.FieldID, selector),
			sqlgraph.To(attendanceday.Table, attendanceday.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attendance.AttendanceDayTable, attendance.AttendanceDayColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStudent chains the current query on the "student" edge.
func (aq *AttendanceQuery) QueryStudent() *StudentQuery {
	query := &StudentQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attendance.Table, attendance.FieldID, selector),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attendance.StudentTable, attendance.StudentColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Attendance entity from the query.
// Returns a *NotFoundError when no Attendance was found.
func (aq *AttendanceQuery) First(ctx context.Context) (*Attendance, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attendance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AttendanceQuery) FirstX(ctx context.Context) *Attendance {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Attendance ID from the query.
// Returns a *NotFoundError when no Attendance ID was found.
func (aq *AttendanceQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attendance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AttendanceQuery) FirstIDX(ctx context.Context) string {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Attendance entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Attendance entity is found.
// Returns a *NotFoundError when no Attendance entities are found.
func (aq *AttendanceQuery) Only(ctx context.Context) (*Attendance, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attendance.Label}
	default:
		return nil, &NotSingularError{attendance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AttendanceQuery) OnlyX(ctx context.Context) *Attendance {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Attendance ID in the query.
// Returns a *NotSingularError when more than one Attendance ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AttendanceQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attendance.Label}
	default:
		err = &NotSingularError{attendance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AttendanceQuery) OnlyIDX(ctx context.Context) string {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Attendances.
func (aq *AttendanceQuery) All(ctx context.Context) ([]*Attendance, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AttendanceQuery) AllX(ctx context.Context) []*Attendance {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Attendance IDs.
func (aq *AttendanceQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := aq.Select(attendance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AttendanceQuery) IDsX(ctx context.Context) []string {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AttendanceQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AttendanceQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AttendanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AttendanceQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttendanceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AttendanceQuery) Clone() *AttendanceQuery {
	if aq == nil {
		return nil
	}
	return &AttendanceQuery{
		config:            aq.config,
		limit:             aq.limit,
		offset:            aq.offset,
		order:             append([]OrderFunc{}, aq.order...),
		predicates:        append([]predicate.Attendance{}, aq.predicates...),
		withAttendanceDay: aq.withAttendanceDay.Clone(),
		withStudent:       aq.withStudent.Clone(),
		// clone intermediate query.
		sql:    aq.sql.Clone(),
		path:   aq.path,
		unique: aq.unique,
	}
}

// WithAttendanceDay tells the query-builder to eager-load the nodes that are connected to
// the "attendanceDay" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttendanceQuery) WithAttendanceDay(opts ...func(*AttendanceDayQuery)) *AttendanceQuery {
	query := &AttendanceDayQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withAttendanceDay = query
	return aq
}

// WithStudent tells the query-builder to eager-load the nodes that are connected to
// the "student" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttendanceQuery) WithStudent(opts ...func(*StudentQuery)) *AttendanceQuery {
	query := &StudentQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withStudent = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Value attendance.Value `json:"value,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Attendance.Query().
//		GroupBy(attendance.FieldValue).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AttendanceQuery) GroupBy(field string, fields ...string) *AttendanceGroupBy {
	grbuild := &AttendanceGroupBy{config: aq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	grbuild.label = attendance.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Value attendance.Value `json:"value,omitempty"`
//	}
//
//	client.Attendance.Query().
//		Select(attendance.FieldValue).
//		Scan(ctx, &v)
func (aq *AttendanceQuery) Select(fields ...string) *AttendanceSelect {
	aq.fields = append(aq.fields, fields...)
	selbuild := &AttendanceSelect{AttendanceQuery: aq}
	selbuild.label = attendance.Label
	selbuild.flds, selbuild.scan = &aq.fields, selbuild.Scan
	return selbuild
}

func (aq *AttendanceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !attendance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AttendanceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Attendance, error) {
	var (
		nodes       = []*Attendance{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withAttendanceDay != nil,
			aq.withStudent != nil,
		}
	)
	if aq.withAttendanceDay != nil || aq.withStudent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, attendance.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Attendance).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Attendance{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withAttendanceDay; query != nil {
		if err := aq.loadAttendanceDay(ctx, query, nodes, nil,
			func(n *Attendance, e *AttendanceDay) { n.Edges.AttendanceDay = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withStudent; query != nil {
		if err := aq.loadStudent(ctx, query, nodes, nil,
			func(n *Attendance, e *Student) { n.Edges.Student = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AttendanceQuery) loadAttendanceDay(ctx context.Context, query *AttendanceDayQuery, nodes []*Attendance, init func(*Attendance), assign func(*Attendance, *AttendanceDay)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Attendance)
	for i := range nodes {
		if nodes[i].attendance_day_attendances == nil {
			continue
		}
		fk := *nodes[i].attendance_day_attendances
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(attendanceday.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "attendance_day_attendances" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AttendanceQuery) loadStudent(ctx context.Context, query *StudentQuery, nodes []*Attendance, init func(*Attendance), assign func(*Attendance, *Student)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Attendance)
	for i := range nodes {
		if nodes[i].student_attendances == nil {
			continue
		}
		fk := *nodes[i].student_attendances
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(student.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "student_attendances" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *AttendanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.fields
	if len(aq.fields) > 0 {
		_spec.Unique = aq.unique != nil && *aq.unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AttendanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aq *AttendanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attendance.Table,
			Columns: attendance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: attendance.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attendance.FieldID)
		for i := range fields {
			if fields[i] != attendance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AttendanceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(attendance.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = attendance.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.unique != nil && *aq.unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AttendanceGroupBy is the group-by builder for Attendance entities.
type AttendanceGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AttendanceGroupBy) Aggregate(fns ...AggregateFunc) *AttendanceGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AttendanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

func (agb *AttendanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range agb.fields {
		if !attendance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AttendanceGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// AttendanceSelect is the builder for selecting fields of Attendance entities.
type AttendanceSelect struct {
	*AttendanceQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (as *AttendanceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AttendanceQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

func (as *AttendanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
