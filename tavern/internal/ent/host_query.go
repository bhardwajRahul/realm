// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"realm.pub/tavern/internal/ent/beacon"
	"realm.pub/tavern/internal/ent/host"
	"realm.pub/tavern/internal/ent/hostcredential"
	"realm.pub/tavern/internal/ent/hostfile"
	"realm.pub/tavern/internal/ent/hostprocess"
	"realm.pub/tavern/internal/ent/predicate"
	"realm.pub/tavern/internal/ent/tag"
)

// HostQuery is the builder for querying Host entities.
type HostQuery struct {
	config
	ctx                  *QueryContext
	order                []host.OrderOption
	inters               []Interceptor
	predicates           []predicate.Host
	withTags             *TagQuery
	withBeacons          *BeaconQuery
	withFiles            *HostFileQuery
	withProcesses        *HostProcessQuery
	withCredentials      *HostCredentialQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*Host) error
	withNamedTags        map[string]*TagQuery
	withNamedBeacons     map[string]*BeaconQuery
	withNamedFiles       map[string]*HostFileQuery
	withNamedProcesses   map[string]*HostProcessQuery
	withNamedCredentials map[string]*HostCredentialQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HostQuery builder.
func (hq *HostQuery) Where(ps ...predicate.Host) *HostQuery {
	hq.predicates = append(hq.predicates, ps...)
	return hq
}

// Limit the number of records to be returned by this query.
func (hq *HostQuery) Limit(limit int) *HostQuery {
	hq.ctx.Limit = &limit
	return hq
}

// Offset to start from.
func (hq *HostQuery) Offset(offset int) *HostQuery {
	hq.ctx.Offset = &offset
	return hq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hq *HostQuery) Unique(unique bool) *HostQuery {
	hq.ctx.Unique = &unique
	return hq
}

// Order specifies how the records should be ordered.
func (hq *HostQuery) Order(o ...host.OrderOption) *HostQuery {
	hq.order = append(hq.order, o...)
	return hq
}

// QueryTags chains the current query on the "tags" edge.
func (hq *HostQuery) QueryTags() *TagQuery {
	query := (&TagClient{config: hq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, host.TagsTable, host.TagsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBeacons chains the current query on the "beacons" edge.
func (hq *HostQuery) QueryBeacons() *BeaconQuery {
	query := (&BeaconClient{config: hq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, selector),
			sqlgraph.To(beacon.Table, beacon.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, host.BeaconsTable, host.BeaconsColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFiles chains the current query on the "files" edge.
func (hq *HostQuery) QueryFiles() *HostFileQuery {
	query := (&HostFileClient{config: hq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, selector),
			sqlgraph.To(hostfile.Table, hostfile.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, host.FilesTable, host.FilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProcesses chains the current query on the "processes" edge.
func (hq *HostQuery) QueryProcesses() *HostProcessQuery {
	query := (&HostProcessClient{config: hq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, selector),
			sqlgraph.To(hostprocess.Table, hostprocess.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, host.ProcessesTable, host.ProcessesColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCredentials chains the current query on the "credentials" edge.
func (hq *HostQuery) QueryCredentials() *HostCredentialQuery {
	query := (&HostCredentialClient{config: hq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, selector),
			sqlgraph.To(hostcredential.Table, hostcredential.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, host.CredentialsTable, host.CredentialsColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Host entity from the query.
// Returns a *NotFoundError when no Host was found.
func (hq *HostQuery) First(ctx context.Context) (*Host, error) {
	nodes, err := hq.Limit(1).All(setContextOp(ctx, hq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{host.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hq *HostQuery) FirstX(ctx context.Context) *Host {
	node, err := hq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Host ID from the query.
// Returns a *NotFoundError when no Host ID was found.
func (hq *HostQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(1).IDs(setContextOp(ctx, hq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{host.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hq *HostQuery) FirstIDX(ctx context.Context) int {
	id, err := hq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Host entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Host entity is found.
// Returns a *NotFoundError when no Host entities are found.
func (hq *HostQuery) Only(ctx context.Context) (*Host, error) {
	nodes, err := hq.Limit(2).All(setContextOp(ctx, hq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{host.Label}
	default:
		return nil, &NotSingularError{host.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hq *HostQuery) OnlyX(ctx context.Context) *Host {
	node, err := hq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Host ID in the query.
// Returns a *NotSingularError when more than one Host ID is found.
// Returns a *NotFoundError when no entities are found.
func (hq *HostQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(2).IDs(setContextOp(ctx, hq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{host.Label}
	default:
		err = &NotSingularError{host.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hq *HostQuery) OnlyIDX(ctx context.Context) int {
	id, err := hq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Hosts.
func (hq *HostQuery) All(ctx context.Context) ([]*Host, error) {
	ctx = setContextOp(ctx, hq.ctx, ent.OpQueryAll)
	if err := hq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Host, *HostQuery]()
	return withInterceptors[[]*Host](ctx, hq, qr, hq.inters)
}

// AllX is like All, but panics if an error occurs.
func (hq *HostQuery) AllX(ctx context.Context) []*Host {
	nodes, err := hq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Host IDs.
func (hq *HostQuery) IDs(ctx context.Context) (ids []int, err error) {
	if hq.ctx.Unique == nil && hq.path != nil {
		hq.Unique(true)
	}
	ctx = setContextOp(ctx, hq.ctx, ent.OpQueryIDs)
	if err = hq.Select(host.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hq *HostQuery) IDsX(ctx context.Context) []int {
	ids, err := hq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hq *HostQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, hq.ctx, ent.OpQueryCount)
	if err := hq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, hq, querierCount[*HostQuery](), hq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (hq *HostQuery) CountX(ctx context.Context) int {
	count, err := hq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hq *HostQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, hq.ctx, ent.OpQueryExist)
	switch _, err := hq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (hq *HostQuery) ExistX(ctx context.Context) bool {
	exist, err := hq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HostQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hq *HostQuery) Clone() *HostQuery {
	if hq == nil {
		return nil
	}
	return &HostQuery{
		config:          hq.config,
		ctx:             hq.ctx.Clone(),
		order:           append([]host.OrderOption{}, hq.order...),
		inters:          append([]Interceptor{}, hq.inters...),
		predicates:      append([]predicate.Host{}, hq.predicates...),
		withTags:        hq.withTags.Clone(),
		withBeacons:     hq.withBeacons.Clone(),
		withFiles:       hq.withFiles.Clone(),
		withProcesses:   hq.withProcesses.Clone(),
		withCredentials: hq.withCredentials.Clone(),
		// clone intermediate query.
		sql:  hq.sql.Clone(),
		path: hq.path,
	}
}

// WithTags tells the query-builder to eager-load the nodes that are connected to
// the "tags" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HostQuery) WithTags(opts ...func(*TagQuery)) *HostQuery {
	query := (&TagClient{config: hq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hq.withTags = query
	return hq
}

// WithBeacons tells the query-builder to eager-load the nodes that are connected to
// the "beacons" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HostQuery) WithBeacons(opts ...func(*BeaconQuery)) *HostQuery {
	query := (&BeaconClient{config: hq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hq.withBeacons = query
	return hq
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HostQuery) WithFiles(opts ...func(*HostFileQuery)) *HostQuery {
	query := (&HostFileClient{config: hq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hq.withFiles = query
	return hq
}

// WithProcesses tells the query-builder to eager-load the nodes that are connected to
// the "processes" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HostQuery) WithProcesses(opts ...func(*HostProcessQuery)) *HostQuery {
	query := (&HostProcessClient{config: hq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hq.withProcesses = query
	return hq
}

// WithCredentials tells the query-builder to eager-load the nodes that are connected to
// the "credentials" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HostQuery) WithCredentials(opts ...func(*HostCredentialQuery)) *HostQuery {
	query := (&HostCredentialClient{config: hq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hq.withCredentials = query
	return hq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Host.Query().
//		GroupBy(host.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (hq *HostQuery) GroupBy(field string, fields ...string) *HostGroupBy {
	hq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &HostGroupBy{build: hq}
	grbuild.flds = &hq.ctx.Fields
	grbuild.label = host.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Host.Query().
//		Select(host.FieldCreatedAt).
//		Scan(ctx, &v)
func (hq *HostQuery) Select(fields ...string) *HostSelect {
	hq.ctx.Fields = append(hq.ctx.Fields, fields...)
	sbuild := &HostSelect{HostQuery: hq}
	sbuild.label = host.Label
	sbuild.flds, sbuild.scan = &hq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a HostSelect configured with the given aggregations.
func (hq *HostQuery) Aggregate(fns ...AggregateFunc) *HostSelect {
	return hq.Select().Aggregate(fns...)
}

func (hq *HostQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range hq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, hq); err != nil {
				return err
			}
		}
	}
	for _, f := range hq.ctx.Fields {
		if !host.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hq.path != nil {
		prev, err := hq.path(ctx)
		if err != nil {
			return err
		}
		hq.sql = prev
	}
	return nil
}

func (hq *HostQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Host, error) {
	var (
		nodes       = []*Host{}
		_spec       = hq.querySpec()
		loadedTypes = [5]bool{
			hq.withTags != nil,
			hq.withBeacons != nil,
			hq.withFiles != nil,
			hq.withProcesses != nil,
			hq.withCredentials != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Host).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Host{config: hq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(hq.modifiers) > 0 {
		_spec.Modifiers = hq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := hq.withTags; query != nil {
		if err := hq.loadTags(ctx, query, nodes,
			func(n *Host) { n.Edges.Tags = []*Tag{} },
			func(n *Host, e *Tag) { n.Edges.Tags = append(n.Edges.Tags, e) }); err != nil {
			return nil, err
		}
	}
	if query := hq.withBeacons; query != nil {
		if err := hq.loadBeacons(ctx, query, nodes,
			func(n *Host) { n.Edges.Beacons = []*Beacon{} },
			func(n *Host, e *Beacon) { n.Edges.Beacons = append(n.Edges.Beacons, e) }); err != nil {
			return nil, err
		}
	}
	if query := hq.withFiles; query != nil {
		if err := hq.loadFiles(ctx, query, nodes,
			func(n *Host) { n.Edges.Files = []*HostFile{} },
			func(n *Host, e *HostFile) { n.Edges.Files = append(n.Edges.Files, e) }); err != nil {
			return nil, err
		}
	}
	if query := hq.withProcesses; query != nil {
		if err := hq.loadProcesses(ctx, query, nodes,
			func(n *Host) { n.Edges.Processes = []*HostProcess{} },
			func(n *Host, e *HostProcess) { n.Edges.Processes = append(n.Edges.Processes, e) }); err != nil {
			return nil, err
		}
	}
	if query := hq.withCredentials; query != nil {
		if err := hq.loadCredentials(ctx, query, nodes,
			func(n *Host) { n.Edges.Credentials = []*HostCredential{} },
			func(n *Host, e *HostCredential) { n.Edges.Credentials = append(n.Edges.Credentials, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range hq.withNamedTags {
		if err := hq.loadTags(ctx, query, nodes,
			func(n *Host) { n.appendNamedTags(name) },
			func(n *Host, e *Tag) { n.appendNamedTags(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range hq.withNamedBeacons {
		if err := hq.loadBeacons(ctx, query, nodes,
			func(n *Host) { n.appendNamedBeacons(name) },
			func(n *Host, e *Beacon) { n.appendNamedBeacons(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range hq.withNamedFiles {
		if err := hq.loadFiles(ctx, query, nodes,
			func(n *Host) { n.appendNamedFiles(name) },
			func(n *Host, e *HostFile) { n.appendNamedFiles(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range hq.withNamedProcesses {
		if err := hq.loadProcesses(ctx, query, nodes,
			func(n *Host) { n.appendNamedProcesses(name) },
			func(n *Host, e *HostProcess) { n.appendNamedProcesses(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range hq.withNamedCredentials {
		if err := hq.loadCredentials(ctx, query, nodes,
			func(n *Host) { n.appendNamedCredentials(name) },
			func(n *Host, e *HostCredential) { n.appendNamedCredentials(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range hq.loadTotal {
		if err := hq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (hq *HostQuery) loadTags(ctx context.Context, query *TagQuery, nodes []*Host, init func(*Host), assign func(*Host, *Tag)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Host)
	nids := make(map[int]map[*Host]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(host.TagsTable)
		s.Join(joinT).On(s.C(tag.FieldID), joinT.C(host.TagsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(host.TagsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(host.TagsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Host]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Tag](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "tags" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (hq *HostQuery) loadBeacons(ctx context.Context, query *BeaconQuery, nodes []*Host, init func(*Host), assign func(*Host, *Beacon)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Host)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Beacon(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(host.BeaconsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.beacon_host
		if fk == nil {
			return fmt.Errorf(`foreign-key "beacon_host" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "beacon_host" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (hq *HostQuery) loadFiles(ctx context.Context, query *HostFileQuery, nodes []*Host, init func(*Host), assign func(*Host, *HostFile)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Host)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.HostFile(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(host.FilesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.host_files
		if fk == nil {
			return fmt.Errorf(`foreign-key "host_files" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "host_files" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (hq *HostQuery) loadProcesses(ctx context.Context, query *HostProcessQuery, nodes []*Host, init func(*Host), assign func(*Host, *HostProcess)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Host)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.HostProcess(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(host.ProcessesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.host_processes
		if fk == nil {
			return fmt.Errorf(`foreign-key "host_processes" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "host_processes" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (hq *HostQuery) loadCredentials(ctx context.Context, query *HostCredentialQuery, nodes []*Host, init func(*Host), assign func(*Host, *HostCredential)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Host)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.HostCredential(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(host.CredentialsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.host_credential_host
		if fk == nil {
			return fmt.Errorf(`foreign-key "host_credential_host" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "host_credential_host" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (hq *HostQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hq.querySpec()
	if len(hq.modifiers) > 0 {
		_spec.Modifiers = hq.modifiers
	}
	_spec.Node.Columns = hq.ctx.Fields
	if len(hq.ctx.Fields) > 0 {
		_spec.Unique = hq.ctx.Unique != nil && *hq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, hq.driver, _spec)
}

func (hq *HostQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(host.Table, host.Columns, sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt))
	_spec.From = hq.sql
	if unique := hq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if hq.path != nil {
		_spec.Unique = true
	}
	if fields := hq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, host.FieldID)
		for i := range fields {
			if fields[i] != host.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hq *HostQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hq.driver.Dialect())
	t1 := builder.Table(host.Table)
	columns := hq.ctx.Fields
	if len(columns) == 0 {
		columns = host.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hq.sql != nil {
		selector = hq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hq.ctx.Unique != nil && *hq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range hq.predicates {
		p(selector)
	}
	for _, p := range hq.order {
		p(selector)
	}
	if offset := hq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedTags tells the query-builder to eager-load the nodes that are connected to the "tags"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (hq *HostQuery) WithNamedTags(name string, opts ...func(*TagQuery)) *HostQuery {
	query := (&TagClient{config: hq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if hq.withNamedTags == nil {
		hq.withNamedTags = make(map[string]*TagQuery)
	}
	hq.withNamedTags[name] = query
	return hq
}

// WithNamedBeacons tells the query-builder to eager-load the nodes that are connected to the "beacons"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (hq *HostQuery) WithNamedBeacons(name string, opts ...func(*BeaconQuery)) *HostQuery {
	query := (&BeaconClient{config: hq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if hq.withNamedBeacons == nil {
		hq.withNamedBeacons = make(map[string]*BeaconQuery)
	}
	hq.withNamedBeacons[name] = query
	return hq
}

// WithNamedFiles tells the query-builder to eager-load the nodes that are connected to the "files"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (hq *HostQuery) WithNamedFiles(name string, opts ...func(*HostFileQuery)) *HostQuery {
	query := (&HostFileClient{config: hq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if hq.withNamedFiles == nil {
		hq.withNamedFiles = make(map[string]*HostFileQuery)
	}
	hq.withNamedFiles[name] = query
	return hq
}

// WithNamedProcesses tells the query-builder to eager-load the nodes that are connected to the "processes"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (hq *HostQuery) WithNamedProcesses(name string, opts ...func(*HostProcessQuery)) *HostQuery {
	query := (&HostProcessClient{config: hq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if hq.withNamedProcesses == nil {
		hq.withNamedProcesses = make(map[string]*HostProcessQuery)
	}
	hq.withNamedProcesses[name] = query
	return hq
}

// WithNamedCredentials tells the query-builder to eager-load the nodes that are connected to the "credentials"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (hq *HostQuery) WithNamedCredentials(name string, opts ...func(*HostCredentialQuery)) *HostQuery {
	query := (&HostCredentialClient{config: hq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if hq.withNamedCredentials == nil {
		hq.withNamedCredentials = make(map[string]*HostCredentialQuery)
	}
	hq.withNamedCredentials[name] = query
	return hq
}

// HostGroupBy is the group-by builder for Host entities.
type HostGroupBy struct {
	selector
	build *HostQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hgb *HostGroupBy) Aggregate(fns ...AggregateFunc) *HostGroupBy {
	hgb.fns = append(hgb.fns, fns...)
	return hgb
}

// Scan applies the selector query and scans the result into the given value.
func (hgb *HostGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hgb.build.ctx, ent.OpQueryGroupBy)
	if err := hgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HostQuery, *HostGroupBy](ctx, hgb.build, hgb, hgb.build.inters, v)
}

func (hgb *HostGroupBy) sqlScan(ctx context.Context, root *HostQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(hgb.fns))
	for _, fn := range hgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*hgb.flds)+len(hgb.fns))
		for _, f := range *hgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*hgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// HostSelect is the builder for selecting fields of Host entities.
type HostSelect struct {
	*HostQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hs *HostSelect) Aggregate(fns ...AggregateFunc) *HostSelect {
	hs.fns = append(hs.fns, fns...)
	return hs
}

// Scan applies the selector query and scans the result into the given value.
func (hs *HostSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hs.ctx, ent.OpQuerySelect)
	if err := hs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HostQuery, *HostSelect](ctx, hs.HostQuery, hs, hs.inters, v)
}

func (hs *HostSelect) sqlScan(ctx context.Context, root *HostQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(hs.fns))
	for _, fn := range hs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*hs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
