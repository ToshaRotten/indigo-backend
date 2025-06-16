package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TypeOfPackaging holds the schema definition for the TypeOfPackaging entity.
type TypeOfPackaging struct {
	ent.Schema
}

// Fields of the TypeOfPackaging.
func (TypeOfPackaging) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().MaxLen(255),
	}
}

// Edges of the TypeOfPackaging.
func (TypeOfPackaging) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("type_of_packaging", Product.Type),
	}
}
