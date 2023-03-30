package schema

import "entgo.io/ent"

// Swagger holds the schema definition for the Swagger entity.
type Swagger struct {
	ent.Schema
}

// Fields of the Swagger.
func (Swagger) Fields() []ent.Field {
	return nil
}

// Edges of the Swagger.
func (Swagger) Edges() []ent.Edge {
	return nil
}
