// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/activity"
	"github.com/vmkevv/rigelapi/ent/attendanceday"
	"github.com/vmkevv/rigelapi/ent/attendancedaysyncs"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/period"
	"github.com/vmkevv/rigelapi/ent/predicate"
)

// ClassPeriodQuery is the builder for querying ClassPeriod entities.
type ClassPeriodQuery struct {
	config
	limit                  *int
	offset                 *int
	unique                 *bool
	order                  []OrderFunc
	fields                 []string
	predicates             []predicate.ClassPeriod
	withAttendanceDays     *AttendanceDayQuery
	withAttendanceDaySyncs *AttendanceDaySyncsQuery
	withActivities         *ActivityQuery
	withClass              *ClassQuery
	withPeriod             *PeriodQuery
	withFKs                bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ClassPeriodQuery builder.
func (cpq *ClassPeriodQuery) Where(ps ...predicate.ClassPeriod) *ClassPeriodQuery {
	cpq.predicates = append(cpq.predicates, ps...)
	return cpq
}

// Limit adds a limit step to the query.
func (cpq *ClassPeriodQuery) Limit(limit int) *ClassPeriodQuery {
	cpq.limit = &limit
	return cpq
}

// Offset adds an offset step to the query.
func (cpq *ClassPeriodQuery) Offset(offset int) *ClassPeriodQuery {
	cpq.offset = &offset
	return cpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cpq *ClassPeriodQuery) Unique(unique bool) *ClassPeriodQuery {
	cpq.unique = &unique
	return cpq
}

// Order adds an order step to the query.
func (cpq *ClassPeriodQuery) Order(o ...OrderFunc) *ClassPeriodQuery {
	cpq.order = append(cpq.order, o...)
	return cpq
}

// QueryAttendanceDays chains the current query on the "attendanceDays" edge.
func (cpq *ClassPeriodQuery) QueryAttendanceDays() *AttendanceDayQuery {
	query := &AttendanceDayQuery{config: cpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(classperiod.Table, classperiod.FieldID, selector),
			sqlgraph.To(attendanceday.Table, attendanceday.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, classperiod.AttendanceDaysTable, classperiod.AttendanceDaysColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAttendanceDaySyncs chains the current query on the "attendanceDaySyncs" edge.
func (cpq *ClassPeriodQuery) QueryAttendanceDaySyncs() *AttendanceDaySyncsQuery {
	query := &AttendanceDaySyncsQuery{config: cpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(classperiod.Table, classperiod.FieldID, selector),
			sqlgraph.To(attendancedaysyncs.Table, attendancedaysyncs.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, classperiod.AttendanceDaySyncsTable, classperiod.AttendanceDaySyncsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryActivities chains the current query on the "activities" edge.
func (cpq *ClassPeriodQuery) QueryActivities() *ActivityQuery {
	query := &ActivityQuery{config: cpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(classperiod.Table, classperiod.FieldID, selector),
			sqlgraph.To(activity.Table, activity.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, classperiod.ActivitiesTable, classperiod.ActivitiesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryClass chains the current query on the "class" edge.
func (cpq *ClassPeriodQuery) QueryClass() *ClassQuery {
	query := &ClassQuery{config: cpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(classperiod.Table, classperiod.FieldID, selector),
			sqlgraph.To(class.Table, class.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, classperiod.ClassTable, classperiod.ClassColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPeriod chains the current query on the "period" edge.
func (cpq *ClassPeriodQuery) QueryPeriod() *PeriodQuery {
	query := &PeriodQuery{config: cpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(classperiod.Table, classperiod.FieldID, selector),
			sqlgraph.To(period.Table, period.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, classperiod.PeriodTable, classperiod.PeriodColumn),
		)
		fromU = sqlgraph.SetNeighbors(cpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ClassPeriod entity from the query.
// Returns a *NotFoundError when no ClassPeriod was found.
func (cpq *ClassPeriodQuery) First(ctx context.Context) (*ClassPeriod, error) {
	nodes, err := cpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{classperiod.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cpq *ClassPeriodQuery) FirstX(ctx context.Context) *ClassPeriod {
	node, err := cpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ClassPeriod ID from the query.
// Returns a *NotFoundError when no ClassPeriod ID was found.
func (cpq *ClassPeriodQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{classperiod.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cpq *ClassPeriodQuery) FirstIDX(ctx context.Context) string {
	id, err := cpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ClassPeriod entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ClassPeriod entity is found.
// Returns a *NotFoundError when no ClassPeriod entities are found.
func (cpq *ClassPeriodQuery) Only(ctx context.Context) (*ClassPeriod, error) {
	nodes, err := cpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{classperiod.Label}
	default:
		return nil, &NotSingularError{classperiod.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cpq *ClassPeriodQuery) OnlyX(ctx context.Context) *ClassPeriod {
	node, err := cpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ClassPeriod ID in the query.
// Returns a *NotSingularError when more than one ClassPeriod ID is found.
// Returns a *NotFoundError when no entities are found.
func (cpq *ClassPeriodQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{classperiod.Label}
	default:
		err = &NotSingularError{classperiod.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cpq *ClassPeriodQuery) OnlyIDX(ctx context.Context) string {
	id, err := cpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ClassPeriods.
func (cpq *ClassPeriodQuery) All(ctx context.Context) ([]*ClassPeriod, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cpq *ClassPeriodQuery) AllX(ctx context.Context) []*ClassPeriod {
	nodes, err := cpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ClassPeriod IDs.
func (cpq *ClassPeriodQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := cpq.Select(classperiod.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cpq *ClassPeriodQuery) IDsX(ctx context.Context) []string {
	ids, err := cpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cpq *ClassPeriodQuery) Count(ctx context.Context) (int, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cpq *ClassPeriodQuery) CountX(ctx context.Context) int {
	count, err := cpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cpq *ClassPeriodQuery) Exist(ctx context.Context) (bool, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cpq *ClassPeriodQuery) ExistX(ctx context.Context) bool {
	exist, err := cpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ClassPeriodQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cpq *ClassPeriodQuery) Clone() *ClassPeriodQuery {
	if cpq == nil {
		return nil
	}
	return &ClassPeriodQuery{
		config:                 cpq.config,
		limit:                  cpq.limit,
		offset:                 cpq.offset,
		order:                  append([]OrderFunc{}, cpq.order...),
		predicates:             append([]predicate.ClassPeriod{}, cpq.predicates...),
		withAttendanceDays:     cpq.withAttendanceDays.Clone(),
		withAttendanceDaySyncs: cpq.withAttendanceDaySyncs.Clone(),
		withActivities:         cpq.withActivities.Clone(),
		withClass:              cpq.withClass.Clone(),
		withPeriod:             cpq.withPeriod.Clone(),
		// clone intermediate query.
		sql:    cpq.sql.Clone(),
		path:   cpq.path,
		unique: cpq.unique,
	}
}

// WithAttendanceDays tells the query-builder to eager-load the nodes that are connected to
// the "attendanceDays" edge. The optional arguments are used to configure the query builder of the edge.
func (cpq *ClassPeriodQuery) WithAttendanceDays(opts ...func(*AttendanceDayQuery)) *ClassPeriodQuery {
	query := &AttendanceDayQuery{config: cpq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpq.withAttendanceDays = query
	return cpq
}

// WithAttendanceDaySyncs tells the query-builder to eager-load the nodes that are connected to
// the "attendanceDaySyncs" edge. The optional arguments are used to configure the query builder of the edge.
func (cpq *ClassPeriodQuery) WithAttendanceDaySyncs(opts ...func(*AttendanceDaySyncsQuery)) *ClassPeriodQuery {
	query := &AttendanceDaySyncsQuery{config: cpq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpq.withAttendanceDaySyncs = query
	return cpq
}

// WithActivities tells the query-builder to eager-load the nodes that are connected to
// the "activities" edge. The optional arguments are used to configure the query builder of the edge.
func (cpq *ClassPeriodQuery) WithActivities(opts ...func(*ActivityQuery)) *ClassPeriodQuery {
	query := &ActivityQuery{config: cpq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpq.withActivities = query
	return cpq
}

// WithClass tells the query-builder to eager-load the nodes that are connected to
// the "class" edge. The optional arguments are used to configure the query builder of the edge.
func (cpq *ClassPeriodQuery) WithClass(opts ...func(*ClassQuery)) *ClassPeriodQuery {
	query := &ClassQuery{config: cpq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpq.withClass = query
	return cpq
}

// WithPeriod tells the query-builder to eager-load the nodes that are connected to
// the "period" edge. The optional arguments are used to configure the query builder of the edge.
func (cpq *ClassPeriodQuery) WithPeriod(opts ...func(*PeriodQuery)) *ClassPeriodQuery {
	query := &PeriodQuery{config: cpq.config}
	for _, opt := range opts {
		opt(query)
	}
	cpq.withPeriod = query
	return cpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Start time.Time `json:"start,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ClassPeriod.Query().
//		GroupBy(classperiod.FieldStart).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cpq *ClassPeriodQuery) GroupBy(field string, fields ...string) *ClassPeriodGroupBy {
	grbuild := &ClassPeriodGroupBy{config: cpq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cpq.sqlQuery(ctx), nil
	}
	grbuild.label = classperiod.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Start time.Time `json:"start,omitempty"`
//	}
//
//	client.ClassPeriod.Query().
//		Select(classperiod.FieldStart).
//		Scan(ctx, &v)
func (cpq *ClassPeriodQuery) Select(fields ...string) *ClassPeriodSelect {
	cpq.fields = append(cpq.fields, fields...)
	selbuild := &ClassPeriodSelect{ClassPeriodQuery: cpq}
	selbuild.label = classperiod.Label
	selbuild.flds, selbuild.scan = &cpq.fields, selbuild.Scan
	return selbuild
}

func (cpq *ClassPeriodQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cpq.fields {
		if !classperiod.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cpq.path != nil {
		prev, err := cpq.path(ctx)
		if err != nil {
			return err
		}
		cpq.sql = prev
	}
	return nil
}

func (cpq *ClassPeriodQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ClassPeriod, error) {
	var (
		nodes       = []*ClassPeriod{}
		withFKs     = cpq.withFKs
		_spec       = cpq.querySpec()
		loadedTypes = [5]bool{
			cpq.withAttendanceDays != nil,
			cpq.withAttendanceDaySyncs != nil,
			cpq.withActivities != nil,
			cpq.withClass != nil,
			cpq.withPeriod != nil,
		}
	)
	if cpq.withClass != nil || cpq.withPeriod != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, classperiod.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ClassPeriod).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ClassPeriod{config: cpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cpq.withAttendanceDays; query != nil {
		if err := cpq.loadAttendanceDays(ctx, query, nodes,
			func(n *ClassPeriod) { n.Edges.AttendanceDays = []*AttendanceDay{} },
			func(n *ClassPeriod, e *AttendanceDay) { n.Edges.AttendanceDays = append(n.Edges.AttendanceDays, e) }); err != nil {
			return nil, err
		}
	}
	if query := cpq.withAttendanceDaySyncs; query != nil {
		if err := cpq.loadAttendanceDaySyncs(ctx, query, nodes,
			func(n *ClassPeriod) { n.Edges.AttendanceDaySyncs = []*AttendanceDaySyncs{} },
			func(n *ClassPeriod, e *AttendanceDaySyncs) {
				n.Edges.AttendanceDaySyncs = append(n.Edges.AttendanceDaySyncs, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := cpq.withActivities; query != nil {
		if err := cpq.loadActivities(ctx, query, nodes,
			func(n *ClassPeriod) { n.Edges.Activities = []*Activity{} },
			func(n *ClassPeriod, e *Activity) { n.Edges.Activities = append(n.Edges.Activities, e) }); err != nil {
			return nil, err
		}
	}
	if query := cpq.withClass; query != nil {
		if err := cpq.loadClass(ctx, query, nodes, nil,
			func(n *ClassPeriod, e *Class) { n.Edges.Class = e }); err != nil {
			return nil, err
		}
	}
	if query := cpq.withPeriod; query != nil {
		if err := cpq.loadPeriod(ctx, query, nodes, nil,
			func(n *ClassPeriod, e *Period) { n.Edges.Period = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cpq *ClassPeriodQuery) loadAttendanceDays(ctx context.Context, query *AttendanceDayQuery, nodes []*ClassPeriod, init func(*ClassPeriod), assign func(*ClassPeriod, *AttendanceDay)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*ClassPeriod)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AttendanceDay(func(s *sql.Selector) {
		s.Where(sql.InValues(classperiod.AttendanceDaysColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.class_period_attendance_days
		if fk == nil {
			return fmt.Errorf(`foreign-key "class_period_attendance_days" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "class_period_attendance_days" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cpq *ClassPeriodQuery) loadAttendanceDaySyncs(ctx context.Context, query *AttendanceDaySyncsQuery, nodes []*ClassPeriod, init func(*ClassPeriod), assign func(*ClassPeriod, *AttendanceDaySyncs)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*ClassPeriod)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.InValues(classperiod.AttendanceDaySyncsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.class_period_attendance_day_syncs
		if fk == nil {
			return fmt.Errorf(`foreign-key "class_period_attendance_day_syncs" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "class_period_attendance_day_syncs" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cpq *ClassPeriodQuery) loadActivities(ctx context.Context, query *ActivityQuery, nodes []*ClassPeriod, init func(*ClassPeriod), assign func(*ClassPeriod, *Activity)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*ClassPeriod)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.InValues(classperiod.ActivitiesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.class_period_activities
		if fk == nil {
			return fmt.Errorf(`foreign-key "class_period_activities" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "class_period_activities" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cpq *ClassPeriodQuery) loadClass(ctx context.Context, query *ClassQuery, nodes []*ClassPeriod, init func(*ClassPeriod), assign func(*ClassPeriod, *Class)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ClassPeriod)
	for i := range nodes {
		if nodes[i].class_class_periods == nil {
			continue
		}
		fk := *nodes[i].class_class_periods
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(class.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "class_class_periods" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cpq *ClassPeriodQuery) loadPeriod(ctx context.Context, query *PeriodQuery, nodes []*ClassPeriod, init func(*ClassPeriod), assign func(*ClassPeriod, *Period)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ClassPeriod)
	for i := range nodes {
		if nodes[i].period_class_periods == nil {
			continue
		}
		fk := *nodes[i].period_class_periods
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(period.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "period_class_periods" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cpq *ClassPeriodQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cpq.querySpec()
	_spec.Node.Columns = cpq.fields
	if len(cpq.fields) > 0 {
		_spec.Unique = cpq.unique != nil && *cpq.unique
	}
	return sqlgraph.CountNodes(ctx, cpq.driver, _spec)
}

func (cpq *ClassPeriodQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cpq *ClassPeriodQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   classperiod.Table,
			Columns: classperiod.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: classperiod.FieldID,
			},
		},
		From:   cpq.sql,
		Unique: true,
	}
	if unique := cpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, classperiod.FieldID)
		for i := range fields {
			if fields[i] != classperiod.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cpq *ClassPeriodQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cpq.driver.Dialect())
	t1 := builder.Table(classperiod.Table)
	columns := cpq.fields
	if len(columns) == 0 {
		columns = classperiod.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cpq.sql != nil {
		selector = cpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cpq.unique != nil && *cpq.unique {
		selector.Distinct()
	}
	for _, p := range cpq.predicates {
		p(selector)
	}
	for _, p := range cpq.order {
		p(selector)
	}
	if offset := cpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ClassPeriodGroupBy is the group-by builder for ClassPeriod entities.
type ClassPeriodGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cpgb *ClassPeriodGroupBy) Aggregate(fns ...AggregateFunc) *ClassPeriodGroupBy {
	cpgb.fns = append(cpgb.fns, fns...)
	return cpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cpgb *ClassPeriodGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cpgb.path(ctx)
	if err != nil {
		return err
	}
	cpgb.sql = query
	return cpgb.sqlScan(ctx, v)
}

func (cpgb *ClassPeriodGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cpgb.fields {
		if !classperiod.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cpgb *ClassPeriodGroupBy) sqlQuery() *sql.Selector {
	selector := cpgb.sql.Select()
	aggregation := make([]string, 0, len(cpgb.fns))
	for _, fn := range cpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cpgb.fields)+len(cpgb.fns))
		for _, f := range cpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cpgb.fields...)...)
}

// ClassPeriodSelect is the builder for selecting fields of ClassPeriod entities.
type ClassPeriodSelect struct {
	*ClassPeriodQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cps *ClassPeriodSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cps.prepareQuery(ctx); err != nil {
		return err
	}
	cps.sql = cps.ClassPeriodQuery.sqlQuery(ctx)
	return cps.sqlScan(ctx, v)
}

func (cps *ClassPeriodSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cps.sql.Query()
	if err := cps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
