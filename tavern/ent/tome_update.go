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
	"github.com/kcarretto/realm/tavern/ent/file"
	"github.com/kcarretto/realm/tavern/ent/predicate"
	"github.com/kcarretto/realm/tavern/ent/tome"
)

// TomeUpdate is the builder for updating Tome entities.
type TomeUpdate struct {
	config
	hooks    []Hook
	mutation *TomeMutation
}

// Where appends a list predicates to the TomeUpdate builder.
func (tu *TomeUpdate) Where(ps ...predicate.Tome) *TomeUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetLastModifiedAt sets the "lastModifiedAt" field.
func (tu *TomeUpdate) SetLastModifiedAt(t time.Time) *TomeUpdate {
	tu.mutation.SetLastModifiedAt(t)
	return tu
}

// SetName sets the "name" field.
func (tu *TomeUpdate) SetName(s string) *TomeUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TomeUpdate) SetDescription(s string) *TomeUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetParameters sets the "parameters" field.
func (tu *TomeUpdate) SetParameters(s string) *TomeUpdate {
	tu.mutation.SetParameters(s)
	return tu
}

// SetNillableParameters sets the "parameters" field if the given value is not nil.
func (tu *TomeUpdate) SetNillableParameters(s *string) *TomeUpdate {
	if s != nil {
		tu.SetParameters(*s)
	}
	return tu
}

// ClearParameters clears the value of the "parameters" field.
func (tu *TomeUpdate) ClearParameters() *TomeUpdate {
	tu.mutation.ClearParameters()
	return tu
}

// SetHash sets the "hash" field.
func (tu *TomeUpdate) SetHash(s string) *TomeUpdate {
	tu.mutation.SetHash(s)
	return tu
}

// SetEldritch sets the "eldritch" field.
func (tu *TomeUpdate) SetEldritch(s string) *TomeUpdate {
	tu.mutation.SetEldritch(s)
	return tu
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (tu *TomeUpdate) AddFileIDs(ids ...int) *TomeUpdate {
	tu.mutation.AddFileIDs(ids...)
	return tu
}

// AddFiles adds the "files" edges to the File entity.
func (tu *TomeUpdate) AddFiles(f ...*File) *TomeUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tu.AddFileIDs(ids...)
}

// Mutation returns the TomeMutation object of the builder.
func (tu *TomeUpdate) Mutation() *TomeMutation {
	return tu.mutation
}

// ClearFiles clears all "files" edges to the File entity.
func (tu *TomeUpdate) ClearFiles() *TomeUpdate {
	tu.mutation.ClearFiles()
	return tu
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (tu *TomeUpdate) RemoveFileIDs(ids ...int) *TomeUpdate {
	tu.mutation.RemoveFileIDs(ids...)
	return tu
}

// RemoveFiles removes "files" edges to File entities.
func (tu *TomeUpdate) RemoveFiles(f ...*File) *TomeUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tu.RemoveFileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TomeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := tu.defaults(); err != nil {
		return 0, err
	}
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TomeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TomeUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TomeUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TomeUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TomeUpdate) defaults() error {
	if _, ok := tu.mutation.LastModifiedAt(); !ok {
		if tome.UpdateDefaultLastModifiedAt == nil {
			return fmt.Errorf("ent: uninitialized tome.UpdateDefaultLastModifiedAt (forgotten import ent/runtime?)")
		}
		v := tome.UpdateDefaultLastModifiedAt()
		tu.mutation.SetLastModifiedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tu *TomeUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := tome.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tome.name": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Hash(); ok {
		if err := tome.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "Tome.hash": %w`, err)}
		}
	}
	return nil
}

func (tu *TomeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tome.Table,
			Columns: tome.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tome.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.LastModifiedAt(); ok {
		_spec.SetField(tome.FieldLastModifiedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(tome.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(tome.FieldDescription, field.TypeString, value)
	}
	if value, ok := tu.mutation.Parameters(); ok {
		_spec.SetField(tome.FieldParameters, field.TypeString, value)
	}
	if tu.mutation.ParametersCleared() {
		_spec.ClearField(tome.FieldParameters, field.TypeString)
	}
	if value, ok := tu.mutation.Hash(); ok {
		_spec.SetField(tome.FieldHash, field.TypeString, value)
	}
	if value, ok := tu.mutation.Eldritch(); ok {
		_spec.SetField(tome.FieldEldritch, field.TypeString, value)
	}
	if tu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tome.FilesTable,
			Columns: []string{tome.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedFilesIDs(); len(nodes) > 0 && !tu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tome.FilesTable,
			Columns: []string{tome.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tome.FilesTable,
			Columns: []string{tome.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tome.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TomeUpdateOne is the builder for updating a single Tome entity.
type TomeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TomeMutation
}

// SetLastModifiedAt sets the "lastModifiedAt" field.
func (tuo *TomeUpdateOne) SetLastModifiedAt(t time.Time) *TomeUpdateOne {
	tuo.mutation.SetLastModifiedAt(t)
	return tuo
}

// SetName sets the "name" field.
func (tuo *TomeUpdateOne) SetName(s string) *TomeUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TomeUpdateOne) SetDescription(s string) *TomeUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetParameters sets the "parameters" field.
func (tuo *TomeUpdateOne) SetParameters(s string) *TomeUpdateOne {
	tuo.mutation.SetParameters(s)
	return tuo
}

// SetNillableParameters sets the "parameters" field if the given value is not nil.
func (tuo *TomeUpdateOne) SetNillableParameters(s *string) *TomeUpdateOne {
	if s != nil {
		tuo.SetParameters(*s)
	}
	return tuo
}

// ClearParameters clears the value of the "parameters" field.
func (tuo *TomeUpdateOne) ClearParameters() *TomeUpdateOne {
	tuo.mutation.ClearParameters()
	return tuo
}

// SetHash sets the "hash" field.
func (tuo *TomeUpdateOne) SetHash(s string) *TomeUpdateOne {
	tuo.mutation.SetHash(s)
	return tuo
}

// SetEldritch sets the "eldritch" field.
func (tuo *TomeUpdateOne) SetEldritch(s string) *TomeUpdateOne {
	tuo.mutation.SetEldritch(s)
	return tuo
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (tuo *TomeUpdateOne) AddFileIDs(ids ...int) *TomeUpdateOne {
	tuo.mutation.AddFileIDs(ids...)
	return tuo
}

// AddFiles adds the "files" edges to the File entity.
func (tuo *TomeUpdateOne) AddFiles(f ...*File) *TomeUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tuo.AddFileIDs(ids...)
}

// Mutation returns the TomeMutation object of the builder.
func (tuo *TomeUpdateOne) Mutation() *TomeMutation {
	return tuo.mutation
}

// ClearFiles clears all "files" edges to the File entity.
func (tuo *TomeUpdateOne) ClearFiles() *TomeUpdateOne {
	tuo.mutation.ClearFiles()
	return tuo
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (tuo *TomeUpdateOne) RemoveFileIDs(ids ...int) *TomeUpdateOne {
	tuo.mutation.RemoveFileIDs(ids...)
	return tuo
}

// RemoveFiles removes "files" edges to File entities.
func (tuo *TomeUpdateOne) RemoveFiles(f ...*File) *TomeUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tuo.RemoveFileIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TomeUpdateOne) Select(field string, fields ...string) *TomeUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tome entity.
func (tuo *TomeUpdateOne) Save(ctx context.Context) (*Tome, error) {
	var (
		err  error
		node *Tome
	)
	if err := tuo.defaults(); err != nil {
		return nil, err
	}
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TomeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Tome)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TomeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TomeUpdateOne) SaveX(ctx context.Context) *Tome {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TomeUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TomeUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TomeUpdateOne) defaults() error {
	if _, ok := tuo.mutation.LastModifiedAt(); !ok {
		if tome.UpdateDefaultLastModifiedAt == nil {
			return fmt.Errorf("ent: uninitialized tome.UpdateDefaultLastModifiedAt (forgotten import ent/runtime?)")
		}
		v := tome.UpdateDefaultLastModifiedAt()
		tuo.mutation.SetLastModifiedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TomeUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := tome.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tome.name": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Hash(); ok {
		if err := tome.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "Tome.hash": %w`, err)}
		}
	}
	return nil
}

func (tuo *TomeUpdateOne) sqlSave(ctx context.Context) (_node *Tome, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tome.Table,
			Columns: tome.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tome.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tome.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tome.FieldID)
		for _, f := range fields {
			if !tome.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tome.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.LastModifiedAt(); ok {
		_spec.SetField(tome.FieldLastModifiedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(tome.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(tome.FieldDescription, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Parameters(); ok {
		_spec.SetField(tome.FieldParameters, field.TypeString, value)
	}
	if tuo.mutation.ParametersCleared() {
		_spec.ClearField(tome.FieldParameters, field.TypeString)
	}
	if value, ok := tuo.mutation.Hash(); ok {
		_spec.SetField(tome.FieldHash, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Eldritch(); ok {
		_spec.SetField(tome.FieldEldritch, field.TypeString, value)
	}
	if tuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tome.FilesTable,
			Columns: []string{tome.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedFilesIDs(); len(nodes) > 0 && !tuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tome.FilesTable,
			Columns: []string{tome.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tome.FilesTable,
			Columns: []string{tome.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tome{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tome.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}