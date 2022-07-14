// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/bug/ent/organizationalunit"
	"entgo.io/bug/ent/organizationalunittype"
	"entgo.io/bug/ent/specialty"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrganizationalUnitCreate is the builder for creating a OrganizationalUnit entity.
type OrganizationalUnitCreate struct {
	config
	mutation *OrganizationalUnitMutation
	hooks    []Hook
}

// SetDisplayName sets the "display_name" field.
func (ouc *OrganizationalUnitCreate) SetDisplayName(s string) *OrganizationalUnitCreate {
	ouc.mutation.SetDisplayName(s)
	return ouc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ouc *OrganizationalUnitCreate) SetNillableDisplayName(s *string) *OrganizationalUnitCreate {
	if s != nil {
		ouc.SetDisplayName(*s)
	}
	return ouc
}

// SetCreatedAt sets the "created_at" field.
func (ouc *OrganizationalUnitCreate) SetCreatedAt(t time.Time) *OrganizationalUnitCreate {
	ouc.mutation.SetCreatedAt(t)
	return ouc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ouc *OrganizationalUnitCreate) SetNillableCreatedAt(t *time.Time) *OrganizationalUnitCreate {
	if t != nil {
		ouc.SetCreatedAt(*t)
	}
	return ouc
}

// SetDeletedAt sets the "deleted_at" field.
func (ouc *OrganizationalUnitCreate) SetDeletedAt(t time.Time) *OrganizationalUnitCreate {
	ouc.mutation.SetDeletedAt(t)
	return ouc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ouc *OrganizationalUnitCreate) SetNillableDeletedAt(t *time.Time) *OrganizationalUnitCreate {
	if t != nil {
		ouc.SetDeletedAt(*t)
	}
	return ouc
}

// SetParentID sets the "parent_id" field.
func (ouc *OrganizationalUnitCreate) SetParentID(s string) *OrganizationalUnitCreate {
	ouc.mutation.SetParentID(s)
	return ouc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (ouc *OrganizationalUnitCreate) SetNillableParentID(s *string) *OrganizationalUnitCreate {
	if s != nil {
		ouc.SetParentID(*s)
	}
	return ouc
}

// SetMetadata sets the "metadata" field.
func (ouc *OrganizationalUnitCreate) SetMetadata(s []string) *OrganizationalUnitCreate {
	ouc.mutation.SetMetadata(s)
	return ouc
}

// SetTypeID sets the "type_id" field.
func (ouc *OrganizationalUnitCreate) SetTypeID(s string) *OrganizationalUnitCreate {
	ouc.mutation.SetTypeID(s)
	return ouc
}

// SetID sets the "id" field.
func (ouc *OrganizationalUnitCreate) SetID(s string) *OrganizationalUnitCreate {
	ouc.mutation.SetID(s)
	return ouc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ouc *OrganizationalUnitCreate) SetNillableID(s *string) *OrganizationalUnitCreate {
	if s != nil {
		ouc.SetID(*s)
	}
	return ouc
}

// SetParent sets the "parent" edge to the OrganizationalUnit entity.
func (ouc *OrganizationalUnitCreate) SetParent(o *OrganizationalUnit) *OrganizationalUnitCreate {
	return ouc.SetParentID(o.ID)
}

// AddChildIDs adds the "children" edge to the OrganizationalUnit entity by IDs.
func (ouc *OrganizationalUnitCreate) AddChildIDs(ids ...string) *OrganizationalUnitCreate {
	ouc.mutation.AddChildIDs(ids...)
	return ouc
}

// AddChildren adds the "children" edges to the OrganizationalUnit entity.
func (ouc *OrganizationalUnitCreate) AddChildren(o ...*OrganizationalUnit) *OrganizationalUnitCreate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouc.AddChildIDs(ids...)
}

// SetOrganizationalUnitTypeIDID sets the "organizational_unit_type_id" edge to the OrganizationalUnitType entity by ID.
func (ouc *OrganizationalUnitCreate) SetOrganizationalUnitTypeIDID(id string) *OrganizationalUnitCreate {
	ouc.mutation.SetOrganizationalUnitTypeIDID(id)
	return ouc
}

// SetOrganizationalUnitTypeID sets the "organizational_unit_type_id" edge to the OrganizationalUnitType entity.
func (ouc *OrganizationalUnitCreate) SetOrganizationalUnitTypeID(o *OrganizationalUnitType) *OrganizationalUnitCreate {
	return ouc.SetOrganizationalUnitTypeIDID(o.ID)
}

// AddSpecialtyIDs adds the "specialties" edge to the Specialty entity by IDs.
func (ouc *OrganizationalUnitCreate) AddSpecialtyIDs(ids ...string) *OrganizationalUnitCreate {
	ouc.mutation.AddSpecialtyIDs(ids...)
	return ouc
}

// AddSpecialties adds the "specialties" edges to the Specialty entity.
func (ouc *OrganizationalUnitCreate) AddSpecialties(s ...*Specialty) *OrganizationalUnitCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ouc.AddSpecialtyIDs(ids...)
}

// Mutation returns the OrganizationalUnitMutation object of the builder.
func (ouc *OrganizationalUnitCreate) Mutation() *OrganizationalUnitMutation {
	return ouc.mutation
}

// Save creates the OrganizationalUnit in the database.
func (ouc *OrganizationalUnitCreate) Save(ctx context.Context) (*OrganizationalUnit, error) {
	var (
		err  error
		node *OrganizationalUnit
	)
	ouc.defaults()
	if len(ouc.hooks) == 0 {
		if err = ouc.check(); err != nil {
			return nil, err
		}
		node, err = ouc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrganizationalUnitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ouc.check(); err != nil {
				return nil, err
			}
			ouc.mutation = mutation
			if node, err = ouc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ouc.hooks) - 1; i >= 0; i-- {
			if ouc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ouc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrganizationalUnit)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrganizationalUnitMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ouc *OrganizationalUnitCreate) SaveX(ctx context.Context) *OrganizationalUnit {
	v, err := ouc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ouc *OrganizationalUnitCreate) Exec(ctx context.Context) error {
	_, err := ouc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouc *OrganizationalUnitCreate) ExecX(ctx context.Context) {
	if err := ouc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouc *OrganizationalUnitCreate) defaults() {
	if _, ok := ouc.mutation.DisplayName(); !ok {
		v := organizationalunit.DefaultDisplayName
		ouc.mutation.SetDisplayName(v)
	}
	if _, ok := ouc.mutation.CreatedAt(); !ok {
		v := organizationalunit.DefaultCreatedAt()
		ouc.mutation.SetCreatedAt(v)
	}
	if _, ok := ouc.mutation.Metadata(); !ok {
		v := organizationalunit.DefaultMetadata
		ouc.mutation.SetMetadata(v)
	}
	if _, ok := ouc.mutation.ID(); !ok {
		v := organizationalunit.DefaultID
		ouc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouc *OrganizationalUnitCreate) check() error {
	if _, ok := ouc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`ent: missing required field "OrganizationalUnit.display_name"`)}
	}
	if _, ok := ouc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrganizationalUnit.created_at"`)}
	}
	if _, ok := ouc.mutation.Metadata(); !ok {
		return &ValidationError{Name: "metadata", err: errors.New(`ent: missing required field "OrganizationalUnit.metadata"`)}
	}
	if _, ok := ouc.mutation.TypeID(); !ok {
		return &ValidationError{Name: "type_id", err: errors.New(`ent: missing required field "OrganizationalUnit.type_id"`)}
	}
	if v, ok := ouc.mutation.ID(); ok {
		if err := organizationalunit.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "OrganizationalUnit.id": %w`, err)}
		}
	}
	if _, ok := ouc.mutation.OrganizationalUnitTypeIDID(); !ok {
		return &ValidationError{Name: "organizational_unit_type_id", err: errors.New(`ent: missing required edge "OrganizationalUnit.organizational_unit_type_id"`)}
	}
	return nil
}

func (ouc *OrganizationalUnitCreate) sqlSave(ctx context.Context) (*OrganizationalUnit, error) {
	_node, _spec := ouc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ouc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected OrganizationalUnit.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (ouc *OrganizationalUnitCreate) createSpec() (*OrganizationalUnit, *sqlgraph.CreateSpec) {
	var (
		_node = &OrganizationalUnit{config: ouc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: organizationalunit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: organizationalunit.FieldID,
			},
		}
	)
	if id, ok := ouc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ouc.mutation.DisplayName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organizationalunit.FieldDisplayName,
		})
		_node.DisplayName = value
	}
	if value, ok := ouc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: organizationalunit.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ouc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: organizationalunit.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := ouc.mutation.Metadata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: organizationalunit.FieldMetadata,
		})
		_node.Metadata = value
	}
	if nodes := ouc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationalunit.ParentTable,
			Columns: []string{organizationalunit.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: organizationalunit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ouc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organizationalunit.ChildrenTable,
			Columns: []string{organizationalunit.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: organizationalunit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ouc.mutation.OrganizationalUnitTypeIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationalunit.OrganizationalUnitTypeIDTable,
			Columns: []string{organizationalunit.OrganizationalUnitTypeIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: organizationalunittype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TypeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ouc.mutation.SpecialtiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   organizationalunit.SpecialtiesTable,
			Columns: organizationalunit.SpecialtiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: specialty.FieldID,
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

// OrganizationalUnitCreateBulk is the builder for creating many OrganizationalUnit entities in bulk.
type OrganizationalUnitCreateBulk struct {
	config
	builders []*OrganizationalUnitCreate
}

// Save creates the OrganizationalUnit entities in the database.
func (oucb *OrganizationalUnitCreateBulk) Save(ctx context.Context) ([]*OrganizationalUnit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oucb.builders))
	nodes := make([]*OrganizationalUnit, len(oucb.builders))
	mutators := make([]Mutator, len(oucb.builders))
	for i := range oucb.builders {
		func(i int, root context.Context) {
			builder := oucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationalUnitMutation)
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
					_, err = mutators[i+1].Mutate(root, oucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oucb *OrganizationalUnitCreateBulk) SaveX(ctx context.Context) []*OrganizationalUnit {
	v, err := oucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oucb *OrganizationalUnitCreateBulk) Exec(ctx context.Context) error {
	_, err := oucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oucb *OrganizationalUnitCreateBulk) ExecX(ctx context.Context) {
	if err := oucb.Exec(ctx); err != nil {
		panic(err)
	}
}