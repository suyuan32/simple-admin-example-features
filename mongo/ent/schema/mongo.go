package schema

import "entgo.io/ent"

// Mongo holds the schema definition for the Mongo entity.
type Mongo struct {
	ent.Schema
}

// Fields of the Mongo.
func (Mongo) Fields() []ent.Field {
	return nil
}

// Edges of the Mongo.
func (Mongo) Edges() []ent.Edge {
	return nil
}
