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
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/period"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/year"
)

// PeriodQuery is the builder for querying Period entities.
type PeriodQuery struct {
	config
	limit            *int
	offset           *int
	unique           *bool
	order            []OrderFunc
	fields           []string
	predicates       []predicate.Period
	withClassPeriods *ClassPeriodQuery
	withYear         *YearQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PeriodQuery builder.
func (pq *PeriodQuery) Where(ps ...predicate.Period) *PeriodQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PeriodQuery) Limit(limit int) *PeriodQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PeriodQuery) Offset(offset int) *PeriodQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PeriodQuery) Unique(unique bool) *PeriodQuery {
	pq.unique = &unique
	return pq
}

// Order adds an order step to the query.
func (pq *PeriodQuery) Order(o ...OrderFunc) *PeriodQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryClassPeriods chains the current query on the "classPeriods" edge.
func (pq *PeriodQuery) QueryClassPeriods() *ClassPeriodQuery {
	query := &ClassPeriodQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(period.Table, period.FieldID, selector),
			sqlgraph.To(classperiod.Table, classperiod.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, period.ClassPeriodsTable, period.ClassPeriodsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryYear chains the current query on the "year" edge.
func (pq *PeriodQuery) QueryYear() *YearQuery {
	query := &YearQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(period.Table, period.FieldID, selector),
			sqlgraph.To(year.Table, year.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, period.YearTable, period.YearColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Period entity from the query.
// Returns a *NotFoundError when no Period was found.
func (pq *PeriodQuery) First(ctx context.Context) (*Period, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{period.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PeriodQuery) FirstX(ctx context.Context) *Period {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Period ID from the query.
// Returns a *NotFoundError when no Period ID was found.
func (pq *PeriodQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{period.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PeriodQuery) FirstIDX(ctx context.Context) string {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Period entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Period entity is found.
// Returns a *NotFoundError when no Period entities are found.
func (pq *PeriodQuery) Only(ctx context.Context) (*Period, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{period.Label}
	default:
		return nil, &NotSingularError{period.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PeriodQuery) OnlyX(ctx context.Context) *Period {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Period ID in the query.
// Returns a *NotSingularError when more than one Period ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PeriodQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{period.Label}
	default:
		err = &NotSingularError{period.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PeriodQuery) OnlyIDX(ctx context.Context) string {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Periods.
func (pq *PeriodQuery) All(ctx context.Context) ([]*Period, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PeriodQuery) AllX(ctx context.Context) []*Period {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Period IDs.
func (pq *PeriodQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := pq.Select(period.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PeriodQuery) IDsX(ctx context.Context) []string {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PeriodQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PeriodQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PeriodQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PeriodQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PeriodQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PeriodQuery) Clone() *PeriodQuery {
	if pq == nil {
		return nil
	}
	return &PeriodQuery{
		config:           pq.config,
		limit:            pq.limit,
		offset:           pq.offset,
		order:            append([]OrderFunc{}, pq.order...),
		predicates:       append([]predicate.Period{}, pq.predicates...),
		withClassPeriods: pq.withClassPeriods.Clone(),
		withYear:         pq.withYear.Clone(),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
	}
}

// WithClassPeriods tells the query-builder to eager-load the nodes that are connected to
// the "classPeriods" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PeriodQuery) WithClassPeriods(opts ...func(*ClassPeriodQuery)) *PeriodQuery {
	query := &ClassPeriodQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withClassPeriods = query
	return pq
}

// WithYear tells the query-builder to eager-load the nodes that are connected to
// the "year" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PeriodQuery) WithYear(opts ...func(*YearQuery)) *PeriodQuery {
	query := &YearQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withYear = query
	return pq
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
//	client.Period.Query().
//		GroupBy(period.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PeriodQuery) GroupBy(field string, fields ...string) *PeriodGroupBy {
	grbuild := &PeriodGroupBy{config: pq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(ctx), nil
	}
	grbuild.label = period.Label
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
//	client.Period.Query().
//		Select(period.FieldName).
//		Scan(ctx, &v)
func (pq *PeriodQuery) Select(fields ...string) *PeriodSelect {
	pq.fields = append(pq.fields, fields...)
	selbuild := &PeriodSelect{PeriodQuery: pq}
	selbuild.label = period.Label
	selbuild.flds, selbuild.scan = &pq.fields, selbuild.Scan
	return selbuild
}

func (pq *PeriodQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !period.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PeriodQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Period, error) {
	var (
		nodes       = []*Period{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withClassPeriods != nil,
			pq.withYear != nil,
		}
	)
	if pq.withYear != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, period.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Period).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Period{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withClassPeriods; query != nil {
		if err := pq.loadClassPeriods(ctx, query, nodes,
			func(n *Period) { n.Edges.ClassPeriods = []*ClassPeriod{} },
			func(n *Period, e *ClassPeriod) { n.Edges.ClassPeriods = append(n.Edges.ClassPeriods, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withYear; query != nil {
		if err := pq.loadYear(ctx, query, nodes, nil,
			func(n *Period, e *Year) { n.Edges.Year = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PeriodQuery) loadClassPeriods(ctx context.Context, query *ClassPeriodQuery, nodes []*Period, init func(*Period), assign func(*Period, *ClassPeriod)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Period)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ClassPeriod(func(s *sql.Selector) {
		s.Where(sql.InValues(period.ClassPeriodsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.period_class_periods
		if fk == nil {
			return fmt.Errorf(`foreign-key "period_class_periods" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "period_class_periods" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *PeriodQuery) loadYear(ctx context.Context, query *YearQuery, nodes []*Period, init func(*Period), assign func(*Period, *Year)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Period)
	for i := range nodes {
		if nodes[i].year_periods == nil {
			continue
		}
		fk := *nodes[i].year_periods
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(year.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "year_periods" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PeriodQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PeriodQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pq *PeriodQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   period.Table,
			Columns: period.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: period.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, period.FieldID)
		for i := range fields {
			if fields[i] != period.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PeriodQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(period.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = period.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PeriodGroupBy is the group-by builder for Period entities.
type PeriodGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PeriodGroupBy) Aggregate(fns ...AggregateFunc) *PeriodGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *PeriodGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

func (pgb *PeriodGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pgb.fields {
		if !period.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PeriodGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql.Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
		for _, f := range pgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pgb.fields...)...)
}

// PeriodSelect is the builder for selecting fields of Period entities.
type PeriodSelect struct {
	*PeriodQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PeriodSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.sql = ps.PeriodQuery.sqlQuery(ctx)
	return ps.sqlScan(ctx, v)
}

func (ps *PeriodSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sql.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
