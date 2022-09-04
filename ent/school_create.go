// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/municipio"
	"github.com/vmkevv/rigelapi/ent/school"
)

// SchoolCreate is the builder for creating a School entity.
type SchoolCreate struct {
	config
	mutation *SchoolMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *SchoolCreate) SetName(s string) *SchoolCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetLat sets the "lat" field.
func (sc *SchoolCreate) SetLat(s string) *SchoolCreate {
	sc.mutation.SetLat(s)
	return sc
}

// SetLon sets the "lon" field.
func (sc *SchoolCreate) SetLon(s string) *SchoolCreate {
	sc.mutation.SetLon(s)
	return sc
}

// SetID sets the "id" field.
func (sc *SchoolCreate) SetID(s string) *SchoolCreate {
	sc.mutation.SetID(s)
	return sc
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (sc *SchoolCreate) AddClassIDs(ids ...string) *SchoolCreate {
	sc.mutation.AddClassIDs(ids...)
	return sc
}

// AddClasses adds the "classes" edges to the Class entity.
func (sc *SchoolCreate) AddClasses(c ...*Class) *SchoolCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddClassIDs(ids...)
}

// SetMunicipioID sets the "municipio" edge to the Municipio entity by ID.
func (sc *SchoolCreate) SetMunicipioID(id string) *SchoolCreate {
	sc.mutation.SetMunicipioID(id)
	return sc
}

// SetNillableMunicipioID sets the "municipio" edge to the Municipio entity by ID if the given value is not nil.
func (sc *SchoolCreate) SetNillableMunicipioID(id *string) *SchoolCreate {
	if id != nil {
		sc = sc.SetMunicipioID(*id)
	}
	return sc
}

// SetMunicipio sets the "municipio" edge to the Municipio entity.
func (sc *SchoolCreate) SetMunicipio(m *Municipio) *SchoolCreate {
	return sc.SetMunicipioID(m.ID)
}

// Mutation returns the SchoolMutation object of the builder.
func (sc *SchoolCreate) Mutation() *SchoolMutation {
	return sc.mutation
}

// Save creates the School in the database.
func (sc *SchoolCreate) Save(ctx context.Context) (*School, error) {
	var (
		err  error
		node *School
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SchoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*School)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SchoolMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SchoolCreate) SaveX(ctx context.Context) *School {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SchoolCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SchoolCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SchoolCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "School.name"`)}
	}
	if _, ok := sc.mutation.Lat(); !ok {
		return &ValidationError{Name: "lat", err: errors.New(`ent: missing required field "School.lat"`)}
	}
	if _, ok := sc.mutation.Lon(); !ok {
		return &ValidationError{Name: "lon", err: errors.New(`ent: missing required field "School.lon"`)}
	}
	return nil
}

func (sc *SchoolCreate) sqlSave(ctx context.Context) (*School, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected School.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (sc *SchoolCreate) createSpec() (*School, *sqlgraph.CreateSpec) {
	var (
		_node = &School{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: school.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: school.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldName,
		})
		_node.Name = value
	}
	if value, ok := sc.mutation.Lat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldLat,
		})
		_node.Lat = value
	}
	if value, ok := sc.mutation.Lon(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldLon,
		})
		_node.Lon = value
	}
	if nodes := sc.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   school.ClassesTable,
			Columns: []string{school.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.MunicipioIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   school.MunicipioTable,
			Columns: []string{school.MunicipioColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: municipio.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.municipio_schools = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SchoolCreateBulk is the builder for creating many School entities in bulk.
type SchoolCreateBulk struct {
	config
	builders []*SchoolCreate
}

// Save creates the School entities in the database.
func (scb *SchoolCreateBulk) Save(ctx context.Context) ([]*School, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*School, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SchoolMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SchoolCreateBulk) SaveX(ctx context.Context) []*School {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SchoolCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SchoolCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
