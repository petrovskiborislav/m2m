// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/bug/ent/organizationalunit"
	"entgo.io/bug/ent/organizationalunittype"
	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/specialty"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrganizationalUnitQuery is the builder for querying OrganizationalUnit entities.
type OrganizationalUnitQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrganizationalUnit
	// eager-loading edges.
	withParent                   *OrganizationalUnitQuery
	withChildren                 *OrganizationalUnitQuery
	withOrganizationalUnitTypeID *OrganizationalUnitTypeQuery
	withSpecialties              *SpecialtyQuery
	modifiers                    []func(*sql.Selector)
	loadTotal                    []func(context.Context, []*OrganizationalUnit) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrganizationalUnitQuery builder.
func (ouq *OrganizationalUnitQuery) Where(ps ...predicate.OrganizationalUnit) *OrganizationalUnitQuery {
	ouq.predicates = append(ouq.predicates, ps...)
	return ouq
}

// Limit adds a limit step to the query.
func (ouq *OrganizationalUnitQuery) Limit(limit int) *OrganizationalUnitQuery {
	ouq.limit = &limit
	return ouq
}

// Offset adds an offset step to the query.
func (ouq *OrganizationalUnitQuery) Offset(offset int) *OrganizationalUnitQuery {
	ouq.offset = &offset
	return ouq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ouq *OrganizationalUnitQuery) Unique(unique bool) *OrganizationalUnitQuery {
	ouq.unique = &unique
	return ouq
}

// Order adds an order step to the query.
func (ouq *OrganizationalUnitQuery) Order(o ...OrderFunc) *OrganizationalUnitQuery {
	ouq.order = append(ouq.order, o...)
	return ouq
}

// QueryParent chains the current query on the "parent" edge.
func (ouq *OrganizationalUnitQuery) QueryParent() *OrganizationalUnitQuery {
	query := &OrganizationalUnitQuery{config: ouq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ouq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ouq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationalunit.Table, organizationalunit.FieldID, selector),
			sqlgraph.To(organizationalunit.Table, organizationalunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, organizationalunit.ParentTable, organizationalunit.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ouq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (ouq *OrganizationalUnitQuery) QueryChildren() *OrganizationalUnitQuery {
	query := &OrganizationalUnitQuery{config: ouq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ouq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ouq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationalunit.Table, organizationalunit.FieldID, selector),
			sqlgraph.To(organizationalunit.Table, organizationalunit.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, organizationalunit.ChildrenTable, organizationalunit.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(ouq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganizationalUnitTypeID chains the current query on the "organizational_unit_type_id" edge.
func (ouq *OrganizationalUnitQuery) QueryOrganizationalUnitTypeID() *OrganizationalUnitTypeQuery {
	query := &OrganizationalUnitTypeQuery{config: ouq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ouq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ouq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationalunit.Table, organizationalunit.FieldID, selector),
			sqlgraph.To(organizationalunittype.Table, organizationalunittype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, organizationalunit.OrganizationalUnitTypeIDTable, organizationalunit.OrganizationalUnitTypeIDColumn),
		)
		fromU = sqlgraph.SetNeighbors(ouq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySpecialties chains the current query on the "specialties" edge.
func (ouq *OrganizationalUnitQuery) QuerySpecialties() *SpecialtyQuery {
	query := &SpecialtyQuery{config: ouq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ouq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ouq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationalunit.Table, organizationalunit.FieldID, selector),
			sqlgraph.To(specialty.Table, specialty.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, organizationalunit.SpecialtiesTable, organizationalunit.SpecialtiesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ouq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrganizationalUnit entity from the query.
// Returns a *NotFoundError when no OrganizationalUnit was found.
func (ouq *OrganizationalUnitQuery) First(ctx context.Context) (*OrganizationalUnit, error) {
	nodes, err := ouq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{organizationalunit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ouq *OrganizationalUnitQuery) FirstX(ctx context.Context) *OrganizationalUnit {
	node, err := ouq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrganizationalUnit ID from the query.
// Returns a *NotFoundError when no OrganizationalUnit ID was found.
func (ouq *OrganizationalUnitQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ouq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{organizationalunit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ouq *OrganizationalUnitQuery) FirstIDX(ctx context.Context) string {
	id, err := ouq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrganizationalUnit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrganizationalUnit entity is found.
// Returns a *NotFoundError when no OrganizationalUnit entities are found.
func (ouq *OrganizationalUnitQuery) Only(ctx context.Context) (*OrganizationalUnit, error) {
	nodes, err := ouq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{organizationalunit.Label}
	default:
		return nil, &NotSingularError{organizationalunit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ouq *OrganizationalUnitQuery) OnlyX(ctx context.Context) *OrganizationalUnit {
	node, err := ouq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrganizationalUnit ID in the query.
// Returns a *NotSingularError when more than one OrganizationalUnit ID is found.
// Returns a *NotFoundError when no entities are found.
func (ouq *OrganizationalUnitQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ouq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{organizationalunit.Label}
	default:
		err = &NotSingularError{organizationalunit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ouq *OrganizationalUnitQuery) OnlyIDX(ctx context.Context) string {
	id, err := ouq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrganizationalUnits.
func (ouq *OrganizationalUnitQuery) All(ctx context.Context) ([]*OrganizationalUnit, error) {
	if err := ouq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ouq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ouq *OrganizationalUnitQuery) AllX(ctx context.Context) []*OrganizationalUnit {
	nodes, err := ouq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrganizationalUnit IDs.
func (ouq *OrganizationalUnitQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := ouq.Select(organizationalunit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ouq *OrganizationalUnitQuery) IDsX(ctx context.Context) []string {
	ids, err := ouq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ouq *OrganizationalUnitQuery) Count(ctx context.Context) (int, error) {
	if err := ouq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ouq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ouq *OrganizationalUnitQuery) CountX(ctx context.Context) int {
	count, err := ouq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ouq *OrganizationalUnitQuery) Exist(ctx context.Context) (bool, error) {
	if err := ouq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ouq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ouq *OrganizationalUnitQuery) ExistX(ctx context.Context) bool {
	exist, err := ouq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrganizationalUnitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ouq *OrganizationalUnitQuery) Clone() *OrganizationalUnitQuery {
	if ouq == nil {
		return nil
	}
	return &OrganizationalUnitQuery{
		config:                       ouq.config,
		limit:                        ouq.limit,
		offset:                       ouq.offset,
		order:                        append([]OrderFunc{}, ouq.order...),
		predicates:                   append([]predicate.OrganizationalUnit{}, ouq.predicates...),
		withParent:                   ouq.withParent.Clone(),
		withChildren:                 ouq.withChildren.Clone(),
		withOrganizationalUnitTypeID: ouq.withOrganizationalUnitTypeID.Clone(),
		withSpecialties:              ouq.withSpecialties.Clone(),
		// clone intermediate query.
		sql:    ouq.sql.Clone(),
		path:   ouq.path,
		unique: ouq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (ouq *OrganizationalUnitQuery) WithParent(opts ...func(*OrganizationalUnitQuery)) *OrganizationalUnitQuery {
	query := &OrganizationalUnitQuery{config: ouq.config}
	for _, opt := range opts {
		opt(query)
	}
	ouq.withParent = query
	return ouq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (ouq *OrganizationalUnitQuery) WithChildren(opts ...func(*OrganizationalUnitQuery)) *OrganizationalUnitQuery {
	query := &OrganizationalUnitQuery{config: ouq.config}
	for _, opt := range opts {
		opt(query)
	}
	ouq.withChildren = query
	return ouq
}

// WithOrganizationalUnitTypeID tells the query-builder to eager-load the nodes that are connected to
// the "organizational_unit_type_id" edge. The optional arguments are used to configure the query builder of the edge.
func (ouq *OrganizationalUnitQuery) WithOrganizationalUnitTypeID(opts ...func(*OrganizationalUnitTypeQuery)) *OrganizationalUnitQuery {
	query := &OrganizationalUnitTypeQuery{config: ouq.config}
	for _, opt := range opts {
		opt(query)
	}
	ouq.withOrganizationalUnitTypeID = query
	return ouq
}

// WithSpecialties tells the query-builder to eager-load the nodes that are connected to
// the "specialties" edge. The optional arguments are used to configure the query builder of the edge.
func (ouq *OrganizationalUnitQuery) WithSpecialties(opts ...func(*SpecialtyQuery)) *OrganizationalUnitQuery {
	query := &SpecialtyQuery{config: ouq.config}
	for _, opt := range opts {
		opt(query)
	}
	ouq.withSpecialties = query
	return ouq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DisplayName string `json:"display_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrganizationalUnit.Query().
//		GroupBy(organizationalunit.FieldDisplayName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ouq *OrganizationalUnitQuery) GroupBy(field string, fields ...string) *OrganizationalUnitGroupBy {
	grbuild := &OrganizationalUnitGroupBy{config: ouq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ouq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ouq.sqlQuery(ctx), nil
	}
	grbuild.label = organizationalunit.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DisplayName string `json:"display_name,omitempty"`
//	}
//
//	client.OrganizationalUnit.Query().
//		Select(organizationalunit.FieldDisplayName).
//		Scan(ctx, &v)
//
func (ouq *OrganizationalUnitQuery) Select(fields ...string) *OrganizationalUnitSelect {
	ouq.fields = append(ouq.fields, fields...)
	selbuild := &OrganizationalUnitSelect{OrganizationalUnitQuery: ouq}
	selbuild.label = organizationalunit.Label
	selbuild.flds, selbuild.scan = &ouq.fields, selbuild.Scan
	return selbuild
}

func (ouq *OrganizationalUnitQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ouq.fields {
		if !organizationalunit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ouq.path != nil {
		prev, err := ouq.path(ctx)
		if err != nil {
			return err
		}
		ouq.sql = prev
	}
	return nil
}

func (ouq *OrganizationalUnitQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrganizationalUnit, error) {
	var (
		nodes       = []*OrganizationalUnit{}
		_spec       = ouq.querySpec()
		loadedTypes = [4]bool{
			ouq.withParent != nil,
			ouq.withChildren != nil,
			ouq.withOrganizationalUnitTypeID != nil,
			ouq.withSpecialties != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OrganizationalUnit).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OrganizationalUnit{config: ouq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ouq.modifiers) > 0 {
		_spec.Modifiers = ouq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ouq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ouq.withParent; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*OrganizationalUnit)
		for i := range nodes {
			fk := nodes[i].ParentID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(organizationalunit.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "parent_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := ouq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*OrganizationalUnit)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*OrganizationalUnit{}
		}
		query.Where(predicate.OrganizationalUnit(func(s *sql.Selector) {
			s.Where(sql.InValues(organizationalunit.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ParentID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "parent_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	if query := ouq.withOrganizationalUnitTypeID; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*OrganizationalUnit)
		for i := range nodes {
			fk := nodes[i].TypeID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(organizationalunittype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "type_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OrganizationalUnitTypeID = n
			}
		}
	}

	if query := ouq.withSpecialties; query != nil {
		edgeids := make([]driver.Value, len(nodes))
		byid := make(map[string]*OrganizationalUnit)
		nids := make(map[string]map[*OrganizationalUnit]struct{})
		for i, node := range nodes {
			edgeids[i] = node.ID
			byid[node.ID] = node
			node.Edges.Specialties = []*Specialty{}
		}
		query.Where(func(s *sql.Selector) {
			joinT := sql.Table(organizationalunit.SpecialtiesTable)
			s.Join(joinT).On(s.C(specialty.FieldID), joinT.C(organizationalunit.SpecialtiesPrimaryKey[1]))
			s.Where(sql.InValues(joinT.C(organizationalunit.SpecialtiesPrimaryKey[0]), edgeids...))
			columns := s.SelectedColumns()
			s.Select(joinT.C(organizationalunit.SpecialtiesPrimaryKey[0]))
			s.AppendSelect(columns...)
			s.SetDistinct(false)
		})
		neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]interface{}, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]interface{}{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []interface{}) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*OrganizationalUnit]struct{}{byid[outValue]: struct{}{}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byid[outValue]] = struct{}{}
				return nil
			}
		})
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "specialties" node returned %v`, n.ID)
			}
			for kn := range nodes {
				kn.Edges.Specialties = append(kn.Edges.Specialties, n)
			}
		}
	}

	for i := range ouq.loadTotal {
		if err := ouq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ouq *OrganizationalUnitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ouq.querySpec()
	if len(ouq.modifiers) > 0 {
		_spec.Modifiers = ouq.modifiers
	}
	_spec.Node.Columns = ouq.fields
	if len(ouq.fields) > 0 {
		_spec.Unique = ouq.unique != nil && *ouq.unique
	}
	return sqlgraph.CountNodes(ctx, ouq.driver, _spec)
}

func (ouq *OrganizationalUnitQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ouq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ouq *OrganizationalUnitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   organizationalunit.Table,
			Columns: organizationalunit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: organizationalunit.FieldID,
			},
		},
		From:   ouq.sql,
		Unique: true,
	}
	if unique := ouq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ouq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationalunit.FieldID)
		for i := range fields {
			if fields[i] != organizationalunit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ouq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ouq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ouq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ouq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ouq *OrganizationalUnitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ouq.driver.Dialect())
	t1 := builder.Table(organizationalunit.Table)
	columns := ouq.fields
	if len(columns) == 0 {
		columns = organizationalunit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ouq.sql != nil {
		selector = ouq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ouq.unique != nil && *ouq.unique {
		selector.Distinct()
	}
	for _, p := range ouq.predicates {
		p(selector)
	}
	for _, p := range ouq.order {
		p(selector)
	}
	if offset := ouq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ouq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrganizationalUnitGroupBy is the group-by builder for OrganizationalUnit entities.
type OrganizationalUnitGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ougb *OrganizationalUnitGroupBy) Aggregate(fns ...AggregateFunc) *OrganizationalUnitGroupBy {
	ougb.fns = append(ougb.fns, fns...)
	return ougb
}

// Scan applies the group-by query and scans the result into the given value.
func (ougb *OrganizationalUnitGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ougb.path(ctx)
	if err != nil {
		return err
	}
	ougb.sql = query
	return ougb.sqlScan(ctx, v)
}

func (ougb *OrganizationalUnitGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ougb.fields {
		if !organizationalunit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ougb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ougb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ougb *OrganizationalUnitGroupBy) sqlQuery() *sql.Selector {
	selector := ougb.sql.Select()
	aggregation := make([]string, 0, len(ougb.fns))
	for _, fn := range ougb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ougb.fields)+len(ougb.fns))
		for _, f := range ougb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ougb.fields...)...)
}

// OrganizationalUnitSelect is the builder for selecting fields of OrganizationalUnit entities.
type OrganizationalUnitSelect struct {
	*OrganizationalUnitQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ous *OrganizationalUnitSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ous.prepareQuery(ctx); err != nil {
		return err
	}
	ous.sql = ous.OrganizationalUnitQuery.sqlQuery(ctx)
	return ous.sqlScan(ctx, v)
}

func (ous *OrganizationalUnitSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ous.sql.Query()
	if err := ous.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
