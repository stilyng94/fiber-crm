package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Lead holds the schema definition for the Lead entity.
type Lead struct {
	ent.Schema
}

// Fields of the Lead.
func (Lead) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(), field.String("company").NotEmpty(), field.String("email").NotEmpty(),
		field.String("phone").NotEmpty()}
}

// Edges of the Lead.
func (Lead) Edges() []ent.Edge {
	return nil
}
