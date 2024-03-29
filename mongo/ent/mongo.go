// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-example-mongo/ent/mongo"
)

// Mongo is the model entity for the Mongo schema.
type Mongo struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Mongo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case mongo.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Mongo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Mongo fields.
func (m *Mongo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mongo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Mongo.
// Note that you need to call Mongo.Unwrap() before calling this method if this Mongo
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Mongo) Update() *MongoUpdateOne {
	return NewMongoClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Mongo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Mongo) Unwrap() *Mongo {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Mongo is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Mongo) String() string {
	var builder strings.Builder
	builder.WriteString("Mongo(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Mongos is a parsable slice of Mongo.
type Mongos []*Mongo
