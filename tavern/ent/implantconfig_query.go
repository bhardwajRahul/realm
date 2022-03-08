// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcarretto/realm/tavern/ent/deploymentconfig"
	"github.com/kcarretto/realm/tavern/ent/implant"
	"github.com/kcarretto/realm/tavern/ent/implantcallbackconfig"
	"github.com/kcarretto/realm/tavern/ent/implantconfig"
	"github.com/kcarretto/realm/tavern/ent/implantserviceconfig"
	"github.com/kcarretto/realm/tavern/ent/predicate"
)

// ImplantConfigQuery is the builder for querying ImplantConfig entities.
type ImplantConfigQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ImplantConfig
	// eager-loading edges.
	withDeploymentConfigs *DeploymentConfigQuery
	withImplants          *ImplantQuery
	withServiceConfigs    *ImplantServiceConfigQuery
	withCallbackConfigs   *ImplantCallbackConfigQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ImplantConfigQuery builder.
func (icq *ImplantConfigQuery) Where(ps ...predicate.ImplantConfig) *ImplantConfigQuery {
	icq.predicates = append(icq.predicates, ps...)
	return icq
}

// Limit adds a limit step to the query.
func (icq *ImplantConfigQuery) Limit(limit int) *ImplantConfigQuery {
	icq.limit = &limit
	return icq
}

// Offset adds an offset step to the query.
func (icq *ImplantConfigQuery) Offset(offset int) *ImplantConfigQuery {
	icq.offset = &offset
	return icq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (icq *ImplantConfigQuery) Unique(unique bool) *ImplantConfigQuery {
	icq.unique = &unique
	return icq
}

// Order adds an order step to the query.
func (icq *ImplantConfigQuery) Order(o ...OrderFunc) *ImplantConfigQuery {
	icq.order = append(icq.order, o...)
	return icq
}

// QueryDeploymentConfigs chains the current query on the "deploymentConfigs" edge.
func (icq *ImplantConfigQuery) QueryDeploymentConfigs() *DeploymentConfigQuery {
	query := &DeploymentConfigQuery{config: icq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := icq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := icq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(implantconfig.Table, implantconfig.FieldID, selector),
			sqlgraph.To(deploymentconfig.Table, deploymentconfig.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, implantconfig.DeploymentConfigsTable, implantconfig.DeploymentConfigsColumn),
		)
		fromU = sqlgraph.SetNeighbors(icq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryImplants chains the current query on the "implants" edge.
func (icq *ImplantConfigQuery) QueryImplants() *ImplantQuery {
	query := &ImplantQuery{config: icq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := icq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := icq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(implantconfig.Table, implantconfig.FieldID, selector),
			sqlgraph.To(implant.Table, implant.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, implantconfig.ImplantsTable, implantconfig.ImplantsColumn),
		)
		fromU = sqlgraph.SetNeighbors(icq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryServiceConfigs chains the current query on the "serviceConfigs" edge.
func (icq *ImplantConfigQuery) QueryServiceConfigs() *ImplantServiceConfigQuery {
	query := &ImplantServiceConfigQuery{config: icq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := icq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := icq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(implantconfig.Table, implantconfig.FieldID, selector),
			sqlgraph.To(implantserviceconfig.Table, implantserviceconfig.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, implantconfig.ServiceConfigsTable, implantconfig.ServiceConfigsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(icq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCallbackConfigs chains the current query on the "callbackConfigs" edge.
func (icq *ImplantConfigQuery) QueryCallbackConfigs() *ImplantCallbackConfigQuery {
	query := &ImplantCallbackConfigQuery{config: icq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := icq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := icq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(implantconfig.Table, implantconfig.FieldID, selector),
			sqlgraph.To(implantcallbackconfig.Table, implantcallbackconfig.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, implantconfig.CallbackConfigsTable, implantconfig.CallbackConfigsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(icq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ImplantConfig entity from the query.
// Returns a *NotFoundError when no ImplantConfig was found.
func (icq *ImplantConfigQuery) First(ctx context.Context) (*ImplantConfig, error) {
	nodes, err := icq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{implantconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (icq *ImplantConfigQuery) FirstX(ctx context.Context) *ImplantConfig {
	node, err := icq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ImplantConfig ID from the query.
// Returns a *NotFoundError when no ImplantConfig ID was found.
func (icq *ImplantConfigQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = icq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{implantconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (icq *ImplantConfigQuery) FirstIDX(ctx context.Context) int {
	id, err := icq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ImplantConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ImplantConfig entity is found.
// Returns a *NotFoundError when no ImplantConfig entities are found.
func (icq *ImplantConfigQuery) Only(ctx context.Context) (*ImplantConfig, error) {
	nodes, err := icq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{implantconfig.Label}
	default:
		return nil, &NotSingularError{implantconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (icq *ImplantConfigQuery) OnlyX(ctx context.Context) *ImplantConfig {
	node, err := icq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ImplantConfig ID in the query.
// Returns a *NotSingularError when more than one ImplantConfig ID is found.
// Returns a *NotFoundError when no entities are found.
func (icq *ImplantConfigQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = icq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{implantconfig.Label}
	default:
		err = &NotSingularError{implantconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (icq *ImplantConfigQuery) OnlyIDX(ctx context.Context) int {
	id, err := icq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ImplantConfigs.
func (icq *ImplantConfigQuery) All(ctx context.Context) ([]*ImplantConfig, error) {
	if err := icq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return icq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (icq *ImplantConfigQuery) AllX(ctx context.Context) []*ImplantConfig {
	nodes, err := icq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ImplantConfig IDs.
func (icq *ImplantConfigQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := icq.Select(implantconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (icq *ImplantConfigQuery) IDsX(ctx context.Context) []int {
	ids, err := icq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (icq *ImplantConfigQuery) Count(ctx context.Context) (int, error) {
	if err := icq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return icq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (icq *ImplantConfigQuery) CountX(ctx context.Context) int {
	count, err := icq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (icq *ImplantConfigQuery) Exist(ctx context.Context) (bool, error) {
	if err := icq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return icq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (icq *ImplantConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := icq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ImplantConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (icq *ImplantConfigQuery) Clone() *ImplantConfigQuery {
	if icq == nil {
		return nil
	}
	return &ImplantConfigQuery{
		config:                icq.config,
		limit:                 icq.limit,
		offset:                icq.offset,
		order:                 append([]OrderFunc{}, icq.order...),
		predicates:            append([]predicate.ImplantConfig{}, icq.predicates...),
		withDeploymentConfigs: icq.withDeploymentConfigs.Clone(),
		withImplants:          icq.withImplants.Clone(),
		withServiceConfigs:    icq.withServiceConfigs.Clone(),
		withCallbackConfigs:   icq.withCallbackConfigs.Clone(),
		// clone intermediate query.
		sql:    icq.sql.Clone(),
		path:   icq.path,
		unique: icq.unique,
	}
}

// WithDeploymentConfigs tells the query-builder to eager-load the nodes that are connected to
// the "deploymentConfigs" edge. The optional arguments are used to configure the query builder of the edge.
func (icq *ImplantConfigQuery) WithDeploymentConfigs(opts ...func(*DeploymentConfigQuery)) *ImplantConfigQuery {
	query := &DeploymentConfigQuery{config: icq.config}
	for _, opt := range opts {
		opt(query)
	}
	icq.withDeploymentConfigs = query
	return icq
}

// WithImplants tells the query-builder to eager-load the nodes that are connected to
// the "implants" edge. The optional arguments are used to configure the query builder of the edge.
func (icq *ImplantConfigQuery) WithImplants(opts ...func(*ImplantQuery)) *ImplantConfigQuery {
	query := &ImplantQuery{config: icq.config}
	for _, opt := range opts {
		opt(query)
	}
	icq.withImplants = query
	return icq
}

// WithServiceConfigs tells the query-builder to eager-load the nodes that are connected to
// the "serviceConfigs" edge. The optional arguments are used to configure the query builder of the edge.
func (icq *ImplantConfigQuery) WithServiceConfigs(opts ...func(*ImplantServiceConfigQuery)) *ImplantConfigQuery {
	query := &ImplantServiceConfigQuery{config: icq.config}
	for _, opt := range opts {
		opt(query)
	}
	icq.withServiceConfigs = query
	return icq
}

// WithCallbackConfigs tells the query-builder to eager-load the nodes that are connected to
// the "callbackConfigs" edge. The optional arguments are used to configure the query builder of the edge.
func (icq *ImplantConfigQuery) WithCallbackConfigs(opts ...func(*ImplantCallbackConfigQuery)) *ImplantConfigQuery {
	query := &ImplantCallbackConfigQuery{config: icq.config}
	for _, opt := range opts {
		opt(query)
	}
	icq.withCallbackConfigs = query
	return icq
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
//	client.ImplantConfig.Query().
//		GroupBy(implantconfig.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (icq *ImplantConfigQuery) GroupBy(field string, fields ...string) *ImplantConfigGroupBy {
	group := &ImplantConfigGroupBy{config: icq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := icq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return icq.sqlQuery(ctx), nil
	}
	return group
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
//	client.ImplantConfig.Query().
//		Select(implantconfig.FieldName).
//		Scan(ctx, &v)
//
func (icq *ImplantConfigQuery) Select(fields ...string) *ImplantConfigSelect {
	icq.fields = append(icq.fields, fields...)
	return &ImplantConfigSelect{ImplantConfigQuery: icq}
}

func (icq *ImplantConfigQuery) prepareQuery(ctx context.Context) error {
	for _, f := range icq.fields {
		if !implantconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if icq.path != nil {
		prev, err := icq.path(ctx)
		if err != nil {
			return err
		}
		icq.sql = prev
	}
	return nil
}

func (icq *ImplantConfigQuery) sqlAll(ctx context.Context) ([]*ImplantConfig, error) {
	var (
		nodes       = []*ImplantConfig{}
		_spec       = icq.querySpec()
		loadedTypes = [4]bool{
			icq.withDeploymentConfigs != nil,
			icq.withImplants != nil,
			icq.withServiceConfigs != nil,
			icq.withCallbackConfigs != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ImplantConfig{config: icq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, icq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := icq.withDeploymentConfigs; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ImplantConfig)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.DeploymentConfigs = []*DeploymentConfig{}
		}
		query.withFKs = true
		query.Where(predicate.DeploymentConfig(func(s *sql.Selector) {
			s.Where(sql.InValues(implantconfig.DeploymentConfigsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.deployment_config_implant_config
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "deployment_config_implant_config" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "deployment_config_implant_config" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.DeploymentConfigs = append(node.Edges.DeploymentConfigs, n)
		}
	}

	if query := icq.withImplants; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ImplantConfig)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Implants = []*Implant{}
		}
		query.withFKs = true
		query.Where(predicate.Implant(func(s *sql.Selector) {
			s.Where(sql.InValues(implantconfig.ImplantsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.implant_config
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "implant_config" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "implant_config" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Implants = append(node.Edges.Implants, n)
		}
	}

	if query := icq.withServiceConfigs; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*ImplantConfig, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.ServiceConfigs = []*ImplantServiceConfig{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*ImplantConfig)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   implantconfig.ServiceConfigsTable,
				Columns: implantconfig.ServiceConfigsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(implantconfig.ServiceConfigsPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, icq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "serviceConfigs": %w`, err)
		}
		query.Where(implantserviceconfig.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "serviceConfigs" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ServiceConfigs = append(nodes[i].Edges.ServiceConfigs, n)
			}
		}
	}

	if query := icq.withCallbackConfigs; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*ImplantConfig, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.CallbackConfigs = []*ImplantCallbackConfig{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*ImplantConfig)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   implantconfig.CallbackConfigsTable,
				Columns: implantconfig.CallbackConfigsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(implantconfig.CallbackConfigsPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, icq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "callbackConfigs": %w`, err)
		}
		query.Where(implantcallbackconfig.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "callbackConfigs" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.CallbackConfigs = append(nodes[i].Edges.CallbackConfigs, n)
			}
		}
	}

	return nodes, nil
}

func (icq *ImplantConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := icq.querySpec()
	_spec.Node.Columns = icq.fields
	if len(icq.fields) > 0 {
		_spec.Unique = icq.unique != nil && *icq.unique
	}
	return sqlgraph.CountNodes(ctx, icq.driver, _spec)
}

func (icq *ImplantConfigQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := icq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (icq *ImplantConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   implantconfig.Table,
			Columns: implantconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: implantconfig.FieldID,
			},
		},
		From:   icq.sql,
		Unique: true,
	}
	if unique := icq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := icq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, implantconfig.FieldID)
		for i := range fields {
			if fields[i] != implantconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := icq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := icq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := icq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := icq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (icq *ImplantConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(icq.driver.Dialect())
	t1 := builder.Table(implantconfig.Table)
	columns := icq.fields
	if len(columns) == 0 {
		columns = implantconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if icq.sql != nil {
		selector = icq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if icq.unique != nil && *icq.unique {
		selector.Distinct()
	}
	for _, p := range icq.predicates {
		p(selector)
	}
	for _, p := range icq.order {
		p(selector)
	}
	if offset := icq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := icq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ImplantConfigGroupBy is the group-by builder for ImplantConfig entities.
type ImplantConfigGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (icgb *ImplantConfigGroupBy) Aggregate(fns ...AggregateFunc) *ImplantConfigGroupBy {
	icgb.fns = append(icgb.fns, fns...)
	return icgb
}

// Scan applies the group-by query and scans the result into the given value.
func (icgb *ImplantConfigGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := icgb.path(ctx)
	if err != nil {
		return err
	}
	icgb.sql = query
	return icgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (icgb *ImplantConfigGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := icgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (icgb *ImplantConfigGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(icgb.fields) > 1 {
		return nil, errors.New("ent: ImplantConfigGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := icgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (icgb *ImplantConfigGroupBy) StringsX(ctx context.Context) []string {
	v, err := icgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (icgb *ImplantConfigGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = icgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implantconfig.Label}
	default:
		err = fmt.Errorf("ent: ImplantConfigGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (icgb *ImplantConfigGroupBy) StringX(ctx context.Context) string {
	v, err := icgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (icgb *ImplantConfigGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(icgb.fields) > 1 {
		return nil, errors.New("ent: ImplantConfigGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := icgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (icgb *ImplantConfigGroupBy) IntsX(ctx context.Context) []int {
	v, err := icgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (icgb *ImplantConfigGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = icgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implantconfig.Label}
	default:
		err = fmt.Errorf("ent: ImplantConfigGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (icgb *ImplantConfigGroupBy) IntX(ctx context.Context) int {
	v, err := icgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (icgb *ImplantConfigGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(icgb.fields) > 1 {
		return nil, errors.New("ent: ImplantConfigGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := icgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (icgb *ImplantConfigGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := icgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (icgb *ImplantConfigGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = icgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implantconfig.Label}
	default:
		err = fmt.Errorf("ent: ImplantConfigGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (icgb *ImplantConfigGroupBy) Float64X(ctx context.Context) float64 {
	v, err := icgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (icgb *ImplantConfigGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(icgb.fields) > 1 {
		return nil, errors.New("ent: ImplantConfigGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := icgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (icgb *ImplantConfigGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := icgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (icgb *ImplantConfigGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = icgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implantconfig.Label}
	default:
		err = fmt.Errorf("ent: ImplantConfigGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (icgb *ImplantConfigGroupBy) BoolX(ctx context.Context) bool {
	v, err := icgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (icgb *ImplantConfigGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range icgb.fields {
		if !implantconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := icgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := icgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (icgb *ImplantConfigGroupBy) sqlQuery() *sql.Selector {
	selector := icgb.sql.Select()
	aggregation := make([]string, 0, len(icgb.fns))
	for _, fn := range icgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(icgb.fields)+len(icgb.fns))
		for _, f := range icgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(icgb.fields...)...)
}

// ImplantConfigSelect is the builder for selecting fields of ImplantConfig entities.
type ImplantConfigSelect struct {
	*ImplantConfigQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ics *ImplantConfigSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ics.prepareQuery(ctx); err != nil {
		return err
	}
	ics.sql = ics.ImplantConfigQuery.sqlQuery(ctx)
	return ics.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ics *ImplantConfigSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ics.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ics *ImplantConfigSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ics.fields) > 1 {
		return nil, errors.New("ent: ImplantConfigSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ics.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ics *ImplantConfigSelect) StringsX(ctx context.Context) []string {
	v, err := ics.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ics *ImplantConfigSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ics.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implantconfig.Label}
	default:
		err = fmt.Errorf("ent: ImplantConfigSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ics *ImplantConfigSelect) StringX(ctx context.Context) string {
	v, err := ics.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ics *ImplantConfigSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ics.fields) > 1 {
		return nil, errors.New("ent: ImplantConfigSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ics.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ics *ImplantConfigSelect) IntsX(ctx context.Context) []int {
	v, err := ics.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ics *ImplantConfigSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ics.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implantconfig.Label}
	default:
		err = fmt.Errorf("ent: ImplantConfigSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ics *ImplantConfigSelect) IntX(ctx context.Context) int {
	v, err := ics.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ics *ImplantConfigSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ics.fields) > 1 {
		return nil, errors.New("ent: ImplantConfigSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ics.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ics *ImplantConfigSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ics.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ics *ImplantConfigSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ics.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implantconfig.Label}
	default:
		err = fmt.Errorf("ent: ImplantConfigSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ics *ImplantConfigSelect) Float64X(ctx context.Context) float64 {
	v, err := ics.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ics *ImplantConfigSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ics.fields) > 1 {
		return nil, errors.New("ent: ImplantConfigSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ics.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ics *ImplantConfigSelect) BoolsX(ctx context.Context) []bool {
	v, err := ics.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ics *ImplantConfigSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ics.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implantconfig.Label}
	default:
		err = fmt.Errorf("ent: ImplantConfigSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ics *ImplantConfigSelect) BoolX(ctx context.Context) bool {
	v, err := ics.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ics *ImplantConfigSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ics.sql.Query()
	if err := ics.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}