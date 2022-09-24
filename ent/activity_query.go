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
	"github.com/vmkevv/rigelapi/ent/area"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/score"
)

// ActivityQuery is the builder for querying Activity entities.
type ActivityQuery struct {
	config
	limit           *int
	offset          *int
	unique          *bool
	order           []OrderFunc
	fields          []string
	predicates      []predicate.Activity
	withScores      *ScoreQuery
	withArea        *AreaQuery
	withClassPeriod *ClassPeriodQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ActivityQuery builder.
func (aq *ActivityQuery) Where(ps ...predicate.Activity) *ActivityQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *ActivityQuery) Limit(limit int) *ActivityQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *ActivityQuery) Offset(offset int) *ActivityQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *ActivityQuery) Unique(unique bool) *ActivityQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *ActivityQuery) Order(o ...OrderFunc) *ActivityQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryScores chains the current query on the "scores" edge.
func (aq *ActivityQuery) QueryScores() *ScoreQuery {
	query := &ScoreQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(activity.Table, activity.FieldID, selector),
			sqlgraph.To(score.Table, score.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, activity.ScoresTable, activity.ScoresColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryArea chains the current query on the "area" edge.
func (aq *ActivityQuery) QueryArea() *AreaQuery {
	query := &AreaQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(activity.Table, activity.FieldID, selector),
			sqlgraph.To(area.Table, area.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, activity.AreaTable, activity.AreaColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryClassPeriod chains the current query on the "classPeriod" edge.
func (aq *ActivityQuery) QueryClassPeriod() *ClassPeriodQuery {
	query := &ClassPeriodQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(activity.Table, activity.FieldID, selector),
			sqlgraph.To(classperiod.Table, classperiod.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, activity.ClassPeriodTable, activity.ClassPeriodColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Activity entity from the query.
// Returns a *NotFoundError when no Activity was found.
func (aq *ActivityQuery) First(ctx context.Context) (*Activity, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{activity.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *ActivityQuery) FirstX(ctx context.Context) *Activity {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Activity ID from the query.
// Returns a *NotFoundError when no Activity ID was found.
func (aq *ActivityQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{activity.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *ActivityQuery) FirstIDX(ctx context.Context) string {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Activity entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Activity entity is found.
// Returns a *NotFoundError when no Activity entities are found.
func (aq *ActivityQuery) Only(ctx context.Context) (*Activity, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{activity.Label}
	default:
		return nil, &NotSingularError{activity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *ActivityQuery) OnlyX(ctx context.Context) *Activity {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Activity ID in the query.
// Returns a *NotSingularError when more than one Activity ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *ActivityQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{activity.Label}
	default:
		err = &NotSingularError{activity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *ActivityQuery) OnlyIDX(ctx context.Context) string {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Activities.
func (aq *ActivityQuery) All(ctx context.Context) ([]*Activity, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *ActivityQuery) AllX(ctx context.Context) []*Activity {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Activity IDs.
func (aq *ActivityQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := aq.Select(activity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *ActivityQuery) IDsX(ctx context.Context) []string {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *ActivityQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *ActivityQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *ActivityQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *ActivityQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ActivityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *ActivityQuery) Clone() *ActivityQuery {
	if aq == nil {
		return nil
	}
	return &ActivityQuery{
		config:          aq.config,
		limit:           aq.limit,
		offset:          aq.offset,
		order:           append([]OrderFunc{}, aq.order...),
		predicates:      append([]predicate.Activity{}, aq.predicates...),
		withScores:      aq.withScores.Clone(),
		withArea:        aq.withArea.Clone(),
		withClassPeriod: aq.withClassPeriod.Clone(),
		// clone intermediate query.
		sql:    aq.sql.Clone(),
		path:   aq.path,
		unique: aq.unique,
	}
}

// WithScores tells the query-builder to eager-load the nodes that are connected to
// the "scores" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ActivityQuery) WithScores(opts ...func(*ScoreQuery)) *ActivityQuery {
	query := &ScoreQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withScores = query
	return aq
}

// WithArea tells the query-builder to eager-load the nodes that are connected to
// the "area" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ActivityQuery) WithArea(opts ...func(*AreaQuery)) *ActivityQuery {
	query := &AreaQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withArea = query
	return aq
}

// WithClassPeriod tells the query-builder to eager-load the nodes that are connected to
// the "classPeriod" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ActivityQuery) WithClassPeriod(opts ...func(*ClassPeriodQuery)) *ActivityQuery {
	query := &ClassPeriodQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withClassPeriod = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Activity.Query().
//		GroupBy(activity.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *ActivityQuery) GroupBy(field string, fields ...string) *ActivityGroupBy {
	grbuild := &ActivityGroupBy{config: aq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	grbuild.label = activity.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Activity.Query().
//		Select(activity.FieldName).
//		Scan(ctx, &v)
func (aq *ActivityQuery) Select(fields ...string) *ActivitySelect {
	aq.fields = append(aq.fields, fields...)
	selbuild := &ActivitySelect{ActivityQuery: aq}
	selbuild.label = activity.Label
	selbuild.flds, selbuild.scan = &aq.fields, selbuild.Scan
	return selbuild
}

func (aq *ActivityQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !activity.ValidColumn(f) {
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

func (aq *ActivityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Activity, error) {
	var (
		nodes       = []*Activity{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [3]bool{
			aq.withScores != nil,
			aq.withArea != nil,
			aq.withClassPeriod != nil,
		}
	)
	if aq.withArea != nil || aq.withClassPeriod != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, activity.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Activity).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Activity{config: aq.config}
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
	if query := aq.withScores; query != nil {
		if err := aq.loadScores(ctx, query, nodes,
			func(n *Activity) { n.Edges.Scores = []*Score{} },
			func(n *Activity, e *Score) { n.Edges.Scores = append(n.Edges.Scores, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withArea; query != nil {
		if err := aq.loadArea(ctx, query, nodes, nil,
			func(n *Activity, e *Area) { n.Edges.Area = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withClassPeriod; query != nil {
		if err := aq.loadClassPeriod(ctx, query, nodes, nil,
			func(n *Activity, e *ClassPeriod) { n.Edges.ClassPeriod = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *ActivityQuery) loadScores(ctx context.Context, query *ScoreQuery, nodes []*Activity, init func(*Activity), assign func(*Activity, *Score)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Activity)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Score(func(s *sql.Selector) {
		s.Where(sql.InValues(activity.ScoresColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.activity_scores
		if fk == nil {
			return fmt.Errorf(`foreign-key "activity_scores" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "activity_scores" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *ActivityQuery) loadArea(ctx context.Context, query *AreaQuery, nodes []*Activity, init func(*Activity), assign func(*Activity, *Area)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Activity)
	for i := range nodes {
		if nodes[i].area_activities == nil {
			continue
		}
		fk := *nodes[i].area_activities
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(area.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "area_activities" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *ActivityQuery) loadClassPeriod(ctx context.Context, query *ClassPeriodQuery, nodes []*Activity, init func(*Activity), assign func(*Activity, *ClassPeriod)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Activity)
	for i := range nodes {
		if nodes[i].class_period_activities == nil {
			continue
		}
		fk := *nodes[i].class_period_activities
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(classperiod.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "class_period_activities" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *ActivityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.fields
	if len(aq.fields) > 0 {
		_spec.Unique = aq.unique != nil && *aq.unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *ActivityQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aq *ActivityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   activity.Table,
			Columns: activity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: activity.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, activity.FieldID)
		for i := range fields {
			if fields[i] != activity.FieldID {
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

func (aq *ActivityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(activity.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = activity.Columns
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

// ActivityGroupBy is the group-by builder for Activity entities.
type ActivityGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *ActivityGroupBy) Aggregate(fns ...AggregateFunc) *ActivityGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *ActivityGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

func (agb *ActivityGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range agb.fields {
		if !activity.ValidColumn(f) {
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

func (agb *ActivityGroupBy) sqlQuery() *sql.Selector {
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

// ActivitySelect is the builder for selecting fields of Activity entities.
type ActivitySelect struct {
	*ActivityQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (as *ActivitySelect) Scan(ctx context.Context, v interface{}) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.ActivityQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

func (as *ActivitySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
