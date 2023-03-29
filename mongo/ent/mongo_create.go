// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-example-mongo/ent/mongo"
)

// MongoCreate is the builder for creating a Mongo entity.
type MongoCreate struct {
	config
	mutation *MongoMutation
	hooks    []Hook
}

// Mutation returns the MongoMutation object of the builder.
func (mc *MongoCreate) Mutation() *MongoMutation {
	return mc.mutation
}

// Save creates the Mongo in the database.
func (mc *MongoCreate) Save(ctx context.Context) (*Mongo, error) {
	return withHooks[*Mongo, MongoMutation](ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MongoCreate) SaveX(ctx context.Context) *Mongo {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MongoCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MongoCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MongoCreate) check() error {
	return nil
}

func (mc *MongoCreate) sqlSave(ctx context.Context) (*Mongo, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MongoCreate) createSpec() (*Mongo, *sqlgraph.CreateSpec) {
	var (
		_node = &Mongo{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(mongo.Table, sqlgraph.NewFieldSpec(mongo.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// MongoCreateBulk is the builder for creating many Mongo entities in bulk.
type MongoCreateBulk struct {
	config
	builders []*MongoCreate
}

// Save creates the Mongo entities in the database.
func (mcb *MongoCreateBulk) Save(ctx context.Context) ([]*Mongo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Mongo, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MongoMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MongoCreateBulk) SaveX(ctx context.Context) []*Mongo {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MongoCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MongoCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
